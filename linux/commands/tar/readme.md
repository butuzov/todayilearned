# `tar` Tape Archives

![`tar` by Julia Evans](tar.jpg)


### Examples

```bash
# zip folder `media` at /mnt into file `media.tar.gz`
# and pipe output to /dev/null
# -----------------

tar -zcvf ~/media.tar.gz -C /mnt media &> /dev/null

# Same, but tar
# -----------------
tar -cvf ~/media.tar -C /mnt media &> /dev/null

# Show tar contents
# -----------------
tar -tf ~/media.tar

# Untar into current directory
# -----------------
tar -vxf ~/media.tar

```
Related: [`zip`](./zip)
