# Jupyter Notebook

Jupyter notebook, formerly known as the IPython notebook, is a flexible tool that helps you create readable analyses, as you can keep code, images, comments, formulae and plots together.

* https://www.dataquest.io/blog/jupyter-notebook-tips-tricks-shortcuts/
* https://ipython.readthedocs.io/en/stable/interactive/magics.html
* https://www.cloudera.com/documentation/data-science-workbench/latest/topics/cdsw_jupyter.html
* https://towardsdatascience.com/the-top-5-magic-commands-for-jupyter-notebooks-2bf0c5ae4bb8

## Kernels

### Clojure

* https://s01blog.wordpress.com/2017/12/10/how-to-create-clojure-notebooks-in-jupyter/
* https://github.com/clojupyter/clojupyter


```bash
brew install clojure
brew install leiningen

git clone https://github.com/roryk/clojupyter
cd clojupyter
make
# use vim to change kerner location to 
# kernelDir:=/Users/butuzov/Py_3.7/share/jupyter/kernels/clojure
make install
```

* Move `clojupyter` bin file to `bin` directory
* add `name` to kernel
* add icons

### Go

```bash
# https://github.com/gopherdata/gophernotes
brew install zeromq
brew install pkg-config
go get -u github.com/gopherdata/gophernotes

# default jupyter location.
mkdir -p ~/Library/Jupyter/kernels/gophernotes
cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* ~/Library/Jupyter/kernels/gophernotes

# in case if it's a virtualenv
mkdir -p $VIRTUAL_ENV/share/jupyter/kernels/gophernotes
cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* $VIRTUAL_ENV/share/jupyter/kernels/gophernotes
```

## Magic Methods

* https://ipython.org/ipython-doc/3/interactive/magics.html

### Autotime 

```bash
# Loading extension
# pip install git+git://github.com/cpcloud/ipython-autotime
```

---

```jupyter
%load_ext autotime
```

### Magic Commands

```python
# IPython's 'magic' functions
%magic
```

#### Additional MAgick Commands

* [Toree Magics](https://github.com/apache/incubator-toree/blob/master/etc/examples/notebooks/magic-tutorial.ipynb)


`%` short lines, `%%` cells.

----

* `%paste` paste (dowsn't work on mac)
* `%cpaste`
* `%run` runs script.
* `%time` time exceution
* `%timeit` time exceution
* `%prun` profiler run
* `%mrun` memory run 
* `%memit` memory reports.


## Extensions

```bash
(venv) > pip install jupyter jupyter_contrib_nbextensions
# autotime extension
(venv) > jupyter contrib nbextension install --sys-prefix
(venv) > jupyter nbextension enable codefolding/main
(venv) > jupyter nbextension disable codefolding/main
```

```python
ls

stdout >>> development/  jupyter-notebook.ipynb  search-engines.md
```

```python
# comment
!ls

stdout >>> development  jupyter-notebook.ipynb  search-engines.md
```


```cell
%load t.py
```

## Customizations

### Config Directory

Using custom directory intead `~/.jupyter/`

```bash
export JUPYTER_CONFIG_DIR="/location/of/the/jupyter/home"
```

### Custom CSS

```css
@import url('https://fonts.googleapis.com/css?family=PT+Serif:700');



/* Change code font */
.CodeMirror pre {
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 10pt;
}

div.output pre{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 10pt;
}

div.output_html td{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 9pt;
}

div.prompt{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 9pt;
}

div.completions select{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 10pt;
}

div.container pre{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 10pt;
}

div.tooltiptext pre{
    font-family: Mensch, 'Fira Code', Monaco;
    font-size: 9pt;
}

div.input_area {
    border-color: rgba(0,0,0,0.10);
    background: rbga(0,0,0,0.5);
}

div.text_cell_render {
    font-family: Georgia, serif;
    font-size: 12pt;
    line-height: 160%;
}

div.text_cell_render code{
    font-family: 'Fira Code', Monaco;
    font-size: 80%;
}


div.text_cell_render h1 { font-size:170%; }
div.text_cell_render h2 { font-size:145%; }
div.text_cell_render h3 { font-size:125%; }
div.text_cell_render h4 { font-size:110%; }
div.text_cell_render h5 { font-size:100%; }
div.text_cell_render h6 { font-size:90%; }

div.text_cell_render h1,
div.text_cell_render h2,
div.text_cell_render h3,
div.text_cell_render h4,
div.text_cell_render h5,
div.text_cell_render h6 {
    font-family: 'PT Serif', serif;
    font-weight: 700;
    font-style:  normal;
}

div.text_cell_render > h1:first-child,
div.text_cell_render > h2:first-child,
div.text_cell_render > h3:first-child,
div.text_cell_render > h4:first-child,
div.text_cell_render > h5:first-child,
div.text_cell_render > h6:first-child {
	min-height:24pt;
	line-height: 24pt;
	margin-top:0;
}


.rendered_html pre,
.rendered_html code {
    font-size: medium;
}

.rendered_html ol, ul {
	margin: 1em 2em;
}

.rendered_html ol ul,
.rendered_html ol ol,
.rendered_html ul ul,
.rendered_html ul ol {
	margin: 0 2em;
	padding-left:0 !important;
}


.rendered_html ol {
    list-style:decimal;
}

.prompt.input_prompt {
    color: rgba(0,0,0,0.5);
}

.cell.command_mode.selected {
    border-color: rgba(0,0,0,0.1);
}

.cell.edit_mode.selected {
    border-color: rgba(0,0,0,0.15);
    box-shadow: 0px 0px 5px #f0f0f0;
    -webkit-box-shadow: 0px 0px 5px #f0f0f0;
}

div.output_scroll {
    -webkit-box-shadow: inset 0 2px 8px rgba(0,0,0,0.1);
    box-shadow: inset 0 2px 8px rgba(0,0,0,0.1);
    border-radius: 2px;
}

#menubar .navbar-inner {
    background: #fff;
    -webkit-box-shadow: none;
    box-shadow: none;
    border-radius: 0;
    border: none;
    font-family: lato;
    font-weight: 400;
}

.navbar-fixed-top .navbar-inner,
.navbar-static-top .navbar-inner {
    box-shadow: none;
    -webkit-box-shadow: none;
    border: none;
}

div#notebook_panel {
    box-shadow: none;
    -webkit-box-shadow: none;
    border-top: none;
}

div#notebook {
    border-top: 1px solid rgba(0,0,0,0.15);
}

#menubar .navbar .navbar-inner,
.toolbar-inner {
    padding-left: 0;
    padding-right: 0;
}

#checkpoint_status,
#autosave_status {
    color: rgba(0,0,0,0.5);
}


/*
	Markdown
*/

.text_cell_render.rendered_html pre {
	margin:0;
	background-color: #f6f8fa;
    border-radius: 3px;
    font-size: 90%;
    line-height: 1.45;
    overflow: auto;
	padding: 8px 8px;
}

.text_cell_render.rendered_html p + pre {
	margin-top: 16px;
}

.text_cell_render.rendered_html pre code {
	background-color: #f6f8fa;
}

span.alert {
    padding:3px 5px;
    font-size:100%;
}
```

## Using Jupyter

### Bash

Its usefull sometimes to run pip directly from jupyter, but its generate a lot of output, so we can pipe it to `/dev/null`

```bash
# silent runs
!pip3 install youtube-dl 2>&1 1>/dev/null
```

```python
!ls -lai

stdout >>> total 32
stdout >>> 5740146 drwxr-xr-x  6 butuzov admin   204 Jun 29 11:28 .
stdout >>> 5738587 drwxr-xr-x 19 butuzov admin   646 May 31 12:21 ..
stdout >>> 5741508 drwxr-xr-x  3 butuzov admin   102 May 30 00:23 .ipynb_checkpoints
stdout >>> 5740804 drwxr-xr-x  3 butuzov admin   102 May 12 15:50 development
stdout >>> 5924839 -rw-r--r--  1 butuzov admin 27325 Jun 29 11:28 jupyter-notebook.ipynb
stdout >>> 5924738 -rw-r--r--  1 butuzov admin  3626 Jun 29 11:07 search-engines.md
```

```python
# getting output of the bash
files = !ls -1
```

```python
# using it
files

result >>> ['development', 'jupyter-notebook.ipynb', 'search-engines.md']
```

### Autocomple Commands

```python
# You can get a help about command by adding ? to the end of comman
len?
```

```python
def demo():
    return 1
```

```python
# source code
demo??
```

```python
# autocomplese commands
# list().<tab>
# list().c<tab>
# private methods
# list().__<tab>
```

```python
# will work too
# import <tab>
```

```python
# show all objects in namespace "warnings"
*Warning?
```

```python
# list methods
str.*find*?
```

```python
list()._*?
```

```python
list()._*
```

Typing `⇧+Tab` You can get help for function/class/etc... 

### In & Out

```python
import math
```

```python
math.sin(2)

result >>> 0.9092974268256817
```

```python
math.cos(-2)

result >>> -0.4161468365471424
```

```python
# last cell
_+1

result >>> 0.5838531634528576
```

```python
# one more back
__+2

result >>> 1.5838531634528576
```

```python
# last one
___+3

result >>> 2.5838531634528574
```

```python
In

result >>> ['',
result >>>  "get_ipython().run_line_magic('ls', '')",
result >>>  "# comment\nget_ipython().system('ls')",
result >>>  "get_ipython().system('ls -lai')",
result >>>  "# getting output of the bash\nfiles = get_ipython().getoutput('ls -1')",
result >>>  '# using it\nfiles',
result >>>  "# You can get a help about command by adding ? to the end of comman\nget_ipython().run_line_magic('pinfo', 'len')",
result >>>  'def demo():\n    return 1',
result >>>  "# source code\nget_ipython().run_line_magic('pinfo2', 'demo')",
result >>>  '# autocomplese commands\n# list().<tab>\n# list().c<tab>\n# private methods\n# list().__<tab>',
result >>>  '# will work two\n# import <tab>',
result >>>  '# show all objects in namespace "warnings"\nget_ipython().run_line_magic(\'psearch\', \'*Warning\')',
result >>>  "# list methods\nget_ipython().run_line_magic('psearch', 'str.*find*')",
result >>>  "get_ipython().set_next_input('list()._*');get_ipython().run_line_magic('psearch', '_*')",
result >>>  "# IPython's 'magic' functions\nget_ipython().run_line_magic('magic', '')",
result >>>  'import math',
result >>>  'math.sin(2)',
result >>>  'math.cos(-2)',
result >>>  '# last cell\n_+1',
result >>>  '# one more back\n__+2',
result >>>  '# last one\n___+3',
result >>>  'In']
```

```python
# dictionary
Out

result >>> {5: ['development', 'jupyter-notebook.ipynb', 'search-engines.md'],
result >>>  16: 0.9092974268256817,
result >>>  17: -0.4161468365471424,
result >>>  18: 0.5838531634528576,
result >>>  19: 1.5838531634528576,
result >>>  20: 2.5838531634528574,
result >>>  21: ['',
result >>>   "get_ipython().run_line_magic('ls', '')",
result >>>   "# comment\nget_ipython().system('ls')",
result >>>   "get_ipython().system('ls -lai')",
result >>>   "# getting output of the bash\nfiles = get_ipython().getoutput('ls -1')",
result >>>   '# using it\nfiles',
result >>>   "# You can get a help about command by adding ? to the end of comman\nget_ipython().run_line_magic('pinfo', 'len')",
result >>>   'def demo():\n    return 1',
result >>>   "# source code\nget_ipython().run_line_magic('pinfo2', 'demo')",
result >>>   '# autocomplese commands\n# list().<tab>\n# list().c<tab>\n# private methods\n# list().__<tab>',
result >>>   '# will work two\n# import <tab>',
result >>>   '# show all objects in namespace "warnings"\nget_ipython().run_line_magic(\'psearch\', \'*Warning\')',
result >>>   "# list methods\nget_ipython().run_line_magic('psearch', 'str.*find*')",
result >>>   "get_ipython().set_next_input('list()._*');get_ipython().run_line_magic('psearch', '_*')",
result >>>   "# IPython's 'magic' functions\nget_ipython().run_line_magic('magic', '')",
result >>>   'import math',
result >>>   'math.sin(2)',
result >>>   'math.cos(-2)',
result >>>   '# last cell\n_+1',
result >>>   '# one more back\n__+2',
result >>>   '# last one\n___+3',
result >>>   'In',
result >>>   '# dictionary\nOut']}
```

```python
_

result >>> ['',
result >>>  "get_ipython().run_line_magic('ls', '')",
result >>>  "# comment\nget_ipython().system('ls')",
result >>>  "get_ipython().system('ls -lai')",
result >>>  "# getting output of the bash\nfiles = get_ipython().getoutput('ls -1')",
result >>>  '# using it\nfiles',
result >>>  "# You can get a help about command by adding ? to the end of comman\nget_ipython().run_line_magic('pinfo', 'len')",
result >>>  'def demo():\n    return 1',
result >>>  "# source code\nget_ipython().run_line_magic('pinfo2', 'demo')",
result >>>  '# autocomplese commands\n# list().<tab>\n# list().c<tab>\n# private methods\n# list().__<tab>',
result >>>  '# will work two\n# import <tab>',
result >>>  '# show all objects in namespace "warnings"\nget_ipython().run_line_magic(\'psearch\', \'*Warning\')',
result >>>  "# list methods\nget_ipython().run_line_magic('psearch', 'str.*find*')",
result >>>  "get_ipython().set_next_input('list()._*');get_ipython().run_line_magic('psearch', '_*')",
result >>>  "# IPython's 'magic' functions\nget_ipython().run_line_magic('magic', '')",
result >>>  'import math',
result >>>  'math.sin(2)',
result >>>  'math.cos(-2)',
result >>>  '# last cell\n_+1',
result >>>  '# one more back\n__+2',
result >>>  '# last one\n___+3',
result >>>  'In',
result >>>  '# dictionary\nOut',
result >>>  '_']
```

## Hilighting Things

Using `<div>` tag

<div class="alert alert-danger">ERROR: </div>

```
<div class="alert alert-danger">ERROR: </div>
```

<div class="alert alert-info">INFO: </div>

```
<div class="alert alert-info">INFO: </div>
```


<div class="alert alert-warning">WARNING: </div>

```
<div class="alert alert-warning">WARNING: </div>
```

<div class="alert alert-success">SUCCESS: </div>

```
<div class="alert alert-success">SUCCESS: </div>
```


---

Using `<span>` tag (slitly changed with `custom.css`)

<span class="alert alert-danger">error</span>, <span class="alert alert-info">info</span>,  <span class="alert alert-warning">warning</span> and <span class="alert alert-success">success</span>