// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RunpodInfra.Runpod.Outputs
{

    [OutputType]
    public sealed class Pod
    {
        public readonly double AdjustedCostPerHr;
        public readonly string AiApiId;
        public readonly string ApiKey;
        public readonly string ConsumerUserId;
        public readonly int ContainerDiskInGb;
        public readonly string ContainerRegistryAuthId;
        public readonly double CostMultiplier;
        public readonly double CostPerHr;
        public readonly string CreatedAt;
        public readonly string DesiredStatus;
        public readonly string DockerArgs;
        public readonly string DockerId;
        public readonly ImmutableArray<string> Env;
        public readonly int GpuCount;
        public readonly int GpuPowerLimitPercent;
        public readonly ImmutableArray<Outputs.Gpu> Gpus;
        public readonly string Id;
        public readonly string ImageName;
        public readonly string LastStartedAt;
        public readonly string LastStatusChange;
        public readonly bool Locked;
        public readonly string MachineId;
        public readonly double MemoryInGb;
        public readonly string Name;
        public readonly string PodType;
        public readonly int Port;
        public readonly string Ports;
        public readonly Outputs.PodRegistry Registry;
        public readonly string TemplateId;
        public readonly int UptimeSeconds;
        public readonly double VcpuCount;
        public readonly int Version;
        public readonly bool VolumeEncrypted;
        public readonly double VolumeInGb;
        public readonly string VolumeKey;
        public readonly string VolumeMountPath;

        [OutputConstructor]
        private Pod(
            double adjustedCostPerHr,

            string aiApiId,

            string apiKey,

            string consumerUserId,

            int containerDiskInGb,

            string containerRegistryAuthId,

            double costMultiplier,

            double costPerHr,

            string createdAt,

            string desiredStatus,

            string dockerArgs,

            string dockerId,

            ImmutableArray<string> env,

            int gpuCount,

            int gpuPowerLimitPercent,

            ImmutableArray<Outputs.Gpu> gpus,

            string id,

            string imageName,

            string lastStartedAt,

            string lastStatusChange,

            bool locked,

            string machineId,

            double memoryInGb,

            string name,

            string podType,

            int port,

            string ports,

            Outputs.PodRegistry registry,

            string templateId,

            int uptimeSeconds,

            double vcpuCount,

            int version,

            bool volumeEncrypted,

            double volumeInGb,

            string volumeKey,

            string volumeMountPath)
        {
            AdjustedCostPerHr = adjustedCostPerHr;
            AiApiId = aiApiId;
            ApiKey = apiKey;
            ConsumerUserId = consumerUserId;
            ContainerDiskInGb = containerDiskInGb;
            ContainerRegistryAuthId = containerRegistryAuthId;
            CostMultiplier = costMultiplier;
            CostPerHr = costPerHr;
            CreatedAt = createdAt;
            DesiredStatus = desiredStatus;
            DockerArgs = dockerArgs;
            DockerId = dockerId;
            Env = env;
            GpuCount = gpuCount;
            GpuPowerLimitPercent = gpuPowerLimitPercent;
            Gpus = gpus;
            Id = id;
            ImageName = imageName;
            LastStartedAt = lastStartedAt;
            LastStatusChange = lastStatusChange;
            Locked = locked;
            MachineId = machineId;
            MemoryInGb = memoryInGb;
            Name = name;
            PodType = podType;
            Port = port;
            Ports = ports;
            Registry = registry;
            TemplateId = templateId;
            UptimeSeconds = uptimeSeconds;
            VcpuCount = vcpuCount;
            Version = version;
            VolumeEncrypted = volumeEncrypted;
            VolumeInGb = volumeInGb;
            VolumeKey = volumeKey;
            VolumeMountPath = volumeMountPath;
        }
    }
}