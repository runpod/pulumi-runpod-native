# Publishing Guide for Runpod Pulumi Provider

This guide follows the official Pulumi publishing requirements from https://www.pulumi.com/docs/iac/build-with-pulumi/publishing-packages/

## Prerequisites

- [ ] Pulumi CLI installed
- [ ] pulumictl installed (`go install github.com/pulumi/pulumictl/cmd/pulumictl@latest`)
- [ ] npm account with @runpod scope access
- [ ] PyPI account with upload permissions
- [ ] NuGet.org account
- [ ] GitHub repository under runpod organization

## Package Names (Updated)

- **npm**: `@runpod/pulumi-runpod`
- **PyPI**: `pulumi-runpod`
- **NuGet**: `Pulumi.Runpod`
- **Go Module**: `github.com/runpod/pulumi-runpod-native/sdk`

## Build Process

```bash
# Install pulumictl if needed
go install github.com/pulumi/pulumictl/cmd/pulumictl@latest

# Set version
export VERSION=v1.10.0

# Build all SDKs
make build
```

## Publishing to Package Registries

### 1. Node.js (npm)

```bash
# Login to npm with @runpod scope
npm login --scope=@runpod

# Publish from the built package
cd sdk/nodejs/bin
npm publish --access public
```

### 2. Python (PyPI)

```bash
# Install twine
pip install twine

# Upload to PyPI
cd sdk/python
python -m twine upload dist/*
```

### 3. .NET (NuGet)

```bash
# Find the built package
cd sdk/dotnet

# Push to NuGet
dotnet nuget push Pulumi.Runpod.*.nupkg \
  --api-key YOUR_NUGET_API_KEY \
  --source https://api.nuget.org/v3/index.json
```

### 4. Go Module

```bash
# Tag the release
git tag sdk/v${VERSION}
git push origin sdk/v${VERSION}
```

## Plugin Binary Release

Create a GitHub release with the provider binary:

```bash
# Create release with pulumictl
pulumictl create release \
  --tag v${VERSION} \
  --repo runpod/pulumi-runpod-native
```

## Pulumi Registry Listing

1. Fork https://github.com/pulumi/registry

2. Add Runpod to `community-packages/package-list.json`:

```json
{
  "repoName": "runpod",
  "repoOrg": "runpod",
  "provider": "runpod",
  "schemaFile": "provider/schema.json"
}
```

3. Submit PR to the Pulumi registry repository

## Deprecation of Old Packages

```bash
# Deprecate old npm package
npm deprecate @runpod-infra/pulumi \
  "This package has moved to @runpod/pulumi-runpod"
```

For PyPI, release a final version with deprecation notice in the description.

## Required Files Checklist

- [x] `provider/schema.json` with metadata:
  - [x] displayName: "Runpod"
  - [x] description
  - [x] publisher: "Runpod"
  - [x] logoUrl
  - [x] keywords
  - [x] pluginDownloadURL
- [x] `docs/_index.md` (overview)
- [x] `docs/installation-configuration.md`
- [x] `docs/logo.png`
- [x] `sdk/nodejs/package.json` with publishConfig
- [x] Registry listing JSON prepared

## CI/CD Automation

Consider setting up GitHub Actions for automated releases:

```yaml
# .github/workflows/release.yml
name: Release
on:
  push:
    tags:
      - v*
jobs:
  publish:
    # Automated build and publish steps
```

## Support

For partnership and additional support, contact Pulumi at partners@pulumi.com