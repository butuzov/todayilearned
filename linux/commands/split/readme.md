# `split` Split a file into pieces


```bash

# We will split our CSV into new_filename every 500 lines
split -l 500 file.csv file_
ls
# file_aaa
# file_aab
# file_aac

# and rename to csv
find . -type f -exec mv '{}' '{}'.csv \;
```
