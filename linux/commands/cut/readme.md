# `cut` Cut out selected portions of each line of a file

```bash
# Cut is for removing columns. To illustrate, if we only wanted the
# first and third columns.
cut -d, -f 1,3 filename.csv

# To select every column other than the first.
cut -d, -f 2- filename.csv

# >>> >>> >>> >>> In combination with other commands, cut serves as a filter.

# Print first 10 lines of column 1 and 3, where "some_string_value" is present
head filename.csv | grep "some_string_value" | cut -d, -f 1,3

#Finding out the number of unique values within the second column.
cat filename.csv | cut -d, -f 2 | sort | uniq | wc -l

# Count occurences of unique values, limiting to first 10 results
cat filename.csv | cut -d, -f 2 | sort | uniq -c | head
```
