<!-- weight: 2 -->

# Kubernetes

<div class="alert alert-primary" role="alert">
This isn't a notes on k8s (actually it is, but...), its a WHOLE ecosystem, its hard to reduce to simple readme.md
</div>

## Runing Local

### Docker Desktop

Read the [full Manual](https://docs.docker.com/desktop/kubernetes/).

- From the Docker Dashboard, select the Settings.
- Select Kubernetes from the left sidebar.
- Next to Enable Kubernetes, select the checkbox.
- Select Apply & Restart to save the settings and then select Install to confirm. This instantiates images required to run the Kubernetes server as containers, and installs the `/usr/local/bin/kubectl` command on your machine.

### Other Options

[minikube](local/readme.md#minikube), [kind](kind/readme.md#kind), [minishift](https://github.com/minishift/minishift), [CRC](https://github.com/crc-org/crc), ansible's [kubespray](https://github.com/kubernetes-sigs/kubespray) with Vagrant (e.g. [k8s-cluster](https://github.com/butuzov/k8s-cluster))

## Running Commands

### Namespaces

```
# set space
kubectl config set-context --current --namespace=ns1

# get namespace
kubectl config view --minify -o jsonpath='{..namespace}'  # json
kubectl config view | grep namespace                      # pipes
kubectl describe sa default | grep Namespace              # using sa
```

### List of Kubernetes API resources and subresources?

```shell
# https://stackoverflow.com/q/49396607
kubectl api-resources -o wide
```
