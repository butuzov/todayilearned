# `test`

```bash
> test 4 -lt 6
# Info about result in the exitcode
> echo $?
0

> test 4 -gt `> echo $?
1

# or
echo $(test 4 -lt 6) $?

# or
test 4 -lt 6 && echo 1 || echo 0
```
### `man test`

  Example | Description
----------|--------------
`-b file`            | True if file exists and is a block special file.
`-c file`            | True if file exists and is a character special file.
`-d file`            | True if file exists and is a directory.
`-e file`            | True if file exists (regardless of type).
`-f file`            | True if file exists and is a regular file.
`-g file`            | True if file exists and its set group ID flag is set.
`-h file`            | True if file exists and is a symbolic link.
`-k file`            | True if file exists and its sticky bit is set.
`-n string`          | True if the length of string is nonzero.
`-p file`            | True if file is a named pipe (FIFO).
`-r file`            | True if file exists and is readable.
`-s file`            | True if file exists and has a size greater than zero.
`-t file_descriptor` | True if the file whose file descriptor number is file_descriptor is open and is associated with a terminal.
`-u file`            | True if file exists and its set user ID flag is set.
`-w file`            | True if file exists and is writable.  True indicates only that the write flag is on.  The file is not writable on a read-only file system even if this test indicates true.
`-x file`            | True if file exists and is executable.  True indicates only that the execute flag is on.  If file is a directory, true indicates that file can be searched.
`-z string`          | True if the length of string is zero.
`-L file`            | True if file exists and is a symbolic link.
`-O file`            | True if file exists and its owner matches the effective user id of this process.
`-G file`            | True if file exists and its group matches the effective group id of this process.
`-S file`            | True if file exists and is a socket.
`file1 -nt file2`    | True if file1 exists and is newer than file2.
`file1 -ot file2`    | True if file1 exists and is older than file2.
`file1 -ef file2`    | True if file1 and file2 exist and refer to the same file.
`string`             | True if string is not the null string.
`s1 = s2`            | True if the strings s1 and s2 are identical.
`s1 != s2`           | True if the strings s1 and s2 are not identical.
`s1 < s2`            | True if string s1 comes before s2 based on the binary value of their characters.
`s1 > s2`            | True if string s1 comes after s2 based on the binary value of their characters.
`n1 -eq n2`          | True if the integers n1 and n2 are algebraically equal.
`n1 -ne n2`          | True if the integers n1 and n2 are not algebraically equal.
`n1 -gt n2`          | True if the integer n1 is algebraically greater than the integer n2.
`n1 -ge n2`          | True if the integer n1 is algebraically greater than or equal to the integer n2.
`n1 -lt n2`          | True if the integer n1 is algebraically less than the integer n2.
`n1 -le n2`          | True if the integer n1 is algebraically less than or equal to the integer n2.
