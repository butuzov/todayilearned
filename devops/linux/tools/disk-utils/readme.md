<!-- menu: "Disk Utils" -->
# Disk utils

- See **5** https://www.thegeekstuff.com/2008/10/6-awesome-linux-cd-command-hacks-productivity-tip3-for-geeks/

## `du` - Disk Usage

```shell
# Usage Examples -----------------------------------------------------------------------------------
du -sh                                             # Summary
du -h                                              # Human Values (K/M/G) instead of bytes
du -ha .                                           # with files
du -h --max-depth 1 .                              # Max Depth 1
du -hd 1                                           # Short way to set max depth
du -Pshx .                                         # Stay in filesystem and don't follow symlinks
du -h . | sort -h | head -n 2                      # Sort with Human values
du -xm . | sort -rn | head -n 2                    # Top 2 Directories
du -kx / | awk '{ if ($1 > 500000) { print $0} }'  # Bigger then 50000
du -hd 1 --time .                                  # Include timestamp
du -hd 1 --exclude=".venv" .                       # Exclude by mask
du -BM .                                           # To use block size
```


## `quota` Disk usage and limits

Prints diskusage quota

```shell
quota -v
```

## `rsunc`

```bash
# sync A into B
rsync -avu --delete "/path/to/A" "/path/to/B"
# Sync A contents into B
rsync -avu --delete "/path/to/A/" "/path/to/B"
```

## TODO: `mount` && `umount`

## Data Duplicator `dd`

```bash
# if -input
# of -output
dd if=/dev/xvdf of=/dev/null bs=1M
# 300 times copy 1 mb of data from random to file
dd if=/dev/random of=data/file bs=1m count=300

# Working With HD Images
dd if=/dev/hda of=hdadisk.img
dd if=hdadisk.img of=/dev/hdb
```

## `diskutil`

MacOS specific tool to manage volumes, disks, partitions [diskutil](../../macos/shell/diskutil/readme.md)

## `d`isplay `f`ree disk space

```shell
> df

> df -a # all the file system

> df -h file # h = 1024
> df -H file # h = 1000
> df --total
> df -x tmpfs # exclude tmpfs
```

## `di`

```shell
> di --help
di version 4.52    Default Format: smbuvpT
Usage: di [-ant] [-d display-size] [-f format] [-x exclude-fstyp-list]
       [-I include-fstyp-list] [file [...]]
   -a   : print all mounted devices
   -d x : size to print blocks in (512 - POSIX, k - kbytes,
          m - megabytes, g - gigabytes, t - terabytes, h - human readable).
   -f x : use format string <x>
   -I x : include only file system types in <x>
   -x x : exclude file system types in <x>
   -l   : display local filesystems only
   -n   : do ntt print header
   -t   : print totals
 Format string values:
    m - mount point                     M - mount point, full length
    b - total kbytes                    B - kbytes available for use
    u - used kbytes                     c - calculated kbytes in use
    f - kbytes free                     v - kbytes available
    p - percentage not avail. for use   1 - percentage used
    2 - percentage of user-available space in use.
    i - total file slots (i-nodes)      U - used file slots
    F - free file slots                 P - percentage file slots used
    s - filesystem name                 S - filesystem name, full length
    t - disk partition type             T - partition type, full length

 > di
Filesystem         Mount               Size     Used    Avail %Used  fs Type
/dev/disk3s1s1     /                 926.4G   493.9G   432.4G   53%  apfs
/dev/disk6s1       /Volumes/maca-u   461.1G    10.0G   451.2G    2%  exfat

 > di -f SMP -I apfs | head -n 4
Filesystem     Mount                      %IUsed
/dev/disk3s1s1 /                             0%
/dev/disk3s5   /System/Volumes/Data          0%
/dev/disk1s3   /System/Volumes/Hardware      0%
```

## TODO: `file`

## TODO: `ls`

## TODO: `ln`

## `lsof` to list open files

```shell
# 	Shows  connections to a specific host
lsof -i@IP
# Shows listening ports
lsof -i -sTCP:LISTEN
# Shows established connections
lsof -i -sTCP:ESTABLISHED
#Shows everything interacting with a given DIR
lsof DIR
# What lisning port 80 ?
lsof -i :80
# Opened by process (ID)
lsof -p 17426
# Opened ports
lsof | grep LISTEN
# List all Network File System(NFS) files
lsof -N
# List all TCP or UDP connections
lsof -i tcp
lsof -i udp
# Kill all process that belongs to what you find out (-t returns process id list)
kill -9 `lsof -t options`
# Execute lsof in repeat mode (-r will continue to list, delay, list until a
# interrupt is given. +r will end when no open files are found.)
lsof /path-of-a-file -r5
# List opened files under a directory (+D will recurse the sub directories also)
lsof +D /path-of-a-directory
# List opened files based on process names starting with (You can give multiple -c
# switch on a single command line)
lsof -c ssh -c init
# List files opened by a specific user
lsof -u username
```
