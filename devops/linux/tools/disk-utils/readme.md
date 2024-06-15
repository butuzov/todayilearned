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

MacOS specific tool to manage volumes, disks, partitions [diskutil](../../../../macos/shell/diskutil/readme.md)

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

## `mdfind` Spotlight's search.

```shell
mdfind search phraze -onlyin "/Volumes/SharedDocuments" | less
```

## `find`

```shell
# Searching for pattern *torrent* exceprt in ./Volumes
find * -path ./Volumes -prune -o -name "*torrent*"

# Searching and Executing Command ( -exec )
find * -path ./Volumes -prune -o -name "*torrent*" -exec file -i '{}' \;

# basic 'find file' commands
#--------------------------
find / -name foo.txt -type f -print             # full command
find / -name foo.txt -type f                    # -print isn't necessary
find / -name foo.txt                            # don't have to specify "type==file"
find . -name foo.txt                            # search under the current dir
find . -name "foo.*"                            # wildcard
find . -name "*.txt"                            # wildcard
find /users/al -name Cookbook -type d           # search '/users/al' dir

#search multiple dirs
#--------------------
find /opt /usr /var -name foo.scala -type f     # search multiple dirs

# search by user
#--------------------------
find . -user butuzov

# case-insensitive searching
#--------------------------
find . -iname foo                               # find foo, Foo, FOo, FOO, etc.
find . -iname foo -type d                       # same thing, but only dirs
find . -iname foo -type f                       # same thing, but only files

# find files with different extensions
#------------------------------------
find . -type f \( -name "*.c" -o -name "*.sh" \)                       # *.c and *.sh files
find . -type f \( -name "*cache" -o -name "*xml" -o -name "*html" \)   # three patterns

# dealing with size
# --------------------------------------------
find / -size +10000k                                              # filesize > 10000k
find / -mount -type d -exec du -s "{}" \; | sort -n               # find directories and sort by size


# find files that don't match a pattern (-not)
# --------------------------------------------
find . -type f -not -name "*.html"                                # find all files not ending in ".html"

# find files by text in the file (find + grep)
#--------------------------------------------
find . -type f -name "*.java" -exec grep -l StringBuffer {} \;    # find StringBuffer in all *.java files
find . -type f -name "*.java" -exec grep -il string {} \;         # ignore case with -i option
find . -type f -name "*.gz" -exec zgrep 'GET /foo' {} \;          # search for a string in gzip'd files

# 5 lines before, 10 lines after grep matches
-------------------------------------------
find . -type f -name "*.scala" -exec grep -B5 -A10 'null' {} \;
# http://alvinalexander.com/linux-unix/find-grep-print-lines-before-after-search-term

# find files and act on them (find + exec)
#----------------------------------------
find /usr/local -name "*.html" -type f -exec chmod 644 {} \;      # change html files to mode 644
find htdocs cgi-bin -name "*.cgi" -type f -exec chmod 755 {} \;   # change cgi files to mode 755
find . -name "*.pl" -exec ls -ld {} \;                            # run ls command on files found

# find and copy
#-------------
find . -type f -name "*.mp3" -exec cp {} /tmp/MusicFiles \;       # cp *.mp3 files to /tmp/MusicFiles

# copy one file to many dirs
#--------------------------
find dir1 dir2 dir3 dir4 -type d -exec cp header.shtml {} \;      # copy the file header.shtml to those dirs

# find and delete
#---------------
find . -type f -name "Foo*" -exec rm {} \;                        # remove all "Foo*" files under current dir
find . -type d -name CVS -exec rm -r {} \;                        # remove all subdirectories named "CVS" under current dir
find ./dist -name "*.js*" -exec cp {} $DIR/lineage.js \;          # copy
find . -type f -name '*.css' -exec sha1sum "$PWD"/{} \; | awk 'END{n=split($2,a,/\//); print (a[n],">",$1)}' # sha1

# find files by modification time
#-------------------------------
find . -mtime 1               # 24 hours
find . -mtime -7              # last 7 days
find . -mtime -7 -type f      # just files
find . -mtime -7 -type d      # just dirs

# find files by modification time using a temp file
#-------------------------------------------------
touch 09301330 poop           # 1) create a temp file with a specific timestamp
find . -mnewer poop           # 2) returns a list of new files
rm poop                       # 3) rm the temp file

# find with time: this works on mac os x
#--------------------------------------
find / -newerct '1 minute ago' -print

# find and tar
#------------
find . -type f -name "*.java" | xargs tar cvf myfile.tar
find . -type f -name "*.java" | xargs tar rvf myfile.tar

# http://alvinalexander.com/blog/post/linux-unix/using-find-xargs-tar-create-huge-archive-cygwin-linux-unix

# find, tar, and xargs
#--------------------
find . -name -type f '*.mp3' -mtime -180 -print0 | xargs -0 tar rvf music.tar
# (-print0 helps handle spaces in filenames)
# http://alvinalexander.com/mac-os-x/mac-backup-filename-directories-spaces-find-tar-xargs

# find and pax (instead of xargs and tar)
#---------------------------------------
find . -type f -name "*html" | xargs tar cvf jw-htmlfiles.tar -
find . -type f -name "*html" | pax -w -f jw-htmlfiles.tar
# http://alvinalexander.com/blog/post/linux-unix/using-pax-instead-of-tar

# find files with permitions
#----------
find . -perm -u+s
```

## `ls`
List directory contents

```bash
# Usage Examples
# ------------------------------------------
ls -1      # One file/directory per line
ls -lah    # Show files size in "human" size

# Order -----------------------------------
ls -lt     # time based order
ls -ltr    # time based reverse order

# Filtering -------------------------------
ls -a      # Show also hidden files
ls -A      # Show also hidden files  but not . and ..

# Options ---------------------------------
ls -i      # Inodes numbers
ls -ln     # UID/GID

# POSIX RegExps
ls -l backup[^[:lower:]^[:digit:]][[:punct:]][^:alpha:]*
```

## `ln`

Make links between files

```bash
ln -s Original/File_Location Target/File_Location

# Correct way
ln -s ../../../../WP_Plugins_Dev/Yojimbo  Yojimbo && ls -la

> drwxr-xr-x@  7 butuzov  admin   238 Mar 13 18:14 .
> drwxr-xr-x@ 11 butuzov  admin   374 Mar 13 18:00 ..
> lrwxr-xr-x   1 butuzov  admin    34 Mar 13 18:14 Yojimbo -> ../../../../WP_Plugins_Dev/Yojimbo

# Incorrect way
ln -s Yojimbo ../www.made.ua/www-main/wp-content/plugins \
  && cd ../www.made.ua/www-main/wp-content/plugins \
  && ls -la

> drwxr-xr-x@  7 butuzov  admin   238 Mar 13 18:09 .
> drwxr-xr-x@ 11 butuzov  admin   374 Mar 13 18:00 ..
> lrwxr-xr-x   1 butuzov  admin    10 Mar 13 18:09 Yojimbo -> Yojimbo


# creates links for cuda libraries
CUDA_LOCATION=/Developer/NVIDIA/CUDA-9.1/lib/
ls -l $CUDA_LOCATION \
    | awk '{print $9}' \
    | xargs  -I {} sh -c "ln -s $CUDA_LOCATION{} /usr/local/lib/{}"
```


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
