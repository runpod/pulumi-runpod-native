# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 gpu_ids: pulumi.Input[str],
                 name: pulumi.Input[str],
                 template_id: pulumi.Input[str],
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 locations: Optional[pulumi.Input[str]] = None,
                 network_volume_id: Optional[pulumi.Input[str]] = None,
                 scaler_type: Optional[pulumi.Input[str]] = None,
                 scaler_value: Optional[pulumi.Input[int]] = None,
                 workers_max: Optional[pulumi.Input[int]] = None,
                 workers_min: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        """
        pulumi.set(__self__, "gpu_ids", gpu_ids)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "template_id", template_id)
        if idle_timeout is not None:
            pulumi.set(__self__, "idle_timeout", idle_timeout)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if network_volume_id is not None:
            pulumi.set(__self__, "network_volume_id", network_volume_id)
        if scaler_type is not None:
            pulumi.set(__self__, "scaler_type", scaler_type)
        if scaler_value is not None:
            pulumi.set(__self__, "scaler_value", scaler_value)
        if workers_max is not None:
            pulumi.set(__self__, "workers_max", workers_max)
        if workers_min is not None:
            pulumi.set(__self__, "workers_min", workers_min)

    @property
    @pulumi.getter(name="gpuIds")
    def gpu_ids(self) -> pulumi.Input[str]:
        return pulumi.get(self, "gpu_ids")

    @gpu_ids.setter
    def gpu_ids(self, value: pulumi.Input[str]):
        pulumi.set(self, "gpu_ids", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "idle_timeout")

    @idle_timeout.setter
    def idle_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_timeout", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter(name="networkVolumeId")
    def network_volume_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_volume_id")

    @network_volume_id.setter
    def network_volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_volume_id", value)

    @property
    @pulumi.getter(name="scalerType")
    def scaler_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scaler_type")

    @scaler_type.setter
    def scaler_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaler_type", value)

    @property
    @pulumi.getter(name="scalerValue")
    def scaler_value(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "scaler_value")

    @scaler_value.setter
    def scaler_value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "scaler_value", value)

    @property
    @pulumi.getter(name="workersMax")
    def workers_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "workers_max")

    @workers_max.setter
    def workers_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "workers_max", value)

    @property
    @pulumi.getter(name="workersMin")
    def workers_min(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "workers_min")

    @workers_min.setter
    def workers_min(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "workers_min", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gpu_ids: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 locations: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_volume_id: Optional[pulumi.Input[str]] = None,
                 scaler_type: Optional[pulumi.Input[str]] = None,
                 scaler_value: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 workers_max: Optional[pulumi.Input[int]] = None,
                 workers_min: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Endpoint resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Endpoint resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gpu_ids: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 locations: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_volume_id: Optional[pulumi.Input[str]] = None,
                 scaler_type: Optional[pulumi.Input[str]] = None,
                 scaler_value: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 workers_max: Optional[pulumi.Input[int]] = None,
                 workers_min: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointArgs.__new__(EndpointArgs)

            if gpu_ids is None and not opts.urn:
                raise TypeError("Missing required property 'gpu_ids'")
            __props__.__dict__["gpu_ids"] = gpu_ids
            __props__.__dict__["idle_timeout"] = idle_timeout
            __props__.__dict__["locations"] = locations
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["network_volume_id"] = network_volume_id
            __props__.__dict__["scaler_type"] = scaler_type
            __props__.__dict__["scaler_value"] = scaler_value
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["workers_max"] = workers_max
            __props__.__dict__["workers_min"] = workers_min
            __props__.__dict__["endpoint"] = None
        super(Endpoint, __self__).__init__(
            'runpod:index:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointArgs.__new__(EndpointArgs)

        __props__.__dict__["endpoint"] = None
        __props__.__dict__["gpu_ids"] = None
        __props__.__dict__["idle_timeout"] = None
        __props__.__dict__["locations"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_volume_id"] = None
        __props__.__dict__["scaler_type"] = None
        __props__.__dict__["scaler_value"] = None
        __props__.__dict__["template_id"] = None
        __props__.__dict__["workers_max"] = None
        __props__.__dict__["workers_min"] = None
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output['outputs.Endpoint']:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="gpuIds")
    def gpu_ids(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gpu_ids")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkVolumeId")
    def network_volume_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "network_volume_id")

    @property
    @pulumi.getter(name="scalerType")
    def scaler_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "scaler_type")

    @property
    @pulumi.getter(name="scalerValue")
    def scaler_value(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "scaler_value")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="workersMax")
    def workers_max(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "workers_max")

    @property
    @pulumi.getter(name="workersMin")
    def workers_min(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "workers_min")

