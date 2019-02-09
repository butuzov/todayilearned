## PIP

```bash
# Reseting PIP Packages
python3 -m pip list | awk 'NR>2 {print $1}' | grep -Ev "pip|setuptools|wheel" | xargs -I {} sh -c "python3 -m pip uninstall {} -y"

# Using Custom cache directory
python3 -m pip install --download-cache /cache/directory matplotlib

# Install packages from requirments.txt
python3 -m pip install -r requirements.txt
```


## Cached Directory

  https://pip.pypa.io/en/stable/reference/pip_install/

  `~/Library/Caches/pip`
