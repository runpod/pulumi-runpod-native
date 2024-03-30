---
title: Runpod
meta_desc: Provides an overview of the Runpod Provider for Pulumi.
layout: package
---

The Runpod provider for Pulumi can be used to provision [Runpod](https://www.runpod.io) resources.
The Runpod provider must be configured with Runpod's API keys to deploy and update resources in Runpod.

## Config

To begin with, please set your runpod API key to use with Pulumi.

```bash
  pulumi config set --secret runpod:token
```

## Example

This is an example of how to deploy it over Golang. We also serve pulumi over Typescript and Python. For more examples, please navigate to the examples directory
or the documents inside docs. If you have any problems in doing so, please contact support@runpod.io.

1. Create a new Pulumi Go example:
```
    pulumi new
```
Select either the Go template or Runpod's Go template.

2. Set your API keys using the config shown above. 

3. Install the official Go package:

```
    go get github.com/runpod/pulumi-runpod-native/sdk/go/runpod@v1.1.8
```
Replace the version above to any that you want. We advise you to pin a certain version as there will be fewer breaking changes.

4. Use this example as a simple building guide for your example project:

```typescript
    import * as pulumi from "@pulumi/pulumi";
    import * as runpod from "@runpod-infra/pulumi";

    const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
        name: "testStorage1",
        size: 20,
        dataCenterId: "US-NJ",
    });
    const myRandomPod = new runpod.Pod("myRandomPod", {
        cloudType: "ALL",
        networkVolumeId: testNetworkStorage.networkStorage.apply(networkStorage => networkStorage.id),
        gpuCount: 1,
        volumeInGb: 50,
        containerDiskInGb: 50,
        minVcpuCount: 2,
        minMemoryInGb: 15,
        gpuTypeId: "NVIDIA GeForce RTX 3070",
        name: "RunPod Pytorch",
        imageName: "runpod/pytorch",
        dockerArgs: "",
        ports: "8888/http",
        volumeMountPath: "/workspace",
        env: [{
            key: "JUPYTER_PASSWORD",
            value: "rns1hunbsstltcpad22d",
        }],
    });
    export const pod = {
        value: myRandomPod.pod,
    };
    export const networkStorage = {
        value: testNetworkStorage.networkStorage,
    };
```

5. PULUMI UP
Create your resources using the command below:

```
    pulumi up
```

6. PULUMI DOWN
If you want to remove your resources, you can use the command below:

```
    pulumi down
```

If you have any issues, please feel free to create an issue or reach out to us directly at support@runpod.io.

> **Note:** For examples in Go and Python, please visit the documentation inside the docs directory or click [here](https://github.com/runpod/pulumi-runpod-native/tree/main/docs).