<!-- menu: MacOS -->
<!-- weight: 80 -->

# MacOS

## Keyboard Shortcuts

### Screenshots

* <kbd>⌘</kbd> + <kbd>⇧</kbd> + 4, then &lt;space&gt; window screenshot as result.

### Pasting

* <kbd>⇧</kbd>+<kbd>⌥</kbd>+<kbd>⌘</kbd>+<kbd>V</kbd> - Insert text without formating.


### Windows Resizing

* <kbd>⇧</kbd>+<kbd>⌥</kbd>+ resize - proportional resizing.
* <kbd>⌥</kbd> + resize - resizing a window to resize it from the center.


## Command recipes

### Hide all icons on the Mac desktop

```shell
defaults write com.apple.finder CreateDesktop -bool false
```

### Refresh Icons

```shell
sudo find /private/var/folders/ -name 'com.apple.dock.iconcache' -delete
sudo find /private/var/folders/ -name 'com.apple.iconservices' -delete
sudo rm -r /Library/Caches/com.apple.iconservices.store
```

### Removing from auto load start

```shell
#  Remove Creative Cloud from Menubar
launchctl unload -w /Library/LaunchAgents/com.adobe.AdobeCreativeCloud.plist
```
