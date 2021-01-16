# GoTOP

Quit: `q` or `<C-c>`

Process navigation:
  - `k` and `<Up>`: up
  - `j` and `<Down>`: down
  - `<C-u>`: half page up
  - `<C-d>`: half page down
  - `<C-b>`: full page up
  - `<C-f>`: full page down
  - `gg` and <Home>: jump to top
  - `G` and <End>: jump to bottom

Process actions:
  - `<Tab>`: toggle process grouping
  - `dd`: kill selected process or group of processes with SIGTERM (15)
  - `d3`: kill selected process or group of processes with SIGQUIT (3)
  - `d9`: kill selected process or group of processes with SIGKILL (9)

Process sorting:
  - `c`: CPU
  - `m`: Mem
  - `p`: PID

Process filtering:
  - `/`: start editing filter
  - (while editing):
    - `<Enter>`: accept filter
    - `<C-c>` and <Escape>: clear filter

CPU and Mem graph scaling:
  - `h`: scale in
  - `l`: scale out

Network:
  - b: toggle between mbps and scaled bytes per second
