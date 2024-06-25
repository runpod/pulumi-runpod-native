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
  templateId: myTemplate.template.apply((template) => {
    console.log("template.id", template);
    return template.id;
  }),
  // templateId: testNetworkStorage.networkStorage.apply((networkStorage) => {
  //   console.log("networkStorage.id", networkStorage.id);
  //   return "sda89dh4i9"  // networkStorage.id;
  // }),
  // templateId: "sda89dh4i9", // myTemplate.template.id,
  workersMax: 2,
  workersMin: 1,
  idleTimeout: 6,
  locations: "US-OR-1",
  networkVolumeId: testNetworkStorage.networkStorage.apply((networkStorage) => {
    console.log("networkStorage.id", networkStorage.id);
    return networkStorage.id;
  }),
  // networkVolumeId: testNetworkStorage.networkStorage.id,
  scalerType: "QUEUE_DELAY",
  scalerValue: 4,
});

export const template = {
  value: myTemplate.template,
};

export const endpoint = {
  value: myEndpoint.endpoint,
};

export const networkStorage = {
  value: testNetworkStorage.networkStorage,
};
