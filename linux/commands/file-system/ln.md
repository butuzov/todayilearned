# Soft and Hard Links

### Files

`ln -s Original/File_Location Target/File_Location`

### Directories

If It's Relative Link - start in **Target** directory.

#### Correct Way
```bash
$ ln -s ../../../../WP_Plugins_Dev/Yojimbo  Yojimbo \
  && ls -la

drwxr-xr-x@  7 butuzov  admin   238 Mar 13 18:14 .
drwxr-xr-x@ 11 butuzov  admin   374 Mar 13 18:00 ..
lrwxr-xr-x   1 butuzov  admin    34 Mar 13 18:14 Yojimbo -> ../../../../WP_Plugins_Dev/Yojimbo
```

#### Incorrect way

```bash
$ ln -s Yojimbo ../www.made.ua/www-main/wp-content/plugins \
  && cd ../www.made.ua/www-main/wp-content/plugins \
  && ls -la

drwxr-xr-x@  7 butuzov  admin   238 Mar 13 18:09 .
drwxr-xr-x@ 11 butuzov  admin   374 Mar 13 18:00 ..
lrwxr-xr-x   1 butuzov  admin    10 Mar 13 18:09 Yojimbo -> Yojimbo
```


### More Examples
```bash
# mass linking of cuda's libraries
ls -l /Developer/NVIDIA/CUDA-9.1/lib/ | awk '{print $9}' | xargs  -I {} sh -c "ln -s /Developer/NVIDIA/CUDA-9.1/lib/{} /usr/local/lib/{}"
```
