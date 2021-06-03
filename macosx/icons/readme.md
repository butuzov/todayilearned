# Icons in Mac

## Replace Icons

![](replace.gif)

## Refresh Icons ([`guide`](https://discussions.apple.com/thread/8441124))

```shell
sudo find /private/var/folders/ -name 'com.apple.dock.iconcache' -delete
sudo find /private/var/folders/ -name 'com.apple.iconservices' -delete
sudo rm -r /Library/Caches/com.apple.iconservices.store
```

## Icons

* Visual Studio Code - `Code.icns`
