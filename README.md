Kube-dash
=========
Kubernetes dashboard.

Frontend
--------

Create a new project in the directory front:

npx @vue/cli create frontend

$HOME/.npm-packages/bin/vue add bootstrap-vue bootstrap

$HOME/.npm-packages/bin/vue add router

$HOME/.npm-packages/bin/vue add pug

$HOME/.npm-packages/bin/vue add axios

### Build

`docker build --tag kubedash:latest --build-arg KUBE_CONFIG=./kubeconfig.yaml .; docker run --network="host" --publish 8008:8008 kubedash`

### TODO


