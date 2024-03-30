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

This is an example of how to deploy over typescript. We also serve pulumi over Python and Go. For more examples, please go read through the examples in the examples directory
or the documents inside docs. If you have any problems in doing so, please contact support@runpod.io.

Filename must be index.ts for Typescript.

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

> **Note:** For examples in Go and Python, please visit the documentation inside the docs directory or click [here](https://github.com/runpod/pulumi-runpod-native/tree/main/docs).