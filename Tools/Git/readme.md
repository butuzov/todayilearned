# Git

## Reading List

* [Start Using Git on the Command Line Like a Pro in Five Minutes](https://medium.com/better-programming/start-using-git-on-the-command-line-like-a-pro-in-5-minutes-36a6e0007e9f)

---

## [Creating Orphaned branches](https://stackoverflow.com/questions/1384325)

```bash
git checkout --orphan gh-pages
git rm -rf .
<do work>
git add your files
git commit -m 'Initial commit'
```

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

