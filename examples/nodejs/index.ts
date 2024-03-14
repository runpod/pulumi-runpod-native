import * as pulumi from "@pulumi/pulumi";
import * as runpod from "@pulumi/runpod";

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
