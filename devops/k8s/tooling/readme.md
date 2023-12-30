<!-- menu: Tooling -->
# Kubernetes Tooling

## krew

[Krew](https://krew.sigs.k8s.io/) is plugin manager for kubectl.

```shell
$ kubectl krew update
$ kubectl krew search
$ kubectl krew install access-matrix
$ kubectl krew upgrade
$ kubectl krew uninstall access-matrix
```

### Plugins

- https://github.com/stern/stern tailf for containers
- https://github.com/alcideio/rbac-tool rbac


## Port forwarding

- `UI` [portfall](https://github.com/Rested/portfall) ` Last Releaased ` @ May 21, 2020