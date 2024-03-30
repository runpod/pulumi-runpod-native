---
title: Runpod Setup
meta_desc: Information on how to install the Runpod Provider for Pulumi.
layout: package
---

## Installation

The Runpod provider is available as a package in all Runpod languages:

* JavaScript/TypeScript: [`@runpod-infra/pulumi`](https://www.npmjs.com/package/@runpod-infra/pulumi)
* Python: [`pulumi-runpod`](https://pypi.org/project/pulumi-runpod/)
* Go: [`github.com/runpod/pulumi-runpod-native/tree/main/sdk/go/runpod`](https://www.github.com/runpod/pulumi-runpod-native)

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
  go get github.com/runpod/pulumi-runpod-native/sdk/go/runpod
```
We advise you to pin to a specific version.

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

To use from Python, follow the following steps:

```bash
  pip install pulumi-runpod
```

## Configuration

To begin with, please set your runpod API key using Pulumi.

```bash
  pulumi config set runpod:token --secret
```