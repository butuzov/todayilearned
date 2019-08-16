# Special Parameters

* https://www.gnu.org/software/bash/manual/bash.html#Special-Parameters

## `$*` vs `$@`  - Expands to the positional parameters, starting from one.

* https://unix.stackexchange.com/q/129072

```bash
# Here, the parameters are regarded as one long quoted string
for a in "$*"; do
    echo $a;
done
output > one two three four

# The string is broken into words by the for loop.
echo -e "\nUsing \$*:"
for a in $*; do
    echo $a;
done
output > one
output > two
output > three
output > four

# This treats each element of $@ as a quoted string.
for a in "$@"; do
    echo $a;
done
output > one
output > two
output > three four

# This treats each element as an unquoted string
for a in $@; do
    echo $a;
done
output > one
output > two
output > three
output > four
```

## `$#`  - Expands to the number of positional parameters in decimal

```bash
result(){
  echo $#
}

result one two "three four"
output > 3
```

## `$?` - Exit Status

```bash
echo "12312312"
echo $?
output > 0
```

## `$-` - The current set of options may be found in

```bash
echo $-
output (example) > hB
# to see what options set check command
man set
```

## `$$` - Expands to the process ID of the shell

```bash
ps -a | grep $$ | grep -v grep
output > 1003 ttys003    0:00.35 /bin/bash -l
```

## `$!` - echo Expands to the process ID of the job most recently placed into the background

```bash
sleep 60 &
output > [1] 1629
echo $!
output > 1629
```

## `$0` - Expands to the name of the shell or shell script.

```bash
echo $0
output > /bin/bash
```

## `$_`  - last argument of previous input.

```bash
mkdir long_long_long_directory_name
cd $_
```


