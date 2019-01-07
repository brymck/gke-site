"""
Inject proxies into Skaffold YAML for development purposes so web requests from containers work.

To do this, we check the environment for proxy-related environment variables:
* http_proxy
* https_proxy
* all_proxy
* no_proxy

If any of those exist, we read in skaffold.yaml and add those to build.artifacts.docker.buildArgs.
We then save that file to skaffold.dev.yaml.

If they don't exist, we just directly copy skaffold.yaml to skaffold.dev.yaml.
"""

import os
import shutil
from urllib.parse import urlparse

import yaml

base_yaml = 'skaffold.yaml'
dev_yaml = 'skaffold.dev.yaml'


def get_proxies():
    """Determine proxies based on environment variables."""
    proxies = {}
    for var_name in ['http_proxy', 'https_proxy', 'all_proxy', 'no_proxy']:
        if var_name in os.environ:
            proxies[var_name] = os.environ[var_name]
        if var_name.upper() in os.environ:
            proxies[var_name.upper()] = os.environ[var_name]
    return proxies


def apply_proxies(proxies):
    """Read in skaffold.yaml, apply proxies as Docker --build-args, save as skaffold.dev.yaml."""
    doc = yaml.load(open(base_yaml, 'r'))
    for artifact in doc['build']['artifacts']:
        if 'docker' not in artifact:
            artifact['docker'] = {}
        if 'buildArgs' not in artifact['docker']:
            artifact['docker']['buildArgs'] = {}
        build_args = artifact['docker']['buildArgs']
        for key, value in proxies.items():
            build_args[key] = value

    stream = open(dev_yaml, 'w')
    yaml.dump(doc, stream, default_flow_style=False)


def main():
    """Entrypoint."""
    proxies = get_proxies()
    if proxies:
        apply_proxies(proxies)
    else:
        shutil.copyfile(base_yaml, dev_yaml)


main()
