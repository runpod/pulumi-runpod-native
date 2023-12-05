package main

import (
	"github.com/pulumi/pulumi-runpod/sdk/go/runpod"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
			GpuTypeId:         pulumi.String("NVIDIA GeForce RTX 3080"),
			Name:              pulumi.String("RunPod Tensorflow"),
			ImageName:         pulumi.String("runpod/tensorflow"),
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
		ctx.Export("pod", map[string]interface{}{
			"value": myRandomPod.Pod,
		})
		ctx.Export("networkStorage", map[string]interface{}{
			"value": testNetworkStorage.NetworkStorage,
		})
		return nil
	})
}
