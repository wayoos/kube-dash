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

- When exec `docker-compose up`, backend can't get a connection with frontend. "Proxy error: Could not proxy request /api/namespaces from localhost:8080 to http://localhost:8008."
- Backend must be ran with --network=host flag, otherwise can't establish connection with cluster. Might come from the kubeconfig.yaml config

