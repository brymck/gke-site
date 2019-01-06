import os
import yaml

proxies = {}
for var_name in ['http_proxy', 'https_proxy', 'HTTP_PROXY', 'HTTPS_PROXY']:
    if var_name in os.environ:
        proxies[var_name] = os.environ[var_name]

doc = yaml.load(open('skaffold.yaml', 'r'))
for artifact in doc['build']['artifacts']:
    if 'docker' not in artifact:
        artifact['docker'] = {}
    if 'buildArgs' not in artifact['docker']:
        artifact['docker']['buildArgs'] = {}
    build_args = artifact['docker']['buildArgs']
    for key, value in proxies.items():
        build_args[key] = value

stream = open('skaffold-with-proxies.yaml', 'w')
yaml.dump(doc, stream, default_flow_style=False)
