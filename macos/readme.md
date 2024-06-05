<!-- menu: MacOS -->
<!-- weight: 80 -->

# MacOS

## Keyboard Shortcuts

Mapping Keyboards Keys
- [using-cmd-on-very-old-pc-keyboard-during-boot](https://apple.stackexchange.com/q/71493/)
- https://karabiner-elements.pqrs.org/

__Actual Keywords__:
- [Logitech MX Keys Mini Minimalist](https://www.amazon.com/dp/B098JPSVKY)
- https://nuphy.com/collections/keyboards/products/air75-v2

### Screenshots

* <kbd>⌘</kbd> + <kbd>⇧</kbd> + 4, then &lt;space&gt; window screenshot as result.

### Pasting

* <kbd>⇧</kbd>+<kbd>⌥</kbd>+<kbd>⌘</kbd>+<kbd>V</kbd> - Insert text without formating.


### Windows Resizing

* <kbd>⇧</kbd>+<kbd>⌥</kbd>+ resize - proportional resizing.
* <kbd>⌥</kbd> + resize - resizing a window to resize it from the center.

### Close ALL windows of same type

* Hold <kbd>⌥</kbd> + close toppest window.


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
