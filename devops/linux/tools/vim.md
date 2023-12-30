# `vim`

## Shortcuts

Shortcut       | Description
---------------|---------------------------
`x`            | to delete the unwanted character

`u`            | to undo the last the command and U to undo the whole line
`CTRL`+`R`     | redo
`:q!`          | to trash all changes
`dw`           |  move the cursor to the beginning of the word to delete that word
`2w`           | to move the cursor two words forward.
`3e`           | to move the cursor to the end of the third word forward.
`0`            | (zero) to move to the start of the line.
`d2w`          | which deletes 2 words ..
`dd`           | to delete the line and 2dd to delete to line .number can be changed for deleting the number of consecutive words
`p`            | puts the previously deleted text after the cursor(Type dd to delete the line and store it in a Vim register. and p to put the line)
`r`            | to replace the letter e.g press re to replace the letter with e
`ce`           | to change until the end of a word (place the cursor on the u in lubw it will delete ubw )
`ce`           | deletes the word and places you in Insert mode
`G`            | to move you to the bottom of the file.
`gg`           | to move you to the start of the file.
`%`            | to find a matching ),], or }
`/`            | backward search n to find the next occurrence and N to search in opposite direction
`?`            | forward search
`:!`           | to run the shell commands like :!dir, :!ls
`:w`           | TEST (where TEST is the filename you chose.) . Save the file
`v`            | starts visual mode for selecting the lines and you can perform operation on that like d delete
`:r`           | Filename will insert the content into the current file
`R`            | to replace more than one character
`y`            | operator to copy text using v visual mode and p to paste it
`yw`           | (copy)yanks one word
`e`            | command moves to the end of a word.
`y`            | operator yanks (copies) text, p puts (pastes) it.
`R`            | enters Replace mode until <ESC> is pressed.
`CTRL`+`W`     | to jump from one window to another
`:cq`          | quit and return error (helpful when using Vim with Git)
`:qa`          | quit all buffers
`:xa`          | quit and save all buffers

## Vim (Updates)
Shortcut       | Description
---------------|---------------------------
`$`            | Go to End of the Line
`A`            | Go to End of the Line and Switch to Insert Mode.
`o`            | Insert Line Under Current and Switch to Insert Mode.
`O`            | Insert Line Above Current and Switch to Insert Mode.
`|`            | Go to Begin of the Line
`:s/old/new/g` | Replace 'new' for 'old' where g is globally
`:s/old/new`   | Replace 'new' for 'old'

## Vim (Delete)
Shortcut       | Description
---------------|---------------------------
`d+G`          | Delete this and next LINES

## How to Quit Vim (Updates)
Shortcut       | Description
---------------|---------------------------
`:wq`          | To save and exit
`:x`           | To save and exit

## Time Travel

`:earlier 35s` | earlier 35s
`:later 12s`   | later 12s


### reading list
- https://coderwall.com/p/rwmdvq/seamlessly-navigate-tmux-and-vim-splits
