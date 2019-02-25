

# Python Tooling

## PEP8 

**Summary**

* Use `space`, no `tab`. Always 4.
* Lines should be no more then 79 characters in length or less, if you need to
  - make string shorter simplifieing meaningfull variables.
  - add `\` to split code to the new line.
  - continuations of long expressions, should be added by one indentation (4 spaces) 
    of their normal indentation level.
* Functions and Classes should be represented by two blank lines.
* Class methods should be separated by one blank lines.
* don't put spaces around list indexes, function calls, or keyword assignments.
* Put one, and only one space before and after variable assignment.
* etc... etc...

**Readme more**:
* https://pep8.org/
* https://www.python.org/dev/peps/pep-0008/
* http://pep8online.com/

## @pyflackes

Code analyzer 

## @pylint

Shows us a styling and logical errors

## @codecov

Code coverage

## @vulture

Dead code

## `flake8`

```bash
git init
git add example.py
git commit -m "first commit"
> ....

python -m flake8 --install-hook git
git config --bool flake8.strict true

echo "# new comment" >> example.py
git commit -m "second commit"
> error messages...
```

## `venv` or `virtualenv`

```bash
# Python 2.7 
$ sudo python -m pip install virtualenv

# Create Virtual Env
$ python -m virtualenv Py2.7

# Activate Environment
$ . Py2.7/bin/activate 

(Py2.7) $ python -V
> Python 2.7.10

# Deactivate Environment
$ deactivate 

# Python 3.6 (or any else)
$ python3 -m venv Py3.7

# Activate Environment
$ . Py3.7/bin/activate
(Py3.7) $ python -V
> Python 3.7.2

# Check ${VIRTUAL_ENV}
$ echo "${VIRTUAL_ENV}"
> /Users/butuzov/Desktop/Py3.7

# Deactivate Environment
$ deactivate 
```

## Jupyter Notebook

https://www.dataquest.io/blog/jupyter-notebook-tips-tricks-shortcuts/

### Extensions

```bash
(venv) > pip install install jupyter jupyter_contrib_nbextensions
(venv) > pip install git+git://github.com/cpcloud/ipython-autotime
(venv) > jupyter contrib nbextension install --sys-prefix
(venv) > jupyter nbextension enable codefolding/main
(venv) > jupyter nbextension disable codefolding/main
```

### Nice extensions

 1. jupyter nbextension enable execute_time/ExecuteTime
 1. jupyter nbextension enable collapsible_headings/main 

### Loading Extensions into cell

```
# Loading extension (autotime extensions)
# pip install git+git://github.com/cpcloud/ipython-autotime
%load_ext autotime
```

### Running Bash/Pip

"!" added at the beging is running bash

```python
!ls -la
```

    total 8
    drwxr-xr-x  4 butuzov staff  128 Mar  3 17:25 .
    drwxr-xr-x 17 butuzov staff  544 Mar  3 17:18 ..
    drwxr-xr-x  3 butuzov staff   96 Mar  3 16:36 .ipynb_checkpoints
    -rw-r--r--  1 butuzov staff 7672 Mar  3 17:25 readme.ipynb

```python
!pip install --upgrade pip
```

    Requirement already up-to-date: pip in /Users/butuzov/Desktop/Py3.7/lib/python3.7/site-packages (19.0.3)

```python
# surrpress output
!pip install --upgrade pip 2>&1 1>/dev/null
```

### Getting Help

type in code cell function+question sighn e.g. `print?` and press Enter

```python
print?
```

