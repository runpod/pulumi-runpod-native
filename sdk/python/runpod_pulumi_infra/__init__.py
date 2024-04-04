# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .network_storage import *
from .pod import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import runpod_pulumi_infra.config as __config
    config = __config
else:
    config = _utilities.lazy_import('runpod_pulumi_infra.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "runpod",
  "mod": "index",
  "fqn": "runpod_pulumi_infra",
  "classes": {
   "runpod:index:NetworkStorage": "NetworkStorage",
   "runpod:index:Pod": "Pod"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "runpod",
  "token": "pulumi:providers:runpod",
  "fqn": "runpod_pulumi_infra",
  "class": "Provider"
 }
]
"""
)
