import * as pulumi from "@pulumi/pulumi";
import * as runpod from "@runpod-infra/pulumi";

const myTemplate = new runpod.Template("TranscriptionTemplate", {
  containerDiskInGb: 5,
  isPublic: false,
  isServerless: true,
  name: "Test StuB Blah 3",
  readme: "# Test",
  ports: "8888/http,22/tcp",
  volumeInGb: 0,
  dockerArgs: "echo hello",
  env: [
    {
      key: "key3",
      value: "value333",
    },
    {
      key: "key4",
      value: "value4",
    },
  ],
  imageName: "runpod/serverless-hello-world:1234",
});

// const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
//   name: "testStorage1",
//   size: 20,
//   dataCenterId: "US-NJ",
// });

// const myRandomPod = new runpod.Pod("myRandomPod", {
//   cloudType: "ALL",
//   networkVolumeId: testNetworkStorage.networkStorage.apply(
//     (networkStorage) => networkStorage.id
//   ),
//   templateId: myTemplate.id,
//   gpuCount: 1,
//   volumeInGb: 50,
//   containerDiskInGb: 50,
//   minVcpuCount: 2,
//   minMemoryInGb: 15,
//   gpuTypeId: "NVIDIA GeForce RTX 3070",
//   name: "RunPod Pytorch",
//   imageName: "runpod/pytorch",
//   dockerArgs: "",
//   ports: "8888/http",
//   volumeMountPath: "/workspace",
//   env: [
//     {
//       key: "JUPYTER_PASSWORD",
//       value: "rns1hunbsstltcpad22d",
//     },
//   ],
// });

// export const pod = {
//   value: myRandomPod.pod,
// };

// export const networkStorage = {
//   value: testNetworkStorage.networkStorage,
// };

export const template = {
  value: myTemplate,
};
