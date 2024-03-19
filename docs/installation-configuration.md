---
title: Runpod Setup
meta_desc: Information on how to install the Runpod Provider for Pulumi.
layout: package
---

## Installation

The Runpod provider is available as a package in all Runpod languages:

* JavaScript/TypeScript: [`@runpod/pulumi`](https://www.npmjs.com/package/@runpod/pulumi)
* Python: [`runpod_pulumi`](https://pypi.org/project/runpod_pulumi/)
* Go: [`github.com/runpod/pulumi-runpod-native/tree/main/sdk/go/runpod`](https://www.github.com/runpod/pulumi-runpod-native)

## Config

To begin with, please set your runpod API key using Pulumi.

```bash
  pulumi config set --secret runpod:token YOUR_API_KEY
```

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
  npm install @runpod/pulumi
```

or `yarn`:

```bash
  yarn add @runpod/pulumi
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

The following configuration points are available for the `runpod` provider:

- `runpod:token` - This is the Runpod API key.
Support for `RUNPOD_API_KEY` environment key is coming soon.