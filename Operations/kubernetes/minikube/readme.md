## minikube

![](minikube.jpg)


```bash
# notes on https://github.com/kubernetes/minikube/issues/567

# 8 Gi of memory
minikube config set memory 8192
# 4 CPU
minikube config set cpus 4
# virtualbox
VBoxManage modifyvm "minikube" --cpus 4 --memory 8196
# starting
minikube start --memory=8196 --cpus=4 --vm-driver=virtualbox
minikube start --memory=8196 --cpus=4 --vm-driver=virtualbox --mount-string "$(pwd):/data"



# prior to 1.4 (pickikng kube version)
minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.13.0 \
  --disk-size=30g \
  --extra-config=apiserver.enable-admission-plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"

# prior to 1.4 (vmware fusion)
minikube start --cpus=8 --memory=32768 --disk-size=50g --vm-driver=vmwarefusion
```

## Usage

### SShing into minikube

```bash
# ssh to minikube
minikube ssh

# or (for copy)
MINIKUBE_IP=`minikube ip`
ssh -o UserKnownHostsFile=/dev/null \
    -o StrictHostKeyChecking=no -o LogLevel=quiet \
    -i ~/.minikube/machines/minikube/id_rsa docker@$MINIKUBE_IP
```


### Loading Docker images

```bash
# copy docker images
docker save <image> | pv | (eval $(minikube docker-env) && docker load)

# copy but from files.
while read image; do
 docker save $image | pv | (eval $(minikube docker-env) && docker load)
done <images.txt;
```

or using insecure registry

```
cat images.txt | xargs -I {} sh -c "docker pull localhost:5000/{}"
```

### Build from source (bew key)

```bash
brew reinstall --build-from-source minikube
```

### Kubernetes

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

# mount folder (issue on 1.13.6)
minikube mount /host-path:/vm-path

# addons
minikube addons list

# dashboard
minikube dashboard --url &
```

### Multinode support
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

### `--help`

`minikube --help`

Minikube is a CLI tool that provisions and manages single-node Kubernetes clusters optimized for development workflows.

#### Basic Commands:
* `start`          Starts a local kubernetes cluster
* `status`         Gets the status of a local kubernetes cluster
* `stop`           Stops a running local kubernetes cluster
* `delete`         Deletes a local kubernetes cluster
* `dashboard`      Access the kubernetes dashboard running within the minikube cluster
* `pause`          pause containers
* `unpause`        unpause Kubernetes

#### Images Commands:
* `docker-env`     Sets up docker env variables; similar to '$(docker-machine env)'
* `podman-env`     Sets up podman env variables; similar to '$(podman-machine env)'
* `cache`          Add or delete an image from the local cache.

#### Configuration and Management Commands:
* `addons`         Modify minikube's kubernetes addons
* `config`         Modify minikube config
* `profile`        Profile gets or sets the current minikube profile
* `update-context` Verify the IP address of the running cluster in kubeconfig.

#### Networking and Connectivity Commands:
* `service`        Gets the kubernetes URL(s) for the specified service in your local cluster
* `tunnel`         tunnel makes services of type LoadBalancer accessible on localhost

#### Advanced Commands:
* `mount`          Mounts the specified directory into minikube
* `ssh`            Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'
* `kubectl`        Run kubectl
* `node`           Node operations

#### Troubleshooting Commands:
* `ssh-key`        Retrieve the ssh identity key path of the specified cluster
* `ip`             Retrieves the IP address of the running cluster
* `logs`           Gets the logs of the running instance, used for debugging minikube, not user code.
* `update-check`   Print current and latest version number
* `version`        Print the version of minikube
* `options`        Show a list of global command-line options (applies to all commands).

#### Other Commands:
* `completion`     Outputs minikube shell completion for the given shell (bash or zsh)
