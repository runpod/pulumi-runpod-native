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
			CloudType:         pulumi.String("ALL"),
			NetworkVolumeId:   testNetworkStorage.NetworkStorage.Id(),
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

> **Note:** For examples in TypeScript and Python, please visit the documentation inside the docs directory or click [here](https://github.com/runpod/pulumi-runpod-native/tree/main/docs).