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

const myRandomPod = new runpod.Pod("myRandomPod", {
  cloudType: "ALL",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    // @ts-ignore
    (networkStorage) => networkStorage.id
  ),
  gpuCount: 1,
  volumeInGb: 50,
  containerDiskInGb: 50,
  minVcpuCount: 2,
  minMemoryInGb: 15,
  gpuTypeId: "NVIDIA GeForce RTX 4090",
  name: "RunPod Pytorch",
  imageName: "runpod/pytorch:latest",
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

const myRandomEndpoint = new runpod.Endpoint("myRandomEndpoint", {
  gpuIds: "AMPERE_16,AMPERE_24,-NVIDIA L4",
  idleTimeout: 100,
  locations: "CA-MTL-2,CA-MTL-3,EU-RO-1,US-CA-1,US-GA-1,US-KS-2,US-OR-1,CA-MTL-1,US-TX-3,EUR-IS-1,EUR-IS-2,SEA-SG-1",
  name: "myRandomEndpoint",
  networkVolumeId: testNetworkStorage.networkStorage.apply(
    // @ts-ignore
    (networkStorage) => networkStorage.id
  ),
  scalerType: 'REQUEST_COUNT',
  scalerValue: 2,
  templateId: myTemplate.template.apply(t => t.id),
  workersMax: 2,
  workersMin: 1,
})

export const template = {
  value: myTemplate.template,
};

export const endpoint = {
  value: myRandomEndpoint.endpoint,
};

export const pod = {
  value: myRandomPod.pod,
};

export const networkStorage = {
  value: testNetworkStorage.networkStorage,
};
