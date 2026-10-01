# crontab

## Schedule Examples

## macOS

```shell
export EDITOR=nano
export VISUAL=nano

# List all tasksa
crontab -l

# Edit Tasks 
crontab -e 

# Re-add your crontab
crontab ~/crontab.txt

# Remove the current crontab
crontab -r

```

## Shedules Examples

```
0 */2 * * *       # Every 2 hours
0 *   * * *       # Once an hour
0 9   * * *       # At 09:00 every day
0 9 * * 1-5       # Run at 9 AM on weekdays only
0 9 * * MON       # Run at 9 AM only on Mondays  
0 9 5,15,25 * *   # 1th,15th, 25th days...
```


## Test Crontab

```
# Add this test job that runs every minute
* * * * * echo "Cron is working! $(date)" >> ~/Desktop/cron_test.log
```

## Reading 
- https://www.cyberciti.biz/faq/linux-unix-osx-crontab-usrbinvi-exited-with-status-1/
- https://dev.to/trueqap/how-to-run-cron-on-macos-in-2025-a-complete-guide-2b8e