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
  name: "Generated Serverless Template",
  readme: "## Hello, World!",
  volumeInGb: 0,
});

const testNetworkStorage = new runpod.NetworkStorage("testNetworkStorage", {
  name: "testStorage1",
  size: 20,
  dataCenterId: "US-NJ",
});

const myRandomPod = new runpod.Pod("myRandomPod", {
  cloudType: "ALL",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    (networkStorage) => networkStorage.id
  ),
  templateId: myTemplate.id,
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
  env: [
    {
      key: "JUPYTER_PASSWORD",
      value: "rns1hunbsstltcpad22d",
    },
  ],
});

export const pod = {
  value: myRandomPod.pod,
};

export const networkStorage = {
  value: testNetworkStorage.networkStorage,
};

export const template = {
  value: myTemplate,
};
