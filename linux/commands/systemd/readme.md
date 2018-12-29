# `systemd` systemd system and service manager


##  History / Starters: `init` process

Startup

* `/sbin/init` -> `/etc/inittab` -> `/etc/rc.d/rc.sysinit` -> `/etc/rc.d/rcX.d/`
  * Run Levels

    0. Halt
    1. Single user mode
    2. Multi user mode (no-networking)
    3. Multi user mode (with networking)
    4. unused
    5. Multi-user (with networking and graphical descktop)
    6. Reboot

  * sample `inittab` contents

    ```bash
    # id:runlevel:action:process
    id:3:initdefault
    si::sysinit:/etc/rc.d/rc.system
    10:0:wait:/etc/rc.d/rc 0
    11:1:wait:/etc/rc.d/rc 1
    12:2:wait:/etc/rc.d/rc 3
    13:3:wait:/etc/rc.d/rc 4
    14:5:wait:/etc/rc.d/rc 5
    15:6:wait:/etc/rc.d/rc 6
    ```

    `rc` stands for run `command`
    `K` - kill
    `S` - start

    ```bash
    ls -l /etc/rc.d/rc2.d
    lrwxrwxrwx. 1 root root 20 Dec  9 19:57 K50netconsole -> ../init.d/netconsole
    lrwxrwxrwx. 1 root root 17 Dec  9 19:57 S10network -> ../init.d/network
    ```

## `chkconfig` Updates and queries runlevel information for system services

```bash
# show what commands will run on particular run-level
chkconfig --list | less

# disable network to work on runLevel 3
chkconfig --level 3 network off

# or on
chkconfig network on

# show particular serice
chkconfig --list network
```

## `service` Run a System V init script

```bash
service --status-all
service network status
```

### `runlevel` Print previous and current System V runlevel

```bash
runlevel
> N 3
```

##  Ubuntu's / Starters: `upstart` Starting Services Asyncchronously and Monitor them

```
/sbin/init -> startup(7) -> /etc/init/rc-sysinit.conf -> telinit(8) -> runlevel(7)
                         -> mountall(8)

```

##  Starters: `systemd`

### `systemd-cgls` Recursively show control group contents

```bash
> systemd-cgls

├─1 /usr/lib/systemd/systemd --switched-root --system --deserialize 21
├─user.slice
│ └─user-1000.slice
│   └─session-4.scope
│     ├─ 4226 sshd: vagrant [priv]
│     ├─ 4229 sshd: vagrant@pts/0
│     ├─ 4230 -bash
│     ├─22447 systemd-cgls
│     └─22448 systemd-cgls
└─system.slice
  ├─sshd.service
  │ └─2453 /usr/sbin/sshd -D -u0
  ├─tuned.service
  │ └─2451 /usr/bin/python2 -Es /usr/sbin/tuned -l -P
  ├─rsyslog.service
  │ └─2449 /usr/sbin/rsyslogd -n
  ├─postfix.service
  │ ├─ 2698 /usr/libexec/postfix/master -w
  │ ├─ 2703 qmgr -l -t unix -u
  │ └─22427 pickup -l -t unix -u
  ├─crond.service
  │ └─1614 /usr/sbin/crond -n
  ├─systemd-logind.service
  │ └─1430 /usr/lib/systemd/systemd-logind
  ├─NetworkManager.service
  │ ├─1429 /usr/sbin/NetworkManager --no-daemon
  │ └─2385 /sbin/dhclient -d -q -sf /usr/libexec/nm-dhcp-helper -pf /var/run/dhclient-eth0.pid -lf /var/lib/NetworkManager/dhclient-5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03-eth0.lea
  ├─gssproxy.service
  │ └─1502 /usr/sbin/gssproxy -D
  ├─polkit.service
  │ └─1425 /usr/lib/polkit-1/polkitd --no-debug
  ├─dbus.service
  │ └─1262 /usr/bin/dbus-daemon --system --address=systemd: --nofork --nopidfile --systemd-activation
  ├─chronyd.service
  │ └─1316 /usr/sbin/chronyd
  ├─rpcbind.service
  │ └─1307 /sbin/rpcbind -w
  ├─auditd.service
  │ └─1092 /sbin/auditd
  ├─systemd-udevd.service
  │ └─1069 /usr/lib/systemd/systemd-udevd
  ├─system-getty.slice
  │ └─getty@tty1.service
  │   └─1613 /sbin/agetty --noclear tty1 linux
  └─systemd-journald.service
    └─1036 /usr/lib/systemd/systemd-journald
```

### systemctl

`systemctl action service.name` sample of command

* `status <target>`
* `enable <target>`
* `disable <target>`
* `start <target>`
* `restart <target>`
* `is-enabled <target>`
* `daemon-reload`
* `stop <target>`
* `isolate <target>`
* `mask <target>`
* `is-active <target>`
* `unmask <target>`

```bash
# remote  execution command
systemctl -H 192.168.0.10 status
```

### journalctl

```bash
# permanent loging (keep logs after reboot)
mkdir -p /var/log/journal
systemd-tmpfiles --create --prefix /var/log/journal
```
`/etc/systemd/journald.conf`
* `Storage=`
  * `auto` - default (and memory, and disk)
  * `persistent` stored in `/var/log/journal` heirarchicaly
  * `volatile` - memory
  * `none` - /dev/null or balckhole.
* `Compress=`
  * `yes` - default option, above certain trashhold to compress items before they written to dick.
  * `no` - no compressing

#### `journalctl` usage

```bash
# show newest first
journalctl -r

# end of page
journalctl -e

# 1- recent entries
journalctl -n 10

# tail -f /var/logs/messages
journalctl -f

# specific unit
journalctl -u -f httpd.service

# format: for the detailes
journalctl -o  verbose

# format: get json back
journalctl -o json-pretty

# debuging journalctl
echo "Yo, systemadministrator" | systemd-cat
journalctl -f -n 10


# messages with text in service catalogue
journalctl -x

# kernel messags
journalctl -k

# since recent boot
journalctl -b

# recorded boot sessions (rutn on persistant storage before)
journalctl --list-boots

# date based queries (ok with today, now, yesterday, etc...)
# 2018-01-01 12:12:12 by default
journalctl --since,--until

# disakusage
journalctl --disk-usage

# rotates the journal
journalctl --rotate
```

### utilities

#### `systemd-analyze`

```bash
systemd-analyze
> Startup finished in 1.171s (kernel) + 2.134s (initrd) + 17.895s (userspace) = 21.201s
```
#### `localectl`

```bash
localectl
>    System Locale: LANG=en_US.UTF-8
>        VC Keymap: us
>       X11 Layout: n/a
```

#### `timedatectl`
```
timedatectl
>       Local time: Sat 2018-12-29 12:31:59 UTC
>   Universal time: Sat 2018-12-29 12:31:59 UTC
>         RTC time: Sat 2018-12-29 12:31:56
>        Time zone: UTC (UTC, +0000)
>      NTP enabled: yes
> NTP synchronized: yes
>  RTC in local TZ: no
>       DST active: n/a

timedatectl list-timezones
> some timezones listed

timedatectl set-timezone "Europe/Kyiv"
```

#### `hostnamectl` - Set Hostname permanently

```bash
# set location
hostnamectl set-location "ukrtelecom, kyiv"
```
#### `systemd-resolve`
```bash
systemd-resolve google.com
```


#### `systemd-inhibit` run untill task completed
```bash
systemd-inhibit wget google.com
```

## [Unit files Scripts](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/sect-managing_services_with_systemd-unit_files)

#### Locations:
  * `/usr/lib/systemd/system` (system ones)
  * `/etc/systemd/system` (drop-in unitfile directory)
  * `/run/systemd/system` (runtime)

```bash
# show all unit files
systemctl list-unit-files

# cat unitfile
systemctl cat <unit>

# create full copy of unitfile for modification
systemctl edit --full <unit>

# unitfiles deltra
systemd-delta
```


sample.service
```ini
[Unit]
Description=Some cool app
Documentation=man:cool.app(7)

[Service]
Type=oneshot
ExecStart=/usr/bin/appcool

[Install]
WantedBy=multi-user.target
```

* https://serverfault.com/questions/690155
* https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/sect-managing_services_with_systemd-unit_files

### Target Types

  * multi-user.target
  * graphical.target
  * rescue.target
  * basic.target
  * sysint.taget

### Timers

In house cron/at replacement for services.

```ini
[Timer]
...
```
https://wiki.archlinux.org/index.php/Systemd/Timers
https://coreos.com/os/docs/latest/scheduling-tasks-with-systemd-timers.html

### Systemctl Containers

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux_atomic_host/7/html/managing_containers/using_systemd_with_containers

# Alternatives
  * OpenRC
  * runit

