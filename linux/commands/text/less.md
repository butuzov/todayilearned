# `less` opposite of more

![less by Julia Evans](less.jpg)

```bash
# view file via less
less And_Then_There_Were_None.txt

# same but strting from 6th line.
less +6 And_Then_There_Were_None.txt
```

## Keyboard shortcuts

*   [Arrows]/[Page Up]/[Page Down]/[Home]/[End]: Navigation.
*   [Space bar]: Next page.
*   b: Previous page.
*   ng: Jump to line number n. Default is the start of the file.
*   nG: Jump to line number n. Default is the end of the file.
*   /pattern: Search for pattern. Regular expressions can be used.
*   n: Go to next match (after a successful search).
*   N: Go to previous match.
*   ‘^ or g: Go to start of file.
*   ‘$ or G: Go to end of file.
*   s: Save current content (got from another program like grep) in a file.
*   =: File information.
*   F: continually read information from file and follow its end. Useful for logs watching. Use Ctrl+C to exit this mode.
*   -option: Toggle command-line option -option.
*   h: Help.
*   q or zz: Quit.

### Forward Search Navigation
*  / – search for a pattern which will take you to the next occurrence.
*  n – for next match in forward
*  N – for previous match in backward

### Backward Search Navigation
*  ? – search for a pattern which will take you to the previous occurrence.
*  n – for next match in forward
*  N – for previous match in backward

### Line Navigation

*  j – navigate forward by one line
*  k – navigate backward by one line


### Screen Navigation

*  CTRL+D – forward half window
*  CTRL+U – backward half window
*  CTRL+F – forward one window
*  CTRL+B – backward one window

### Count numbers as options

*  5j – 5 lines forward.
*  10k – 10 lines backward.

###  Some more less options

*  :p - Examine the previous file in the command line list.
*  :d - Remove the current file from the list of files.
*  v – using the configured editor edit the current file.
*  &pattern – display only the matching lines, not all.
*  CTRL+G – show the current file name along with line, byte and percentage statistics.
*  mx – mark the current position with the letter `x`.
*  ‘x – go to the marked position `x`.
