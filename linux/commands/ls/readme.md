
# `ls` List directory contents


```bash
# One file/directory per line
ls -1

# human size
ls -lah

# time based order
ls -lt

# time based reverse order
ls -ltr

# Hidden withou . and ..
ls -A

# Hidden with
ls -a

# Inodes numbers
ls -i

# UID/GID
ls -ln
```

### Example of `ls -li`
 inode  | Type-Permissions | Number of Links | Owner   | Group | Size | Last-Modified | Name
--------|------------------|-----------------|---------|-------|------|---------------|-----------
5649755 | drwxr-xr-x       | 70              | butuzov | staff | 2380 | Jan 23 23:45  | KubeCon-2018
3990493 | drwxr-xr-x       | 56              | butuzov | staff | 1904 | Jan  2 13:13  | GopherCon-2017
