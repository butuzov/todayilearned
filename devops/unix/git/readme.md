# `git`

## Reading

- [Start Using Git on the Command Line Like a Pro in Five Minutes](https://medium.com/better-programming/start-using-git-on-the-command-line-like-a-pro-in-5-minutes-36a6e0007e9f)
- [Git CLI tips and tricks to be more productive](https://www.codementor.io/@kpunith8/git-cli-tips-and-tricks-to-be-more-productive-1a3pb4fyvn)
- https://github.com/tj/git-extras
- https://ndpsoftware.com/git-cheatsheet.html

## Configuration

### Conditional configs

- [Conditional Git Configuration](https://blog.scottlowe.org/2023/12/15/conditional-git-configuration/)

```shell
# ~/.gitconfig
[includeIf "gitdir:~/Work/Code/Repos/"]
    path = ~/Work/Code/Repos/.gitconfig
```

```shell
# ~/Work/Code/Repos/.gitconfig
[user]
    email = name@work-domain.com
    name = Scott Lowe
[commit]
    gpgsign = false
```

### Default Commit Message

```shell
# Setting Commit Message (Default) for repo (Not works in SourceTree)
git config commit.template $(pwd)/.hugo/git-extras/.commit-msg-template
```

## Recipes

### Find Large Files

```shell
# via https://stackoverflow.com/q/10622179/
git rev-list --objects --all |
  git cat-file --batch-check='%(objecttype) %(objectname) %(objectsize) %(rest)' |
  sed -n 's/^blob //p' |
  sort --numeric-sort --key=2 |
  cut -c 1-12,41- |
  $(command -v gnumfmt || echo numfmt) --field=2 --to=iec-i --suffix=B --padding=7 --round=nearest
```

### Orphant Branches

```shell
git checkout --orphan gh-pages
git rm -rf .
<do work>
git add your files
git commit -m 'Initial commit'
```

### Garbage Collection

- [How to remove unreferenced blobs from my Git repository](https://stackoverflow.com/questions/1904860)

```shell
git -c gc.reflogExpire=0 -c gc.reflogExpireUnreachable=0 -c gc.rerereresolved=0 \
    -c gc.rerereunresolved=0 -c gc.pruneExpire=now gc

git gc --prune=now
```

### Rebase

```shell
# Update current branch with changes from base branch
git stash --include-untracked             # Stashing uncommited
git rebase base-branch                    # Rebasing Changes
git stash pop                             # Apply stash and delete

# Or
git checkout target
git rebase base-branch


# Edit Old Commit
git rebase --interactive 'bbc643cd^'
git commit --all --amend
git rebase --continue

# Squash
git log --oneline
# Edit Commit Messages
git rebase -i HEAD~4
```

### Upstream

```shell
git clone https://github.com/butuzov/kubernetes-website .
git remote add upstream https://github.com/kubernetes/website.git # Add the remote, call it "upstream":
git fetch upstream                                                # get updates
git merge upstream/master master                                  # merge updaes
git push -f origin master                                         # push
```

### History

```shell
# What happend to `content/docs/examples/`
git log --full-history -- content/docs/examples/
```
