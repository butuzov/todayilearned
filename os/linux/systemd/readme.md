# systemd

## Service

```
# /etc/systemd/system/myservice.service
[Unit]
Description=My Custom Service
After=network.target

[Service]
ExecStart=/usr/local/bin/myscript.sh
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

## Commands

```shell
sudo systemctl daemon-reload    # reload daemon
sudo systemctl start myservice   # start
sudo systemctl enable myservice  # enable
sudo systemctl status myservice  # status
sudo systemctl restart myservice #
sudo systemctl stop myservice    #
sudo systemctl disable myservice #
```

• Відключення автозапуску:  `
• Перезавантаження без повного stop/start:  `sudo systemctl reload myservice`  (за наявності ExecReload).
• Маскування (заборона запуску навіть вручну):  `sudo systemctl mask myservice` 
• Перегляд стану:  `sudo systemctl status myservice'

логи сервісу
Для перегляду логів скористайтесь командою journalctl:
• Усі логи сервісу:  journalctl -u myservice 
• Логи в реальному часі:  journalctl -u myservice -f 
• Логи за певний період:  journalctl -u myservice --since "2025-11-01" 
• Логи поточного сеансу запуску:  journalctl -u myservice -b 
journalctl дозволяє зручно фільтрувати логи за часом, сер

```
[Unit]
Description=RADA + service
After=syslog.target
After=network.target

[Service]
ExecStart=/home/radaruser/radaplus/wandering_willow/bin/Rada
Environment="LD_LIBRARY_PATH=/home/radaruser/radaplus/wandering_willow/bin"
Restart=always
User=radaruser
Group=radaruser

[Install]
WantedBy=multi-user.target
```
