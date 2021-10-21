# Git related things

## Git Ugly

![gitpretty.png](gitpretty.png)

## Reading List

* [Start Using Git on the Command Line Like a Pro in Five Minutes](https://medium.com/better-programming/start-using-git-on-the-command-line-like-a-pro-in-5-minutes-36a6e0007e9f)
* [Git CLI tips and tricks to be more productive](https://www.codementor.io/@kpunith8/git-cli-tips-and-tricks-to-be-more-productive-1a3pb4fyvn)
* https://github.com/tj/git-extras
* https://ndpsoftware.com/git-cheatsheet.html

## How To's

### [Creating Orphaned branches](https://stackoverflow.com/questions/1384325)

```bash
git checkout --orphan gh-pages
git rm -rf .
<do work>
git add your files
git commit -m 'Initial commit'
```

### [How do I update a GitHub forked repository?](https://stackoverflow.com/q/7244321/)  

#### fork https://github.com/kubernetes/website

```
git clone https://github.com/butuzov/kubernetes-website .
# Add the remote, call it "upstream":
git remote add upstream https://github.com/kubernetes/website.git
# get updates
git fetch upstream
# merge updaes
git merge upstream/master master
# push
git push -f origin master
```

#### How to see what happened to the file in path of the repo?

```bash
git log --full-history -- content/docs/examples/
```

### Rabase

[How to use git rebase effectively](https://www.codementor.io/@kpunith8/how-to-use-git-rebase-effectively-1a8f0o39hf)

#### Rebase Changes from upstream into your branch

```
git checkout target
git rebase master
```

#### [Edit Old Commits](https://stackoverflow.com/questions/1186535/how-to-modify-a-specified-commit)

```bash
git rebase --interactive 'bbc643cd^'
# edit to commit
git commit --all --amend
git rebase --continue
```

### Squash branch 

```bash
git log --oneline
git rebase -i HEAD~4
> pick to squash
```