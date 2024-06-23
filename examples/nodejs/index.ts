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

const myEndpoint = new runpod.Endpoint("testEndpoint", {
  //   gpuIds: "AMPERE_16",
  name: "Pulumi Endpoint Test V1",
  templateId: myTemplate.template.id, // .apply((template) => template.id),
  workersMax: 2,
  workersMin: 1,
  gpuIds: "ADA_24,AMPERE_24,AMPERE_16",
  idleTimeout: 5,
  //   locations: null,
  //   networkVolumeId: null,
  scalerType: "QUEUE_DELAY",
  scalerValue: 4,
});
// const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
//     name: "testStorage1",
//     size: 5,
//     dataCenterId: "EU-RO-1",
// });
// const myRandomPod = new runpod.Pod("myRandomPod", {
//     cloudType: "ALL",
//     networkVolumeId: testNetworkStorage.networkStorage.apply(networkStorage => networkStorage.id),
//     gpuCount: 1,
//     volumeInGb: 50,
//     containerDiskInGb: 50,
//     minVcpuCount: 2,
//     minMemoryInGb: 15,
//     gpuTypeId: "NVIDIA GeForce RTX 4090",
//     name: "RunPod Pytorch",
//     imageName: "runpod/pytorch",
//     dockerArgs: "",
//     ports: "8888/http",
//     volumeMountPath: "/workspace",
//     env: [{
//         key: "JUPYTER_PASSWORD",
//         value: "rns1hunbsstltcpad22d",
//     }],
// });
export const template = {
  value: myTemplate.template,
};
export const endpoint = {
  value: myEndpoint.endpoint,
};
// export const pod = {
//     value: myRandomPod.pod,
// };
// export const networkStorage = {
//     value: testNetworkStorage.networkStorage,
// };
