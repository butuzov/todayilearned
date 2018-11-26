# `grep` File pattern searcher

![grep by Julia Evans](grep.jpg)

```bash
# safe cleanup local pip folder
python3 -m pip list | awk 'NR>2 {print $1}' | grep -Ev "pip|setuptools|wheel" | xargs -I {} sh -c "python3 -m pip uninstall {} -y"

# getting local ip addresses info
ifconfig | grep inet | grep -v inet6 | awk '{print $2}'
```
