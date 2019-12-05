# Git


## How do I update a GitHub forked repository? ([`stackoverflow.com/q/7244321`](https://stackoverflow.com/q/7244321/))

```bash
# fork https://github.com/kubernetes/website
git clone https://github.com/butuzov/kubernetes-website
cd kubernetes-website/

# Add the remote, call it "upstream":
git remote add upstream https://github.com/kubernetes/website.git

# get updates
git fetch upstream

# merge
git merge upstream/master master

# push
git push -f origin master
```


## How to see what happend to file in path of the repo?

```bash
git log --full-history -- content/docs/examples/
```

## Empty Branch
