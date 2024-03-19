---
title: Runpod Setup
meta_desc: Information on how to install the Runpod Provider for Pulumi.
layout: package
---

## Installation

The Runpod provider is available as a package in all Runpod languages:

* JavaScript/TypeScript: [`@runpod-infra/pulumi`](https://www.npmjs.com/package/@runpod-infra/pulumi)
* Python: [`runpod_pulumi`](https://pypi.org/project/runpod_pulumi/)
* Go: [`github.com/runpod/pulumi-runpod-native/tree/main/sdk/go/runpod`](https://www.github.com/runpod/pulumi-runpod-native)

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
  npm install @runpod-infra/pulumi
```

or `yarn`:

```bash
  yarn add @runpod-infra/pulumi
```

### Python

To use from Python, install using `pip`:

```bash
  pip install runpod_pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
  go get github.com/runpod/pulumi-runpod-native
```

## Configuration

To begin with, please set your runpod API key using Pulumi.

```bash
  pulumi config set --secret runpod:token YOUR_API_KEY
```