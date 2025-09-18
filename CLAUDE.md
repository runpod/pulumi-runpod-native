# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Build and Test
- `make build` - Build provider and all SDKs (Go, Node.js, Python, .NET)
- `make provider` - Build only the provider binary
- `make test_provider` - Run provider tests with parallel execution
- `make test_all` - Run all tests including SDK tests
- `make lint` - Run Go linter across provider, sdk, and tests directories

### SDK Generation
- `make codegen` - Generate code from schema
- `make go_sdk` - Generate Go SDK
- `make nodejs_sdk` - Generate Node.js SDK  
- `make python_sdk` - Generate Python SDK
- `make dotnet_sdk` - Generate .NET SDK

### Development Workflow
- `make ensure` - Update go.mod files in provider, sdk, and tests
- `make generate` - Generate Go client from Swagger definition using oapi-codegen

### Testing Integration
- `make up` - Deploy example resources to test environment
- `make down` - Destroy test resources and clean up stack

## Architecture Overview

This is a Pulumi native provider for Runpod's cloud GPU platform, built using the pulumi-go-provider framework.

### Core Components

**Provider Structure:**
- `provider/provider.go` - Main provider implementation with resource registration
- `provider/schema.json` - JSON schema defining all resources and their properties
- Individual resource files: `pod.go`, `template.go`, `networkStorage.go`, `endpoint.go`

**Resource Types:**
- **Pod** - GPU compute instances with configurable specs, storage, and environments
- **Template** - Reusable container configurations for serverless and persistent workloads  
- **NetworkStorage** - Persistent network-attached storage volumes
- **Endpoint** - Serverless API endpoints with auto-scaling capabilities

**SDK Generation:**
- Multi-language SDK support (Go, Node.js, Python, .NET) generated from single schema
- Language-specific packaging and distribution via npm, PyPI, NuGet, and Go modules

### Key Patterns

- Resources use the `infer.Resource` pattern with separate Args/State structs
- HTTP client interactions use GraphQL mutations for resource operations
- Configuration via `runpod:token` secret for API authentication
- All resources support standard Pulumi lifecycle (Create, Read, Update, Delete)

### Testing Structure

- `tests/` directory contains provider integration tests
- Example configurations in `examples/` for each supported language
- YAML examples for testing resource deployments

## Configuration

Set Runpod API token:
```bash
pulumi config set runpod:token --secret <your-token>
```

Or via environment variable:
```bash
export RUNPOD_API_KEY=<your-token>
```