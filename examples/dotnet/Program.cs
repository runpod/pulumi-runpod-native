using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Runpod = Pulumi.Runpod;

return await Deployment.RunAsync(() => 
{
    var testNetworkStorage = new Runpod.NetworkStorage("testNetworkStorage", new()
    {
        Name = "testStorage1",
        Size = 20,
        DataCenterId = "US-NJ",
    });

    var myRandomPod = new Runpod.Pod("myRandomPod", new()
    {
        CloudType = "ALL",
        NetworkVolumeId = testNetworkStorage.NetworkStorage.Apply(networkStorage => networkStorage.Id),
        GpuCount = 1,
        VolumeInGb = 50,
        ContainerDiskInGb = 50,
        MinVcpuCount = 2,
        MinMemoryInGb = 15,
        GpuTypeId = "NVIDIA GeForce RTX 3070",
        Name = "RunPod Pytorch",
        ImageName = "runpod/pytorch",
        DockerArgs = "",
        Ports = "8888/http",
        VolumeMountPath = "/workspace",
        Env = new[]
        {
            new Runpod.Inputs.PodEnvArgs
            {
                Key = "JUPYTER_PASSWORD",
                Value = "rns1hunbsstltcpad22d",
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["pod"] = 
        {
            { "value", myRandomPod.Pod },
        },
        ["networkStorage"] = 
        {
            { "value", testNetworkStorage.NetworkStorage },
        },
    };
});

