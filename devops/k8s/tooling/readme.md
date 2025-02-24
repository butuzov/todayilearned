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

- `UI` [portfall](https://github.com/Rested/portfall) `Last Releaased` @ May 21, 2020

## Highlighting YAML

- [`yh`](https://github.com/andreazorzetto/yh) - YAML Highlighter
  `kubectl get pod alpine -o yaml | yh`
- `yq` - `kubectl get pod alpine -o yaml | yq`
- [`bat`](https://github.com/sharkdp/bat)
  `cat config.yaml | bat -l=yaml`
