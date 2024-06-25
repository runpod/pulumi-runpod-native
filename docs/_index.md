---
title: RunPod
meta_desc: Provides an overview of the RunPod Provider for Pulumi.
layout: package
---

The RunPod provider for Pulumi can be used to provision [RunPod](https://www.runpod.io) resources. The RunPod provider must be configured with RunPod's API keys to deploy and update resources in RunPod.

## Config

To begin with, please set your RunPod API key using Pulumi.

```bash
pulumi config set --secret runpod:token
```

## Note

Please make sure that you are inside the Python virtual environment created by Pulumi when using the Python SDK.

## Example

{{< chooser language "typescript,go,python,yaml" />}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as runpod from "@runpod-infra/pulumi";

const myTemplate = new runpod.Template("testTemplate", {
  containerDiskInGb: 5,
  dockerArgs: "python handler.py",
  env: [
    {
      key: "key1",
      value: "value1",
    },
    {
      key: "key2",
      value: "value2",
    },
  ],
  imageName: "runpod/serverless-hello-world:latest",
  isServerless: true,
  name: "Testing Pulumi V1",
  readme: "## Hello, World!",
  volumeInGb: 0,
});

const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
  name: "testStorage1",
  size: 10,
  dataCenterId: "US-OR-1",
});

const myEndpoint = new runpod.Endpoint("testEndpoint", {
  gpuIds: "AMPERE_16",
  name: "Pulumi Endpoint Test V2 -fb",
  templateId: myTemplate.template.id,
  workersMax: 2,
  workersMin: 1,
  idleTimeout: 6,
  locations: "US-OR-1",
  networkVolumeId: testNetworkStorage.networkStorage.id,
  scalerType: "QUEUE_DELAY",
  scalerValue: 4,
});

const myRandomPod = new runpod.Pod("myRandomPod", {
  cloudType: "ALL",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    (networkStorage) => networkStorage.id
  ),
  gpuCount: 1,
  volumeInGb: 50,
  containerDiskInGb: 50,
  minVcpuCount: 2,
  minMemoryInGb: 15,
  gpuTypeId: "NVIDIA GeForce RTX 4090",
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
export const template = {
  value: myTemplate.template,
};
export const endpoint = {
  value: myEndpoint.endpoint,
};
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
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/runpod/pulumi-runpod-native/sdk/go/runpod"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testNetworkStorage, err := runpod.NewNetworkStorage(ctx, "testNetworkStorage", &runpod.NetworkStorageArgs{
			Name:         pulumi.String("testStorage1"),
			Size:         pulumi.Int(5),
			DataCenterId: pulumi.String("EU-RO-1"),
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
			GpuTypeId:         pulumi.String("NVIDIA GeForce RTX 4090"),
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
		myTemplate, err := runpod.NewTemplate(ctx, "myTemplate", &runpod.TemplateArgs{
			ContainerDiskInGb:       pulumi.Int(5),
			DockerArgs:              pulumi.String("python3 -m http.server 8080"),
		    Env: runpod.PodEnvArray{
				&runpod.PodEnvArgs{
					Key:   pulumi.String("JUPYTER_PASSWORD"),
					Value: pulumi.String("rns1hunbsstltcpad22d"),
				},
			},
			ImageName:       pulumi.String("runpod/serverless-hello-world:latest"),
			IsServerless:    pulumi.Bool(true),
			Name:            pulumi.String("Testing Pulumi V1"),
			Readme:          pulumi.String("## Hello, World!"),
			VolumeInGb:      pulumi.Int(0),
		})

		if err != nil {
			return err
		}

		myEndpoint, err := runpod.NewEndpoint(ctx, "myEndpoint", &runpod.EndpointArgs{
			GpuIds:         pulumi.String("AMPERE_16"),
			Name:           pulumi.String("Pulumi Endpoint Test V2 -fb"),
			TemplateId:     myTemplate.Template.Id(),
			WorkersMax:     pulumi.Int(2),
			WorkersMin:     pulumi.Int(1),
			IdleTimeout:    pulumi.Int(6),
			Locations:      pulumi.String("EU-RO-1"),
			NetworkVolumeId: testNetworkStorage.NetworkStorage.Id(),
			ScalerType:     pulumi.String("QUEUE_DELAY"),
			ScalerValue:    pulumi.Int(4),
		})

		if err != nil {
			return err
		}

		ctx.Export("pod", myRandomPod.Pod.Id())
		ctx.Export("networkStorage", testNetworkStorage.NetworkStorage.Id())
		ctx.Export("template", myTemplate.Template.Id())
		ctx.Export("endpoint", myEndpoint.Endpoint.Id())
		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language python %}}

```
  source venv/bin/activate
```

```python
import pulumi
import runpodinfra as runpod
from loguru import logger
import json

def fetch_id(a):
    if type(a) == runpod.outputs.NetworkStorage:
        return a.id
    else:
        return a

try:
    test_network_storage = runpod.NetworkStorage(
            "testNetworkStorage", name="testStorage1", size=10, data_center_id="US-OR-1"
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
        gpu_type_id="NVIDIA GeForce RTX 4090",
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
    my_template = runpod.Template(
        "myTemplate",
        container_disk_in_gb=5,
        docker_args="python handler.py",
        env=[
            runpod.PodEnvArgs(
                key="key1",
                value="value1",
            ).__dict__,
            runpod.PodEnvArgs(
                key="key2",
                value="value2",
            ).__dict__,
        ],
        image_name="runpod/serverless-hello-world:latest",
        is_serverless=True,
        name="Testing Pulumi V1",
        readme="## Hello, World!",
        volume_in_gb=0,
    )

    my_endpoint = runpod.Endpoint(
        "myEndpoint",
        gpu_ids="AMPERE_16",
        name="Pulumi Endpoint Test V2 -fb",
        template_id=my_template.template.id,
        workers_max=2,
        workers_min=1,
        idle_timeout=6,
        locations="US-OR-1",
        network_volume_id=test_network_storage.network_storage.apply(lambda x: fetch_id(x)),
        scaler_type="QUEUE_DELAY",
        scaler_value=4,
    )

    print(my_random_pod)
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
    pulumi.export(
        "template",
        {
            "value": my_template.template,
        },
    )
    pulumi.export(
        "endpoint",
        {
            "value": my_endpoint.endpoint,
        },
    )
except Exception as e:
    logger.exception(e)
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
resources:
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
      env:
        - key: "JUPYTER_PASSWORD"
          value: "rns1hunbsstltcpad22d"

  myTemplate:
    type: runpod:Template
    properties:
      containerDiskInGb: 5
      dockerArgs: "python handler.py"
      env:
        - key: "key1"
          value: "value1"
        - key: "key2"
          value: "value2"
      imageName: "runpod/serverless-hello-world:latest"
      isServerless: true
      name: "Testing Pulumi V1"
      readme: "## Hello, World!"
      volumeInGb: 0

  myEndpoint:
    type: runpod:Endpoint
    properties:
      gpuIds: "AMPERE_16"
      name: "Pulumi Endpoint Test V2 -fb"
      templateId: ${myTemplate.template.id}
      workersMax: 2
      workersMin: 1
      idleTimeout: 6
      locations: "US-OR-1"
      networkVolumeId: ${testNetworkStorage.networkStorage.id}
      scalerType: "QUEUE_DELAY"
      scalerValue: 4

outputs:
  pod: ${myRandomPod.pod}
  networkStorage: ${testNetworkStorage.networkStorage}
  template: ${myTemplate.template}
  endpoint: ${myEndpoint.endpoint}
```

{{% /choosable %}}
