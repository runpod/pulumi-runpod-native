import pulumi
import runpodinfra as runpod
from loguru import logger
import json

def fetch_id(a):
    if type(a) == runpod.outputs.NetworkStorage:
        return a.id
    else:
        return a

try:
    test_network_storage = runpod.NetworkStorage("testNetworkStorage",
        name="testStorage1",
        size=5,
        data_center_id="EU-RO-1")

    my_random_pod = runpod.Pod("myRandomPod",
        cloud_type="ALL",
        network_volume_id=test_network_storage.network_storage.apply(lambda x : fetch_id(x)),
        gpu_count=1,
        volume_in_gb=50,
        container_disk_in_gb=50,
        min_vcpu_count=2,
        min_memory_in_gb=15,
        gpu_type_id="NVIDIA GeForce RTX 4090",
        name="RunPod Pytorch",
        image_name="runpod/pytorch",
        docker_args="",
        ports="8888/http",
        volume_mount_path="/workspace",
        env=[runpod.PodEnvArgs(
            key="JUPYTER_PASSWORD",
            value="rns1hunbsstltcpad22d",
        )])

    my_random_template = runpod.Template("myRandomTemplate",
        container_disk_in_gb = 5,
        container_registry_auth_id = "",
        docker_args="python handler.py",
        env=[{"key": "hi", "value": "hello"}],
        image_name="runpod/serverless-hello-world:latest",
        is_public=False,
        is_serverless=False,
        name="Generated Serverless Template",
        ports="1293/http",
        readme="Some readme", # pass some value to this. Won't work otherwise
        start_jupyter=False,
        start_ssh=False,
        volume_in_gb=20,
        volume_mount_path="/workspace",
    )

    pulumi.export("pod", {
        "value": my_random_pod.id,
    })
    pulumi.export("networkStorage", {
        "value": test_network_storage.network_storage,
    })
    pulumi.export("testTemplate", {
        "value": my_random_template.id,
    })
except Exception as e:
    logger.exception(e)