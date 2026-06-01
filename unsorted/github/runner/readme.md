# Self Hosted Runner


## RTFM

- https://github.com/features/actions
- https://github.com/actions/runner
- https://docs.github.com/en/actions/how-tos/manage-runners/self-hosted-runners
- https://docs.github.com/en/actions/concepts/runners/self-hosted-runners
- [How to Configure Self-Hosted Runners in GitHub Actions](https://oneuptime.com/blog/post/2026-01-25-github-actions-self-hosted-runners/view)
- [hCustom GitHub Runner Images With Pre- and Post-Job Scripts](https://www.kenmuse.com/blog/custom-github-runner-images-with-pre-and-post-job-scripts/)

## Quick Start 

1. Register Runner via https://github.com/user/repo/settings/actions/runners
2. Follow for registration 
	```shell
	./config.sh --url https://github.com/your-org/your-repo --token YOUR_TEMP_TOKEN
	```
3. Run It.
	```
	./run.sh
	```
### Run as a background service

```shell
# Instead of keeping a terminal open:
sudo ./svc.sh install
sudo ./svc.sh start

# Check status:
sudo ./svc.sh status

# Logs:
tail -f _diag/*.log
```

### Tweak Job Pipeline

```yaml
...
	runs-on: self-hosted
...
```

