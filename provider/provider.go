// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version = "v1.2.7"

const Name string = "runpod"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*NetworkStorage, NetworkStorageArgs, NetworkStorageState](),
			infer.Resource[*Template, TemplateArgs, TemplateState](),
			infer.Resource[*Pod, PodArgs, PodState](),
			infer.Resource[*Endpoint, EndpointArgs, EndpointState](),
		},
		Config: infer.Config[*Config](),
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "Runpod",
			Description: "The Runpod Pulumi provider provides resources to interact with Runpod's native APIs.",
			Keywords: []string{
				"pulumi",
				"runpod",
				"gpus",
				"ml",
				"ai",
			},
			Homepage:          "https://runpod.io",
			License:           "Apache-2.0",
			Repository:        "https://github.com/runpod/pulumi-runpod-native",
			PluginDownloadURL: "github://api.github.com/runpod/pulumi-runpod-native",
			Publisher:         "Runpod",
			LogoURL:           "https://avatars.githubusercontent.com/u/95939477?s=200&v=4",
			LanguageMap: map[string]interface{}{
				"go": map[string]interface{}{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/runpod/pulumi-runpod-native/sdk/go/runpod",
				},
				"nodejs": map[string]interface{}{
					"packageName": "@runpod-infra/pulumi",
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.42.0",
					},
					"devDependencies": map[string]string{
						"typescript": "^3.7.0",
					},
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					"packageName": "runpodinfra",
				},
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
					"rootNamespace": "RunpodInfra",
				},
			},
		},
	})
}

type Config struct {
	Token string `pulumi:"token"`
}

func (c *Config) Configure(ctx p.Context) error {
	if c.Token == "" {
		Token, exists := os.LookupEnv("RUNPOD_API_KEY")
		if exists {
			c.Token = Token
			return nil
		}
		return fmt.Errorf("API key is required")
	}
	return nil
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Token, "Runpod API Token")
}

func Schema(version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	s, err := integration.NewServer("runpod", semver.MustParse(version), Provider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
