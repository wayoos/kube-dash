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

### Examples

### TODO

- In tools.go, fix runCmd() function. Executing `runCmd("kubectl")`, `runCmd("pluto version")` work fine. `runCmd("kubectl version")` times out.
