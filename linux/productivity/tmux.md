# `tmux`

```shell
brew install tmux
```

> Tmux is a terminal multiplexer; it allows you to create several "pseudo terminals" from a single terminal.

## Managing

### Sessions

| Action                    | Key                           | Commands           | Shell                              |
| ------------------------- | ----------------------------- | ------------------ | ---------------------------------- |
| **Create**                |                               | `:new -s db`       | `tmux new -s db`                   |
| **Create** or **Attache** |                               | `:new -A -s db`    | `tmux new -A -s db`                |
| **Close**                 |                               | `:kill-session`    | `tmux kill-session -t name`        |
| **Close** all but current |                               | `:kill-session -a` | `tmux kill-session -a`             |
| **Attach**                |                               | `:attach -t vpn`   | `tmux attach -t vpn`               |
| **Deattach**              | <kbd>lead</kbd>, <kbd>d</kbd> |                    |                                    |
| **List**                  | <kbd>lead</kbd>, <kbd>s</kbd> | `:ls`              | `tmux ls`                          |
| **List** also windows     | <kbd>lead</kbd>, <kbd>w</kbd> |                    |                                    |
| **Rename**                | <kbd>lead</kbd>, <kbd>$</kbd> |                    | `tmux rename-session -t smap smtp` |
| Navigate **Next**         | <kbd>lead</kbd>, <kbd>(</kbd> |                    |                                    |
| Navigate **Prev**         | <kbd>lead</kbd>, <kbd>)</kbd> |                    |                                    |

### Windows

| Action                        | Key                             |
| ----------------------------- | ------------------------------- |
| **Create** Window             | <kbd>lead</kbd>, <kbd>c</kbd>   |
| **Close** Window              | <kbd>lead</kbd>, <kbd>&</kbd>   |
| **Switch** to N-Window        | <kbd>lead</kbd>, <kbd>0-9</kbd> |
| **Rename** Window             | <kbd>lead</kbd>, <kbd>,</kbd>   |
| **List** Window               | <kbd>lead</kbd>, <kbd>w</kbd>   |
| **Toggle** last Active Window | <kbd>lead</kbd>, <kbd>l</kbd>   |
| **Next** Window               | <kbd>lead</kbd>, <kbd>n</kbd>   |
| **Prev** Window               | <kbd>lead</kbd>, <kbd>p</kbd>   |
| **Find** Window               | <kbd>lead</kbd>, <kbd>f</kbd>   |

### Panes

| Action                    | Key                                           | Commnds              |
| ------------------------- | --------------------------------------------- | -------------------- |
| **Close**                 | <kbd>lead</kbd>, <kbd>x</kbd>                 |                      |
| **Switch Panes**          | <kbd>lead</kbd>, <kbd>←, →, ↑, ↓</kbd>        |                      |
| **Split Vertically**      | <kbd>lead</kbd>, <kbd>"</kbd>                 | `:split-window -h`   |
| **Split Horizontally**    | <kbd>lead</kbd>, <kbd>%</kbd>                 | `:split-window -v`   |
| **Resize**                | <kbd>lead</kbd>, <kbd>b</kbd>                 | `:resize-pane -L 10` |
| **Show Panes Nnmbers**    | <kbd>lead</kbd>, <kbd>q</kbd>                 |                      |
| **Switch** to N-Pane      | <kbd>lead</kbd>, <kbd>q</kbd> ,<kbd>0-9</kbd> |                      |
| **Toogle** fullscreen     | <kbd>lead</kbd>, <kbd>z</kbd>                 |                      |
| **Toogle** layouts        | <kbd>lead</kbd>, <kbd>space</kbd>             |                      |
| **Convert** to **Window** | <kbd>lead</kbd>, <kbd>!</kbd>                 |                      |
| **Move** pane Left        | <kbd>lead</kbd>, <kbd>{</kbd>                 |                      |
| **Move** pane Right       | <kbd>lead</kbd>, <kbd>}</kbd>                 |                      |

## Config `~/tmux.conf`

### Lead Key

```
set -g prefix C-z
unbind C-z
bind C-z send-prefix

```

### Other

```tmux
# Enable mouse control (clickable windows, panes, resizable panes)
set -g mouse on
setw -g mouse on

# Commands
bind R source-file ~/.tmux.conf \; display "Reloaded!" # Reload
```

## Read more

- https://tmuxai.dev/
- [Managing tmux Sessions – A Beginner‘s Guide](https://thelinuxcode.com/managing-tmux-sessions-ubuntu-20-04-lts/)
