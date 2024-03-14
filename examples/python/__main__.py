import pulumi
import pulumi_runpod as runpod

test_network_storage = runpod.NetworkStorage("testNetworkStorage",
    name="testStorage1",
    size=20,
    data_center_id="US-NJ")
my_random_pod = runpod.Pod("myRandomPod",
    cloud_type="ALL",
    network_volume_id=test_network_storage.network_storage.id,
    gpu_count=1,
    volume_in_gb=50,
    container_disk_in_gb=50,
    min_vcpu_count=2,
    min_memory_in_gb=15,
    gpu_type_id="NVIDIA GeForce RTX 3070",
    name="RunPod Pytorch",
    image_name="runpod/pytorch",
    docker_args="",
    ports="8888/http",
    volume_mount_path="/workspace",
    env=[runpod.PodEnvArgs(
        key="JUPYTER_PASSWORD",
        value="rns1hunbsstltcpad22d",
    )])
pulumi.export("pod", {
    "value": my_random_pod.pod,
})
pulumi.export("networkStorage", {
    "value": test_network_storage.network_storage,
})
