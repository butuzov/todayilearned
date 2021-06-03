# CKA Notes

## Exam Experience

* https://habr.com/ru/company/flant/blog/425683/
* https://blog.heptio.com/93d20af32557

## Resources

* https://github.com/krzko/awesome-cka
* Exam Overview + References: [walidshaari/Kubernetes-Certified-Administrator](https://github.com/walidshaari/Kubernetes-Certified-Administrator)
* https://unofficial-kubernetes.readthedocs.io/en/latest
* https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands
* https://kubernetes.io/docs/reference/kubectl/cheatsheet/
* (good) https://unofficialism.info/posts/unofficial-tips-for-cka-and-ckad-exams/

## Practice Excercies

* (vagrant) https://github.com/jecnua/certified-kubernetes-administrator-cka-kata
* (good-1) https://github.com/stretchcloud/cka-lab-practice
* (good-2) https://github.com/wonkarthik/k8s-questions
* (good-3) https://github.com/dgkanatsios/CKAD-exercises

* https://github.com/arush-sal/cka-practice-environment 
    * https://github.com/kubernauts/cka-practice-environment (`docker-compose`)
    

## Tips and Tricks

* `k explain <pod>.<spec>`
* Read - *Kubernetes the Hard Way*
* `alias k=kubectl`
* `complete -F __start_kubectl k`
* Watch and track **time**
* Watch and track **questions**
* No Copy/Paste, **TYPE**

### Getting Help

* https://kubernetes.io/docs/
* https://kubernetes.io/docs/reference/kubectl/cheatsheet/
* or ask kubectl itself `kubectl explain <resource>.<key>`

### Exam Topic - Core Concepts (Config) - `19%`

```bash
kubectl config view --minify
```

### Exam Topics - Installation, Configuration and Validation `12%`

### Exam Topics - Security `12%`

```bash
k get secret(s)
k describe secrets default-token-59rrj
k get secrets default-token-59rrj -o json | jq --row-output '.data"ca.crt"'
```

### Exam Topics - Networking `11%`

* [CIRD change](https://capstonec.com/help-i-need-to-change-the-pod-cidr-in-my-kubernetes-cluster/)

```bash
# route to pod
k port-forward <pod> port
curl localhost:port
```

### Exam Topics - Cluster Maintenance `11%`

### Exam Topics - Troubleshooting `10%`

### Exam Topics - Storage `7%`

### Exam Topics - Application Lifecycle Management `8%`

<div class="alert alert-warning"><center>Practice Auto Scaling (vagrant)</center></div>

https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#scaling-a-deployment

```bash
# show the revisions history
k rollout history deployment <deploymentName>  
# show pod tempalte from the revision 2
k rollout history deployment <deploymentName> --revision=2
# check thee status
k rollout status deployment <deploymentName>
# undo previous deployment - will change only pod, not deployment stategy
k rollout undo deployment <deploymentName>
# undo to revision #2
k rollout undo deployment <deploymentName> --to-revision=2
```

### Exam Topics - Scheduling `5%`

### Exam Topics - Monitoring and Logging - `5%`

```bash
# after monitor resources deployed
k top nodes
k top pods

# check logs
k logs <pod>

# container in pod 
k logs <pod> <container>
k logs -f -c <container> <pod>
k logs deployment/<deployment> -c <container>
```

## Exam

* `kubeadm`
* `systemd`

### Real Exam Questions: Set 1

https://docs.google.com/document/d/1AMVwvVabPoYt-o1k8Uo7UlmlfsjQKVHDhDyKP3QqbOM/edit


1. Create a node that has a SSD and label it as such. 
  * Create a pod that is only scheduled on SSD nodes.
1. Create 2 pod definitions: 
  * the second pod should be scheduled to run anywhere the first pod is running 
  * 2nd pod runs alongside the first pod
  
1. Create a deployment running nginx version 1.12.2 that will run in 2 pods
  * Scale this to 4 pods.
  * Scale it back to 2 pods.
  * Upgrade this to 1.13.8
  * Check the status of the upgrade
  * How do you do this in a way that you can see history of what happened?
  * Undo the upgrade
1. Create a service that uses a scratch disk.
  * Change the service to mount a disk from the host.
  * Change the service to mount a persistent volume.
1. Create a pod that has a liveness check
1. Create a service that manually requires endpoint creation - and create that too
1. Create a daemon set
   * Change the update strategy to do a rolling update but delaying 30 seconds between pod updates
1. Create a static pod
1. Create a busybox container without a manifest. Then edit the manifest.
1. Create a pod that uses secrets
	a. Pull secrets from environment variables
	b. Pull secrets from a volume
	c. Dump the secrets out via kubectl to show it worked
1. Create a job that runs every 3 minutes and prints out the current time.
1. Create a job that runs 20 times, 5 containers at a time, and prints "Hello parallel world"
1. Create a service that uses an external load balancer and points to a 3 pod cluster running nginx.
1. Create a horizontal autoscaling group that starts with 2 pods and scales when CPU usage is over 50%.
1. Create a custom resource definition
	a. Display it in the API with curl
1. Create a networking policy such that only pods with the label access=granted can talk to it.
  * Create an nginx pod and attach this policy to it. 
  * Create a busybox pod and attempt to talk to nginx - should be blocked
  * Attach the label to busybox and try again - should be allowed
1. Create a service that references an externalname.
  * Test that this works from another pod
1. Create a pod that runs all processes as user 1000.
1. Create a namespace
  * Run a pod in the new namespace
  * Put memory limits on the namespace
  * Limit pods to 2 persistent volumes in this namespace
1. Write an ingress rule that redirects calls to /foo to one service and to /bar to another
1. Write a service that exposes nginx on a nodeport
	* Change it to use a cluster port
	* Scale the service
	* Change it to use an external IP
	* Change it to use a load balancer
1. Deploy nginx with 3 replicas and then expose a port
	a. Use port forwarding to talk to a specific port
    
1. Make an API call using CURL and proper certs
1. Upgrade a cluster with kubeadm
1. Get logs for a pod
1. Deploy a pod with the wrong image name (like --image=nginy) and find the error message.
1. Get logs for kubectl
1. Get logs for the scheduler
1. Restart kubelet
   ```bash
   systemctl stop kubelet.service
   systemctl start kubelet.service
   # or   
   systemctl restart kubelet.service
   ```

**Non-K8S**
1. Convert a CRT to a PEM
	a. Convert it back
1. Backup an etcd cluster
1. List the members of an etcd cluster
1. Find the health of etcd


### Questions (3) - Set of 3 Questions

1. сертификаты протухли вчера
2. сделать снапшот etcd и откатится до него на 3ъ машинах
3. есть кластер из 3ъ машин, 2 умерли, привезли 2 новых - восстановить кластер до 3х машин

# Problems Already

* metrics-server deployment - `CreateContainerConfigError` [link](https://stackoverflow.com/questions/50424754/pod-status-as-createcontainerconfigerror-in-minikube-cluster) + `Error: container has runAsNonRoot and image will run as root`