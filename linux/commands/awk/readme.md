# `awk` Pattern-directed scanning and processing language

  * [AWK Tutorial][awk-tutorial]

  [awk-tutorial]: http://www.grymoire.com/Unix/Awk.html

![awk by Julia Evans](awk.jpg)

### Vars

   Keyword | Meaning
-----------|---------
  FS       | The Input Field Separator Variable
  OFS      | The Output Field Separator Variable
  NF       | The Number of Fields Variable
  NR       | The Number of Records Variable
  RS       | The Record Separator Variable
  ORS      | The Output Record Separator Variable
  FILENAME | The Current Filename Variable

### Examples

```bash
# 10 Little Soldiers - Causes of Deathes
cat And_Then_There_Were_None.txt | awk '(NR-2)%3 == 0 {print $0}'
> One choked his little self, and then there were nine.
> One overslept himself, and then there were eight.
> One said he'd stay there, and then there were seven.
> One chopped himself in half, and then there were six.
> A bumble-bee stung one, and then there were five.
> One got in chancery, and then there were four.
> A red herring swallowed one, and then there were three.
> A big bear hugged one, and then there were two.
> One got frizzled up, and then there was one.
> He went and hanged himself and then there were none.
```

Simple `ansible` inventory generator

```bash
#!/usr/bin/env bash
DIR=$(pwd | sed "s/\//\\\\\//g")

echo "[all:vars]"
echo "ansible_connection=ssh"
echo "ansible_user=vagrant"
echo "ansible_host=127.0.0.1"
echo "ansible_ssh_common_args='-o StrictHostKeyChecking=no'"
echo ""

vagrant ssh-config| \
    grep -iE "key|host|port" | \
    grep -ivE "strict|hosts" | \
    sed "s/^  //" | \
    sed "s/$DIR\///" | \
    awk '{
        row= ( NR/4 == int(NR/4)) ? NR/4 : int(NR/4)+1;
        if ( array[row] == "" ) {
            array[row] = $2
        } else {
            array[row]= array[row] " " $2
        }
    }
    END {
        for( i=1; i<= row; i++){
            print(  array[i] )
        }
    }
    ' | \
    sed "s/\"//g" | \
    awk '{print $1, "ansible_port="$3, "ansible_private_key_file="$4}' | \
    awk '{
        split( $1, a, /-/)
        if ( groups[a[1]] == "" ){
            groups[a[1]] = a[1]
            printf("\n[%s]\n", a[1])
        }
        print $1, $2, $3
    }'
```
