# Git


<!-- TOC depthfrom:2 depthto:4 -->

- [Reading List](#reading-list)
- [Branching](#branching)
    - [Creating Orphaned branches](#creating-orphaned-branches)
- [Forks](#forks)
    - [How do I update a GitHub forked repository? stackoverflow.com/q/7244321](#how-do-i-update-a-github-forked-repository-stackoverflowcomq7244321)
- [History](#history)
    - [How to see what happened to the file in path of the repo?](#how-to-see-what-happened-to-the-file-in-path-of-the-repo)
- [Rebase](#rebase)
    - [Rebase Changes from upstream into your branch](#rebase-changes-from-upstream-into-your-branch)

<!-- /TOC -->
## Reading List

* [Start Using Git on the Command Line Like a Pro in Five Minutes](https://medium.com/better-programming/start-using-git-on-the-command-line-like-a-pro-in-5-minutes-36a6e0007e9f)
* [Git CLI tips and tricks to be more productive](https://www.codementor.io/@kpunith8/git-cli-tips-and-tricks-to-be-more-productive-1a3pb4fyvn)
* https://github.com/tj/git-extras

## Branching
### [Creating Orphaned branches](https://stackoverflow.com/questions/1384325)

```bash
git checkout --orphan gh-pages
git rm -rf .
<do work>
git add your files
git commit -m 'Initial commit'
```

## Forks
### How do I update a GitHub forked repository? ([`stackoverflow.com/q/7244321`](https://stackoverflow.com/q/7244321/))

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

## History
### How to see what happened to the file in path of the repo?

```bash
git log --full-history -- content/docs/examples/
```

## Rebase

### Rebase Changes from upstream into your branch

```
git checkout target
git rebase master
```

