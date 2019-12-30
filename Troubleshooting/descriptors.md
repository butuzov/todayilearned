# Way to many open file descriptors.

* "-bash: fork: Resource temporarily unavailable"
  https://apple.stackexchange.com/questions/373035

## Workflow

```
# current limit
ulimit -u

# currently (process of current user)
ps -xu $(id -ru) | wc -l

# killing zomby process
ps -xu $(id -ru) | grep zmb | awk '{print $2}' | \
  xargs -I {} sh -c "kill -9 {}"
```