---
title: Runpod
meta_desc: Provides an overview of the Runpod Provider for Pulumi.
layout: package
---

The Runpod provider for Pulumi can be used to provision [Runpod](https://www.runpod.io) resources.
The Runpod provider must be configured with Runpod's API keys to deploy and update resources in Runpod.

## Config

To begin with, please set your runpod API key using Pulumi.

```bash
  pulumi config set --secret runpod:token YOUR_API_KEY
```

## Example

{{< chooser language "typescript,go,python" >}}

{{% choosable language typescript %}}

Filename must be index.ts.

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

{{% /choosable %}}

{{% choosable language go %}}

File name must be main.go.

```go
    package main

    import (
        "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
        "github.com/runpod/pulumi-runpod-native/sdk/go/runpod"
    )

    func main() {
        pulumi.Run(func(ctx *pulumi.Context) error {
            testNetworkStorage, err := runpod.NewNetworkStorage(ctx, "testNetworkStorage", &runpod.NetworkStorageArgs{
                Name:         pulumi.String("testStorage1"),
                Size:         pulumi.Int(20),
                DataCenterId: pulumi.String("US-NJ"),
            })
            if err != nil {
                return err
            }
            myRandomPod, err := runpod.NewPod(ctx, "myRandomPod", &runpod.PodArgs{
                CloudType: pulumi.String("ALL"),
                NetworkVolumeId: testNetworkStorage.NetworkStorage.ApplyT(func(networkStorage runpod.NetworkStorageType) (*string, error) {
                    return &networkStorage.Id, nil
                }).(pulumi.StringPtrOutput),
                GpuCount:          pulumi.Int(1),
                VolumeInGb:        pulumi.Int(50),
                ContainerDiskInGb: pulumi.Int(50),
                MinVcpuCount:      pulumi.Int(2),
                MinMemoryInGb:     pulumi.Int(15),
                GpuTypeId:         pulumi.String("NVIDIA GeForce RTX 3070"),
                Name:              pulumi.String("RunPod Pytorch"),
                ImageName:         pulumi.String("runpod/pytorch"),
                DockerArgs:        pulumi.String(""),
                Ports:             pulumi.String("8888/http"),
                VolumeMountPath:   pulumi.String("/workspace"),
                Env: runpod.PodEnvArray{
                    &runpod.PodEnvArgs{
                        Key:   pulumi.String("JUPYTER_PASSWORD"),
                        Value: pulumi.String("rns1hunbsstltcpad22d"),
                    },
                },
            })
            if err != nil {
                return err
            }
            ctx.Export("pod", myRandomPod)
		    ctx.Export("networkStorage", testNetworkStorage)
            return nil
        })
    }
```

{{% /choosable %}}

{{% choosable language python %}}

File name must be __main__.py

```python
    import pulumi
    import pulumi_runpod as runpod
    from loguru import logger
    import json

    def fetch_id(a):
        if type(a) == runpod.outputs.NetworkStorage:
            return a.id
        else:
            return a

    try:
        test_network_storage = runpod.NetworkStorage(
                "testNetworkStorage", name="testStorage1", size=20, data_center_id="US-NJ"
            )
        my_random_pod = runpod.Pod(
            "myRandomPod",
            cloud_type="ALL",
            network_volume_id=test_network_storage.network_storage.apply(lambda x : fetch_id(x)),
            gpu_count=1,
            volume_in_gb=50,
            container_disk_in_gb=50,
            min_vcpu_count=2,
            min_memory_in_gb=15,
            gpu_type_id="NVIDIA GeForce RTX 3070",
            name="RunPod Pytorch",
            image_name="runpod/pytorch",
            docker_args="",
            ports="8888/http",
            volume_mount_path="/workspace",
            env=[
                runpod.PodEnvArgs(
                    key="JUPYTER_PASSWORD",
                    value="rns1hunbsstltcpad22d",
                ).__dict__,
            ],
        )
        pulumi.export(
            "pod",
            {
                "value": my_random_pod.pod,
            },
        )
        pulumi.export(
            "networkStorage",
            {
                "value": test_network_storage.network_storage,
            },
        )
    except Exception as e:
        logger.exception(e)
```

{{% /choosable %}}

{{% choosable language yaml %}}

name: provider-runpod-native
runtime: yaml

```yaml
    resources:
    plugins:
    providers:
        - name: runpod
    testNetworkStorage:
        type: runpod:NetworkStorage
        properties:
        name: "testStorage1"
        size: 20
        dataCenterId: "US-NJ"

    myRandomPod:
        type: runpod:Pod
        properties:
        cloudType: ALL
        networkVolumeId: ${testNetworkStorage.networkStorage.id}
        gpuCount: 1
        volumeInGb: 50
        containerDiskInGb: 50
        minVcpuCount: 2
        minMemoryInGb: 15
        gpuTypeId: "NVIDIA GeForce RTX 3070"
        name: "RunPod Pytorch"
        imageName: "runpod/pytorch"
        dockerArgs: ""
        ports: "8888/http"
        volumeMountPath: "/workspace"
        env: [{ key: "JUPYTER_PASSWORD", value: "rns1hunbsstltcpad22d" }]

    outputs:
    pod:
        value: ${myRandomPod.pod}
    networkStorage:
        value: ${testNetworkStorage.networkStorage}
```