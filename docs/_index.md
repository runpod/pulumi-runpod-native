---
title: Runpod
meta_desc: Provides an overview of the Runpod Provider for Pulumi.
layout: package
---

The Runpod provider for Pulumi can be used to provision Runpod resources.
The Runpod provider must be configured with Runpod's API keys to deploy and update resources in Aquasec.

## Example

{{< chooser language "typescript,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as runpod from "@pierre78181/runpod";

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

{{% /choosable %}}
{{% choosable language go %}}

```go
COMING SOON
```