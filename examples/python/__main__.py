import pulumi
import runpodinfra as runpod
from loguru import logger
import json

def fetch_id(a):
    if type(a) == runpod.outputs.NetworkStorage:
        return a.id
    else:
        return a

def fetch_template_id(a):
    if type(a) == runpod.outputs.Template:
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
        is_serverless=True,
        name="Generated Serverless Template",
        ports="1293/http",
        readme="Some readme", # pass some value to this. Won't work otherwise
        start_jupyter=False,
        start_ssh=False,
        volume_in_gb=20,
        volume_mount_path="/workspace",
    )

    my_random_endpoint = runpod.Endpoint("myRandomEndpoint",
        gpu_ids="AMPERE_16,AMPERE_24,-NVIDIA L4",
        idle_timeout=100,
        locations="CA-MTL-2,CA-MTL-3,EU-RO-1,US-CA-1,US-GA-1,US-KS-2,US-OR-1,CA-MTL-1,US-TX-3,EUR-IS-1,EUR-IS-2,SEA-SG-1",
        name="myRandomEndpoint",
        network_volume_id=test_network_storage.network_storage.apply(lambda x : fetch_id(x)),
        scaler_type='REQUEST_COUNT',
        scaler_value=2,
        template_id=my_random_template.template.apply(lambda x : fetch_template_id(x)),
        workers_max=2,
        workers_min=1,
    )

    pulumi.export("pod", {
        "value": my_random_pod.pod,
    })
    pulumi.export("networkStorage", {
        "value": test_network_storage.network_storage,
    })
    pulumi.export("testTemplate", {
        "value": my_random_template.template,
    })
    pulumi.export("testEndpoint", {
        "value": my_random_endpoint.endpoint,
    })
except Exception as e:
    logger.exception(e)