PROJECT_NAME := Pulumi Runpod Resource Provider

PACK             := runpod
PACKDIR          := sdk
PROJECT          := github.com/runpod/pulumi-runpod-native
NODE_MODULE_NAME := @runpod-infra/pulumi
NUGET_PKG_NAME   := Pulumi.Runpod

PROVIDER        := pulumi-resource-${PACK}
VERSION         ?= $(shell pulumictl get version)
PROVIDER_PATH   := provider
VERSION_PATH    := ${PROVIDER_PATH}.Version

GOPATH			:= $(shell go env GOPATH)
CODEGEN         := pulumi-gen-${PACK}
WORKING_DIR     := $(shell pwd)
EXAMPLES_DIR    := ${WORKING_DIR}/examples/yaml
TESTPARALLELISM := 4
SCHEMA_FILE     := provider/schema.json
PULUMI_CONVERT := 0

ensure::
	cd provider && go mod tidy
	cd sdk && go mod tidy
	cd tests && go mod tidy

provider:: codegen generate
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

codegen::
	(cd provider && VERSION=${VERSION} go generate cmd/${PROVIDER}/main.go)
	(cd provider && go build -o $(WORKING_DIR)/bin/${CODEGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/$(CODEGEN))
	$(WORKING_DIR)/bin/${CODEGEN} $(SCHEMA_FILE) --version ${VERSION}

generate:
	@echo "Generating Go client from Swagger definition..."
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go generate ./${PROVIDER_PATH}/provider.go

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

test_provider::
	cd tests && go test -short -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM} ./...

dotnet_sdk:: DOTNET_VERSION := ${VERSION}
dotnet_sdk::
	rm -rf sdk/dotnet
	pulumi package gen-sdk --language dotnet $(SCHEMA_FILE)
	cd ${PACKDIR}/dotnet/&& \
		echo "${DOTNET_VERSION}" >version.txt

go_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/go
	pulumi package gen-sdk --language go $(SCHEMA_FILE)

nodejs_sdk:: VERSION := ${VERSION}
nodejs_sdk::
	rm -rf sdk/nodejs
	PULUMI_CONVERT=$(PULUMI_CONVERT) PULUMI_DISABLE_AUTOMATIC_PLUGIN_ACQUISITION=$(PULUMI_CONVERT) pulumi package gen-sdk --language nodejs $(SCHEMA_FILE)
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock bin/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' bin/package.json && \
		rm ./bin/package.json.bak

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S), Linux)
    SED_CMD = sed -i -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' setup.py
	SED_CMD_REMOVE_OS = sed -i -e '/^import os$$/d' setup.py
	SED_REMOVE_PYI_LINE = sed -i -e '/\.pyi$$/d' sdk/python/runpodinfra.egg-info/SOURCES.txt
else ifeq ($(UNAME_S), Darwin)
    SED_CMD = sed -i '' -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' setup.py
	SED_CMD_REMOVE_OS = sed -i '' '/^import os$$/d' setup.py
	SED_REMOVE_PYI_LINE = sed -i "" '/\.pyi$$/d' sdk/python/runpodinfra.egg-info/SOURCES.txt
else
    $(error Unsupported OS: $(UNAME_S))
endif

python_sdk:: PYPI_VERSION := ${VERSION}
python_sdk::
	rm -rf sdk/python
	PULUMI_CONVERT=$(PULUMI_CONVERT) PULUMI_DISABLE_AUTOMATIC_PLUGIN_ACQUISITION=$(PULUMI_CONVERT) pulumi package gen-sdk --language python $(SCHEMA_FILE)
	pip install -r requirements.txt
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		$(SED_CMD) && \
		$(SED_CMD_REMOVE_OS) && \
		python3 -m pip list && \
		python3 setup.py clean --all 2>/dev/null && \
		python3 setup.py build sdist

# The codegened examples are shit
# gen_examples: gen_go_example \
# 		gen_nodejs_example \
# 		gen_python_example \
# 		gen_dotnet_example

# The codegened examples are shit
# gen_%_example:
# 	rm -rf ${WORKING_DIR}/examples/$*
# 	pulumi convert \
# 		--cwd ${WORKING_DIR}/examples/yaml \
# 		--logtostderr \
# 		--generate-only \
# 		--non-interactive \
# 		--language $* \
# 		--out ${WORKING_DIR}/examples/$*

define pulumi_login
    export PULUMI_CONFIG_PASSPHRASE=asdfqwerty1234; \
    pulumi login --local;
endef

up::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack init dev && \
	pulumi stack select dev && \
	pulumi config set name dev && \
	pulumi up --config="runpod:token=${RUNPOD_TOKEN}" -y

down::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack select dev && \
	pulumi destroy -y && \
	pulumi stack rm dev -y

devcontainer::
	git submodule update --init --recursive .devcontainer
	git submodule update --remote --merge .devcontainer
	cp -f .devcontainer/devcontainer.json .devcontainer.json

.PHONY: build

build:: provider python_sdk go_sdk nodejs_sdk
	rm -rf sdk/python/build/lib/runpodinfra/config/__init__.pyi

build-and-push:
	VERSION=$(VERSION) $(MAKE) build
	git add .
	git commit -am "Bump version to $(VERSION)"
	git push --force
	git tag $(VERSION)
	git push origin $(VERSION)

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint::
	for DIR in "provider" "sdk" "tests" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

# install_dotnet_sdk
install:: install_nodejs_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

GO_TEST 	 := go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

test_all:: test_provider
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/dotnet && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

run-ci-locally:
	act -P catthehacker/ubuntu:full-latest -e tag.json --secret-file .secrets --container-architecture linux/amd64 -W .github/workflows/release.yml