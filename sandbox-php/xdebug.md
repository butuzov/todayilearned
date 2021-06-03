## `xdebug`

### Install

```bash
# on mac
pecl install xdebug
```

### Settings
```ini
; remove filepathes from debug output
ini_set( 'xdebug.overload_var_dump', 1 );
; php 8 , turn lines reporting
xdebug.mode=off
```