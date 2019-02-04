# `xargs` Construct Argument List(s) and Execute Utility

![xargs by Julia Evans](xargs.jpg)


The xargs utility reads space, tab, newline and end-of-file delimited strings from the standard input and executes utility with the strings as arguments.

Any arguments specified on the command line are given to utility upon each invocation, followed by some number of the arguments read from the standard input of xargs.  The utility is repeatedly executed until standard input is exhausted.

Spaces, tabs and newlines may be embedded in arguments using single (`'`) or double (`"`) quotes or backslashes (`\`).  Single quotes escape all non-single quote characters, excluding newlines, up to the matching single quote.  Double quotes escape all non-double quote characters, excluding newlines, up to the matching double quote.  Any single character, including newlines, may be escaped by a backslash.


#### Examples

  ```bash
  # Getting Git Subcomands Help Pages to Files.
  #-----------------------------------------------
  git help -a | \
    awk '{printf "%s\n%s\n", $1, $2}' | \
    head -n 172 | \
    tail -n 156 | \
    xargs -I {} sh -c "git {} -h > {}.md"

  # Copy the list of files and directories which start with an
  # uppercase letter in the current directory to destdir
  # -------
  ls -1d [A-Z]* | xargs -J % cp -rp % destdir
  ```

## Options

  Key   |  Short  | Description
--------|---------|--------------------------------------------------------------------
  `-0`  |         | Change xargs to expect NUL (`\0`) characters as separators, instead of spaces and newlines.  This is expected to be used in concert with the `-print0` function in `find(1)`.
  `-E`  | eofstr  | Use eofstr as a logical EOF marker.
  `-I`  | replstr |  Execute utility for each input line, replacing one or more occurrences of replstr in up to replacements (or 5 if no `-R` flag is speci fied) arguments to utility with the entire line of input.  The resulting arguments, after replacement is done, will not be allowed to grow beyond 255 bytes; this is implemented by concatenating as much of the argument containing replstr as possible, to the con structed arguments to utility, up to 255 bytes. The 255 byte limit does not apply to arguments to utility which do not contain replstr, and furthermore, no replacement will be done on utility itself.  Implies `-x`.
  `-J`  | replstr |  If this option is specified, xargs will use the data read from standard input to replace the first occurrence of replstr instead of appending that data after all other arguments.  This option will not affect how many arguments will be read from input (`-n`), or the size of the command(s) xargs will generate (`-s`).  The option just moves where those arguments will be placed in the command(s) that are executed.  The replstr must show up as a distinct argument to xargs.  It will not be recognized if, for instance, it is in the middle of a quoted string.  Furthermore, only the first occurrence of the replstr will be replaced.
  `-L`  | number |  Call utility for every number non-empty lines read.  A line ending with a space continues to the next non-empty line.  If EOF is reached and fewer lines have been read than number then utility will be called with the available lines.  The `-L` and `-n` options are mutually-exclusive; the last one given will be used.
  `-n`  | number | Set the maximum number of arguments taken from standard input for each invocation of utility.  An invocation of utility will use less than number standard input arguments if the number of bytes accumulated (see the `-s` option) exceeds the specified size or there are fewer than number arguments remaining for the last invocation of utility.  The current default value for number is 5000.
  `-o`  |     | Reopen stdin as `/dev/tty` in the child process before executing the command.  This is useful if you want xargs to run an interactive application.
  `-P`  | maxprocs |  Parallel mode: run at most maxprocs invocations of utility at once.
  `-p`  |      | Echo each command to be executed and ask the user whether it should be executed.  An affirmative response, `y' in the POSIX locale, causes the command to be executed, any other response causes it to be skipped.  No commands are executed if the process is not attached to a terminal.
  `-R`  | replacements | Specify the maximum number of arguments that -I will do replacement in.  If replacements is negative, the number of arguments in which to replace is unbounded.
  `-s`  | size |  Set the maximum number of bytes for the command line length provided to utility.  The sum of the length of the utility name, the arguments passed to utility (including NULL terminators) and the current environment will be less than or equal to this number. The current default value for size is ARG_MAX - 4096.
  `-t`  |      | Echo the command to be executed to standard error immediately before it is executed.
  `-x`  |      | Force xargs to terminate immediately if a command line containing number arguments will not fit in the specified (or default) command line length.


Undefined behavior may occur if utility reads from the standard input.

The xargs utility exits immediately (without processing any further input) if a command line cannot be assembled, utility cannot be invoked, an invocation of utility is terminated by a signal, or an invocation of utility exits with a value of 255.


###### @todo: Add more examples
