# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "v1.8.9"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "runpod Pulumi Package - Development Version"


setup(name='runpodinfra',
      python_requires='>=3.8',
      version=VERSION,
      description="The Runpod Pulumi provider provides resources to interact with Runpod's native APIs.",
      long_description=readme(),
      long_description_content_type='text/markdown',
      keywords='pulumi runpod gpus ml ai',
      url='https://runpod.io',
      project_urls={
          'Repository': 'https://github.com/runpod/pulumi-runpod-native'
      },
      license='Apache-2.0',
      packages=find_packages(),
      package_data={
          'runpodinfra': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
