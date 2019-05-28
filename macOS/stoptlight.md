# Spotlight

* [Spotlight troubleshooting tips](http://osxdaily.com/2007/02/15/spotlight-wont-work-fix-a-broken-spotlight-menu-with-these-troubleshooting-tips/)
* [apple.stackexchange.com for #spotlight](https://apple.stackexchange.com/questions/tagged/spotlight?sort=votes&pagesize=50)

## Indexing

```bash
# disable indexing
sudo mdutil -a -i off

# enable indexing
sudo mdutil -a -i on

# rebuild index manualy
sudo mdutil -E -i on /
```
