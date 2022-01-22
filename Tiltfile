# -*- mode: Python -*-
config.define_string("namespace")
cfg = config.parse()

version_settings(constraint='>=0.23.1')

allow_k8s_contexts('playground')

load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/main ./cmd/enable-it'

local_resource(
  'compile',
  compile_cmd,
  deps=['./cmd', './internal']
)

docker_build_with_restart(
  'example',
  '.',
  dockerfile='Dockerfile',
  entrypoint='/app/main',
  live_update=[
    sync('./main', 'main'),
  ],
)


k8s_yaml(helm('deploy/', values='values.yaml'))

k8s_resource(workload='example', port_forwards=8080, resource_deps=['compile'])