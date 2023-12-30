<!-- tags: package manager, k8s -->

# helm

Kubernetes Package manager

## Recipes


```shell
> helm pull dagster/dagster & ls
dagster-1.5.4.tgz
> helm pull dagster/dagster --version=1.5.5
Error: chart "dagster" matching 1.5.5 not found in dagster index. (try 'helm repo update'): no chart version found for dagster-1.5.5
> helm repo update
```
