# `date` Print or set the system date and time

```bash
DATE="$(date +'%y-%m-%d')"
echo ${DATE}

# Generating password using date
date +%s | sha256sum | base64 | head -c 32
```
