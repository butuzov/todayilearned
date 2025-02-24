<!-- weight: 10 -->
<!-- menu: Local Cluster -->
<!-- seotitle: Running Local Kubernetes CLuster with minikube, kind, etc.. -->

# Running Local Kubernetes Cluster

## microk8s

- [Homepage](https://microk8s.io/)
- [Tutorials](https://microk8s.io/docs/tutorials)

```shell
# Install microk8s & microk8s's cluster
brew install ubuntu/microk8s/microk8s
microk8s install

# Check the status while Kubernetes starts
microk8s status --wait-ready

# Turn on the services you want
# https://microk8s.io/docs/addons
microk8s enable dashboard
microk8s enable dns
microk8s enable registry
microk8s enable istio

microk8s kubectl get all --all-namespaces

# switch to kubectl
microk8s config > ~/.kube/config
```

## Kind

`kind is a tool for running local Kubernetes clusters using Docker container “nodes”. kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.

- [Homepage](https://kind.sigs.k8s.io/)

```shell
brew install kind

# Crate Cluster
kind create cluster
kind create cluster --image kindest/node:latest # Specific Node
kind create cluster --name dev                  # Specific Name


# Delete Cluster
kind delete cluster
```

## Minikube

```shell
brew install minikube

# Customizing Run
minikube config set memory 8192                       # 8 Gi of memory
minikube config set cpus 4                            # 4 CPU
VBoxManage modifyvm "minikube" --cpus 4 --memory 8196 # virtualbox

# starting
minikube start --memory=8196 --cpus=4 --vm-driver=virtualbox
minikube start --memory=8196 --cpus=4 --vm-driver=virtualbox --mount-string "$(pwd):/data"

# LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount
# prior to 1.4 (pickikng kube version)
plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount"
plugins="${plugins},DefaultStorageClass,MutatingAdmissionWebhook"
minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.13.0 \
  --disk-size=30g \
  --extra-config=apiserver.enable-admission-plugins="$plugins"

# prior to 1.4 (vmware fusion)
minikube start --cpus=8 --memory=32768 --disk-size=50g --vm-driver=vmwarefusion
```

```help
# Basic Commands

- start           Starts a local kubernetes cluster
- status          Gets the status of a local kubernetes cluster
- stop            Stops a running local kubernetes cluster
- delete          Deletes a local kubernetes cluster
- dashboard       Access the kubernetes dashboard running within the minikube cluster
- pause           pause containers
- unpause         unpause Kubernetes

# Images Commands

- docker-env      Sets up docker env variables; similar to '$(docker-machine env)'
- podman-env      Sets up podman env variables; similar to '$(podman-machine env)'
- cache           Add or delete an image from the local cache.

# Configuration and Management Commands

- addons          Modify minikube's kubernetes addons
- config          Modify minikube config
- profile         Profile gets or sets the current minikube profile
- update-context  Verify the IP address of the running cluster in kubeconfig.

# Networking and Connectivity Commands

- service         Gets the kubernetes URL(s) for the specified service in your local cluster
- tunnel          tunnel makes services of type LoadBalancer accessible on localhost

# Advanced Commands

- mount           Mounts the specified directory into minikube
- ssh             Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'
- kubectl         Run kubectl
- node            Node operations

# Troubleshooting Commands

- ssh-key         Retrieve the ssh identity key path of the specified cluster
- ip              Retrieves the IP address of the running cluster
- logs            Gets the logs of the running instance, used for debugging minikube, not user code.
- update-check    Print current and latest version number
- version         Print the version of minikube
- options         Show a list of global command-line options (applies to all commands).

# Other Commands

completion     Outputs minikube shell completion for the given shell (bash or zsh)
```

### Recipe: `SSH`ing into minikube

```shell
# ssh to minikube
minikube ssh

# or (for copy)
MINIKUBE_IP=`minikube ip`
ssh -o UserKnownHostsFile=/dev/null \
    -o StrictHostKeyChecking=no -o LogLevel=quiet \
    -i ~/.minikube/machines/minikube/id_rsa docker@$MINIKUBE_IP
```

### Recipe: Loading Docker images

```bash
# copy docker images
docker save <image> | pv | (eval $(minikube docker-env) && docker load)

# copy but from files.
while read image; do
 docker save $image | pv | (eval $(minikube docker-env) && docker load)
done <images.txt;
```

or using insecure registry

```shell
cat images.txt | xargs -I {} sh -c "docker pull localhost:5000/{}"
```

### Recipe: Multinode support

```bash
> minikube node --help
Operations on nodes

Available Commands:
  add         Adds a node to the given cluster.
  delete      Deletes a node from a cluster.
  start       Starts a node.
  stop        Stops a node in a cluster.

> minikube node add
> minikube status
```

### Extras

in addition to kubernetes functionality minikube has

```bash
# service info
minikube service <srv-name>

# ip
minikube ip

# status
minikube status

# logs
minikube logs

# cluster
kubectl cluster-info

# mount folder
minikube mount /host-path:/vm-path

# addons
minikube addons list

# dashboard
minikube dashboard --url &
```
