# tmux

```bash
# list sessions
tmux ls

# attach session
tmux attach -t <num>

# start new session with name
tmux new -s <name>

# rename session
tmux rename-session -t 0 <name>
```

## Shortcuts

* `^d` or `exit` - close
* `^b` + `⇧+%` - split to right
* `^b` + `<-` - switch to left
* `^b` + `->` - switch to right
* `^b` + `z` - toggle pane fullscreen
* `^b` + `,` - rename current window


## In Session

* `^b` + `c` - new window
* `^b` + `p` - switch to prev
* `^b` + `n` - switch to new
* `^b` + `NUM` - switch to window Num



## Unsorted (`^b` + `?` result)

```
bind-key    -T copy-mode    C-Space           send-keys -X begin-selection                                                                      [164/164]
bind-key    -T copy-mode    C-a               send-keys -X start-of-line
bind-key    -T copy-mode    C-b               send-keys -X cursor-left
bind-key    -T copy-mode    C-c               send-keys -X cancel
bind-key    -T copy-mode    C-e               send-keys -X end-of-line
bind-key    -T copy-mode    C-f               send-keys -X cursor-right
bind-key    -T copy-mode    C-g               send-keys -X clear-selection
bind-key    -T copy-mode    C-k               send-keys -X copy-end-of-line
bind-key    -T copy-mode    C-n               send-keys -X cursor-down
bind-key    -T copy-mode    C-p               send-keys -X cursor-up
bind-key    -T copy-mode    C-r               command-prompt -i -I "#{pane_search_string}" -p "(search up)" "send -X search-backward-incremental \"%%%\""
bind-key    -T copy-mode    C-s               command-prompt -i -I "#{pane_search_string}" -p "(search down)" "send -X search-forward-incremental \"%%%\"
"
bind-key    -T copy-mode    C-v               send-keys -X page-down
bind-key    -T copy-mode    C-w               send-keys -X copy-selection-and-cancel
bind-key    -T copy-mode    Escape            send-keys -X cancel
bind-key    -T copy-mode    Space             send-keys -X page-down
bind-key    -T copy-mode    ,                 send-keys -X jump-reverse
bind-key    -T copy-mode    ;                 send-keys -X jump-again
bind-key    -T copy-mode    F                 command-prompt -1 -p "(jump backward)" "send -X jump-backward \"%%%\""
bind-key    -T copy-mode    N                 send-keys -X search-reverse
bind-key    -T copy-mode    R                 send-keys -X rectangle-toggle
bind-key    -T copy-mode    T                 command-prompt -1 -p "(jump to backward)" "send -X jump-to-backward \"%%%\""
bind-key    -T copy-mode    f                 command-prompt -1 -p "(jump forward)" "send -X jump-forward \"%%%\""
bind-key    -T copy-mode    g                 command-prompt -p "(goto line)" "send -X goto-line \"%%%\""
bind-key    -T copy-mode    n                 send-keys -X search-again
bind-key    -T copy-mode    q                 send-keys -X cancel
bind-key    -T copy-mode    t                 command-prompt -1 -p "(jump to forward)" "send -X jump-to-forward \"%%%\""
bind-key    -T copy-mode    MouseDown1Pane    select-pane
bind-key    -T copy-mode    MouseDrag1Pane    select-pane ; send-keys -X begin-selection
bind-key    -T copy-mode    MouseDragEnd1Pane send-keys -X copy-selection-and-cancel
bind-key    -T copy-mode    WheelUpPane       select-pane ; send-keys -X -N 5 scroll-up
bind-key    -T copy-mode    WheelDownPane     select-pane ; send-keys -X -N 5 scroll-down
bind-key    -T copy-mode    DoubleClick1Pane  select-pane ; send-keys -X select-word
bind-key    -T copy-mode    TripleClick1Pane  select-pane ; send-keys -X select-line
bind-key    -T copy-mode    Home              send-keys -X start-of-line
bind-key    -T copy-mode    End               send-keys -X end-of-line
bind-key    -T copy-mode    NPage             send-keys -X page-down
bind-key    -T copy-mode    PPage             send-keys -X page-up
bind-key    -T copy-mode    Up                send-keys -X cursor-up
bind-key    -T copy-mode    Down              send-keys -X cursor-down
bind-key    -T copy-mode    Left              send-keys -X cursor-left
bind-key    -T copy-mode    Right             send-keys -X cursor-right
bind-key    -T copy-mode    M-1               command-prompt -N -I 1 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-2               command-prompt -N -I 2 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-3               command-prompt -N -I 3 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-4               command-prompt -N -I 4 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-5               command-prompt -N -I 5 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-6               command-prompt -N -I 6 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-7               command-prompt -N -I 7 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-8               command-prompt -N -I 8 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-9               command-prompt -N -I 9 -p (repeat) "send -N \"%%%\""
bind-key    -T copy-mode    M-<               send-keys -X history-top
bind-key    -T copy-mode    M->               send-keys -X history-bottom
bind-key    -T copy-mode    M-R               send-keys -X top-line
bind-key    -T copy-mode    M-b               send-keys -X previous-word
bind-key    -T copy-mode    M-f               send-keys -X next-word-end
bind-key    -T copy-mode    M-m               send-keys -X back-to-indentation
bind-key    -T copy-mode    M-r               send-keys -X middle-line
bind-key    -T copy-mode    M-v               send-keys -X page-up
bind-key    -T copy-mode    M-w               send-keys -X copy-selection-and-cancel
bind-key    -T copy-mode    M-{               send-keys -X previous-paragraph
bind-key    -T copy-mode    M-}               send-keys -X next-paragraph
bind-key    -T copy-mode    M-Up              send-keys -X halfpage-up
bind-key    -T copy-mode    M-Down            send-keys -X halfpage-down
bind-key    -T copy-mode    C-Up              send-keys -X scroll-up
bind-key    -T copy-mode    C-Down            send-keys -X scroll-down
bind-key    -T copy-mode-vi C-b               send-keys -X page-up
bind-key    -T copy-mode-vi C-c               send-keys -X cancel
bind-key    -T copy-mode-vi C-d               send-keys -X halfpage-down
bind-key    -T copy-mode-vi C-e               send-keys -X scroll-down
```