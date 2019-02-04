# `tar` Tape Archives

![`tar` by Julia Evans](tar.jpg)


### Examples

```bash
# zip folder `media` at /mnt into file `media.tar.gz`
# and pipe output to /dev/null
tar -zcvf ~/media.tar.gz -C /mnt media &> /dev/null

# Let tar to find out which algorith to use for compression.
tar acvf ~/media.tar.gz -C /mnt media &> /dev/null

# Same, but tar
tar -cvf ~/media.tar -C /mnt media &> /dev/null

# Excluding files from archiving
tar -cf notes.tar * --exclude "*.md"

# Or using a list of patterns from file
tar -cf notes.tar * --X list.file

# Adding extra (folder) to existing tar
tar -rvf ~/media.tar -C /mnt extra

# Show tar contents
tar -tf ~/media.tar

# Untar into current directory
tar -vxf ~/media.tar

# merge to archivers
tar -Af one.tar two.tar

# add file if modified
tar -uf notes.tar test.txt
tar -uvvf notes.tar test.txt



# вудуеу
# ---------------
tar -f notes.tar --delete test.txt
```
Related: [`zip`](./zip)
