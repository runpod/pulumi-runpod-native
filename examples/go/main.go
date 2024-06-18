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
		myRandomTemplate, err := runpod.NewTemplate(ctx, "myRandomTemplate", &runpod.TemplateArgs{
			ContainerDiskInGb:       pulumi.Int(5),
			ContainerRegistryAuthId: pulumi.String(""),
			DockerArgs:              pulumi.String("python3 -m http.server 8080"),
			Env: runpod.PodEnvArray{
				&runpod.PodEnvArgs{
					Key:   pulumi.String("JUPYTER_PASSWORD"),
					Value: pulumi.String("rns1hunbsstltcpad22d"),
				},
			},
			ImageName:       pulumi.String("nginx:latest"),
			IsPublic:        pulumi.BoolPtr(false),
			IsServerless:    pulumi.BoolPtr(false),
			Name:            pulumi.String("RunPod Nginx"),
			Ports:           pulumi.String("8080/http"),
			Readme:          pulumi.String("Please set this even if you don't have a readme"),
			StartJupyter:    pulumi.BoolPtr(false),
			StartSsh:        pulumi.BoolPtr(false),
			VolumeInGb:      pulumi.Int(5),
			VolumeMountPath: pulumi.String("/usr/share/nginx/html"),
		})
		if err != nil {
			return err
		}

		ctx.Export("pod", myRandomPod.Pod.Id())
		ctx.Export("networkStorage", testNetworkStorage.NetworkStorage.Id())
		ctx.Export("template", myRandomTemplate.Template.Id())
		return nil
	})
}
