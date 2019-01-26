# `grep` File pattern searcher

![grep by Julia Evans](grep.jpg)

```bash
# search for "HEY" in text.txt and show lines that has it
grep HEY text.txt

# same as above AND highlight searched term
grep --color=auto HEY text.txt

# same as above AND lines surrounding match
grep -C 1 -color=auto HEY text.txt

# same as above AND case insensetive search
grep -C 1 -color=auto -i HEY text.txt

# almost same as above but inverted. shows everything that not match term
grep -i -v HEY text.txt

# we can search in directories as well.
grep -i password password-dn

# we can search in directories as well (recursive search)
grep -R password password-dn

# treat binary as text
grep -a password pass.bin

# safe cleanup local pip folder
python3 -m pip list | awk 'NR>2 {print $1}' | grep -Ev "pip|setuptools|wheel" | xargs -I {} sh -c "python3 -m pip uninstall {} -y"

# getting local ip addresses info
ifconfig | grep inet | grep -v inet6 | awk '{print $2}'
```
