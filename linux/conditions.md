# Conditions in BASH

## if, elif, else
Syntax quick recap...

```bash
  if [[ CONDITION ]]; then
    echo 'if case'
  elif [[ CONDITION ]]]; then
    echo 'elif case'
  else
    echo 'else case'
  fi
```


## File-Based Conditions Types:

#### `[ -a existingfile ]`
File *existingfile* exists.

```bash
  if [ -a tmp.tmp ]; then
    # Make sure we’re not bothered by an old temporary file
    rm -f tmp.tmp
  fi
```


#### `[ -b blockspecialfile ]`
File *blockspecialfile* exists and is block special.

```bash
  # Block special files are special kernel files found in /dev,
  # mainly used for ATA devices like hard disks, cd-roms and floppy disks.
  if [ -b /dev/fd0 ]; then
    # Write an image to a floppy
    dd if=floppy.img of=/dev/fd0
  fi
```


#### `[ -c characterspecialfile ]`
File *characterspecialfile* exists and is character special.

```bash
  # Character special files are special kernel files found in /dev,
  # used for all kinds of purposes (audio hardware, tty’s, but also /dev/null).
  if [ -c /dev/dsp ]; then
    # This actually works for certain raw wav files
    cat raw.wav > /dev/dsp
  fi
```


#### `[ -d directory ]`
File *directory* exists and is a directory.

```bash
# In UNIX-style, directories are a special kind of file.
if [ -d ~/.kde ]; then
    echo "You seem to be a kde user."
fi
```


#### `[ -e existingfile ]`
File *existingfile* exists.

```bash
  # (same as -a, see that entry for an example)
  if [ -e tmp.tmp ]; then
    # Make sure we’re not bothered by an old temporary file
    rm -f tmp.tmp
  fi
```


#### `[ -f regularfile ]`
File *regularfile* exists and is a regular file.

```bash
# A regular file is neither a block or character special
# file nor a directory.
if [ -f ~/.bashrc ]; then
    source ~/.bashrc
fi
```


#### `[ -g sgidfile ]`
File *sgidfile* exists and is set-group-ID.

```bash
# When the SGID-bit is set on a directory, all files
# created in that directory will inherit the group of
# the directory.
if [ -g . ]; then
   echo "Created files are inheriting the group ‘$(ls -ld . | awk ‘{ print $4 }’)’ from the working directory."
fi
```


#### `[ -G fileownedbyeffectivegroup ]`
File *fileownedbyeffectivegroup* exists and is owned by the effective group ID.

```bash
# The effective group id is the primary group id of the executing user.
if [ ! -G file ]; then
  # An exclamation mark inverts the outcome of the condition following it
  # Change the group if it’s not the effective one
  chgrp $(id -g) file
fi
```


#### `[ -h symboliclink ]`
File *symboliclink* exists and is a symbolic link.

```bash
if [ -h $pathtofile ]; then
  pathtofile=$(readlink -e $pathtofile)
  # Make sure $pathtofile contains the actual file and not a
  # symlink to it
fi
```


#### `[ -k stickyfile ]`
File *stickyfile* exists and has its sticky bit set.

```bash
# The sticky bit has got quite a history, but is now used to
# prevent world-writable directories from having their contents
# deletable by anyone.
if [ ! -k /tmp ]; then
  # An exclamation mark inverts the outcome of the condition
  # following it
  echo "Warning! Anyone can delete and/or rename your files in /tmp!"
fi
```


##### `[ -L symboliclink ]`
File *symboliclink* exists and is a symbolic link.

```bash
  if [ -L $pathtofile ]; then
    pathtofile=$(readlink -e $pathtofile)
    # Make sure $pathtofile contains the actual file and not a
    # symlink to it
  fi
```


#### `[ -N modifiedsincelastread ]`
File *modifiedsincelastread* exists and was modified after the last read.

```bash
if [ -N /etc/crontab ]; then
  # SIGHUP makes crond reread all crontabs
  killall -HUP crond
fi
```


#### `[ -O fileownedbyeffectiveuser ]`
File *fileownedbyeffectiveuser* exists and is owned by the user executing the script.

```bash
if [ -O file ]; then
  # Makes the file private, which is a bad idea if you don’t own it
  chmod 600 file
fi
```


#### `[ -p namedpipe ]`
File *namedpipe* exists and is a named pipe.

```bash
# A named pipe is a file in /dev/fd/ that can be read just once.
if [ -p $file ]; then
  cp $file tmp.tmp  # Make sure we’ll be able to read
  file="tmp.tmp"    # the file as many times as we like
fi
```


#### `[ -r readablefile ]`
File *readablefile* exists and is readable to the script.

```bash
if [-r file ]; then
  content=$(cat file) # Set $content to the content of the file
fi
```


#### `[ -s nonemptyfile ]`
File *nonemptyfile* exists and has a size of more than 0 bytes.

```bash
if [ -s logfile ]; then
  gzip logfile    # Backup the old logfile
  touch logfile   # Before creating a fresh one.
fi
```


#### `[ -S socket ]`
File *socket* exists and is a socket.

```bash
# A socket file is used for inter-process communication, and features
# an interface similar to a network connection.
if [ -S /var/lib/mysql/mysql.sock ]; then
  mysql –socket=/var/lib/mysql/mysql.sock # See this MySQL tip
fi
```


#### `[ -t openterminal ]`
File descriptor *openterminal* exists and refers to an open terminal.

```bash
# Virtually everything is done using files on Linux/UNIX,
# and the terminal is no exception.
if [ -t /dev/pts/3 ]; then
  echo -e "nHello there. Message from terminal $(tty) to you." > /dev/pts/3
  # Anyone using that terminal will actually see this message!
fi
```


#### `[ -u suidfile ]`
File *suidfile* exists and is set-user-ID.

```bash
# Setting the suid-bit on a file causes execution of that file to be
# done with the credentials of the owner of the file, not of the executing
# user.
if [ -u executable ]; then
  echo "Running program executable as user $(ls -l executable | awk ‘{ print $3 }’)."
fi
```


#### `[ -w writeablefile ]`
file *writeablefile* exists and is writeable to the script.

```bash
if [ -w /dev/hda ]; then
  grub-install /dev/hda
fi
```


#### `[ -x executablefile ]`
File *executablefile* exists and is executable for the script.

```bash
# Note that the execute permission on a directory means that it’s
# searchable (you can see which files it contains).
if [ -x /root ]; then
  echo "You can view the contents of the /root directory."
fi
```


#### `[ newerfile -nt olderfile ]`
* File *newerfile* was changed more recently than *olderfile*
* If   *newerfile* exists and *olderfile* doesn*t.

```bash
if [ story.txt1 -nt story.txt ]; then
  echo "story.txt1 is newer than story.txt; I suggest continuing with the former."
fi
```


#### `[ olderfile -ot newerfile ]`
* File *olderfile* was changed longer ago than *newerfile*
* If *newerfile* exists and *olderfile* doesn't.

```bash
if [ /mnt/remote/remotefile -ot localfile ]; then
  cp -f localfile /mnt/remote/remotefile
  # Make sure the remote location has the newest version of the file, too
fi
```


#### `[ same -ef file ]`
file *same* and file *file* refer to the same *device/inode* number.

```bash
if [ /dev/cdrom -ef /dev/dvd ]; then
  echo "Your primary cd drive appears to read dvd’s, too."
fi
```


## String-Based Conditions Types:

#### `[ STRING1 == STRING2 ]`
STRING1 is equal to STRING2.

```bash
# Note: you can also use a single "=" instead of a "=="
if [ "$1" == "moo" ]; then
  echo $cow # Ever tried executing ‘apt-get moo’?
fi
```


#### `[ STRING1 != STRING2 ]`
STRING1 is not equal to STRING2.

```bash
if [ "$userinput" != "$password" ]; then
  echo "Access denied! Wrong password!"
  exit 1 # Stops script execution right here
fi
```


#### `[ STRING1 > STRING2 ]`
STRING1 sorts after STRING2 in the current locale (lexographically).


#### `[ STRING1 < STRING2 ]`
STRING1 sorts before STRING2 in the current locale (lexographically).


####  `[ -n NONEMPTYSTRING ]`
NONEMPTYSTRING has a length of more than zero.

```bash
  # This condition only accepts valid strings, so be sure to
  # quote anything you give to it.
  if [ -n "$userinput" ]; then
    # Only parse if the user actually gave some input.
    userinput=parse($userinput)
  fi
  # Note that you can also omit the "-n", as brackets with
  # just a string in it behave the same.
```


#### `[ -z EMPTYSTRING ]`
EMPTYSTRING is an empty string.

```bash
# This condition also accepts non-string input, like an uninitialized variable:
if [ -z $uninitializedvar ]; then
    uninitializedvar="initialized"
    # -z returns true on an uninitialized variable, so we initialize it here.
fi
```

#### `[[ STRING1 =~ REGEXPATTERN ]]`
STRING1 matches REGEXPATTERN. (** Double-bracket syntax only: ** =)

```bash
# If you are familiar with Regular Expressions,
# you can use this conditions to perform a regex match.
if [[ "$email" =~ "b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,4}b" ]]; then
  echo "$email contains a valid e-mail address."
fi
```


##  Arithmetic (number-based) Conditions Types:

#### `[ NUM1 -eq NUM2 ]`
NUM1 is **EQ**ual to NUM2.

#### `[ NUM1 -ne NUM2 ]`
NUM1 is **N**ot **E**qual to NUM2.

#### `[ NUM1 -gt NUM2 ]`
NUM1 is **G**reater **T**han NUM2.

#### `[ NUM1 -ge NUM2 ]`
NUM1 is **G**reater than or **E**qual to NUM2.

####   `[ NUM1 -lt NUM2 ]`
NUM1 is **L**ess **T**han NUM2.

#### `[ NUM1 -le NUM2 ]`
NUM1 is **L**ess than or **E**qual to NUM2.


## Miscellaneous conditions:

#### `[ -o shelloption ]`
shell option *shelloption* is enabled.

```bash
  # Shell options modify the behaviour of bash, except a
  # few unmodifiable ones that indicate the shell status.

  if [ ! -o checkwinsize ]; then
    # An exclamation mark inverts the outcome of the condition following it
    echo "Shell option checkwinsize is disabled; enabling it so you can resize you terminal window without problems."
    shopt -s checkwinsize
    # This shell option is modifiable
  fi

  if [ -o login_shell ]; then
    # This shell option is not modifiable
    echo "This a a login shell."
  fi
```
