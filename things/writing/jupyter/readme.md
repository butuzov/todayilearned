# Jupyter

Short start is ...

```shell
brew install jupyter

# show main pathes
jupyter --paths

# installed (jupyterlab) extensions
jupyter labextension list
# enabling one
jupyter labextension enable @jupyterlab/completer-extension:base-service
@jupyterlab/completer-extension:base-service
```

Jupyter notebook, formerly known as the IPython notebook, is a flexible tool that helps you create readable analyses, as you can keep code, images, comments, formulae and plots together.

- https://www.dataquest.io/blog/jupyter-notebook-tips-tricks-shortcuts/
- https://ipython.readthedocs.io/en/stable/interactive/magics.html

## Kernels

> Important: Select default env (` ~/Library/jupyterlab-desktop/jlab_server`) or create new one (`Manage Python Environments` -> `Environments` -> `Add New One`)

[`kernels`](https://docs.jupyter.org/en/latest/projects/kernels.html) allows to use multiple programming languages while working with notebook. Here is how to isntall kernel to [your notebook](https://github.com/jupyterlab/jupyterlab-desktop/issues/304#issuecomment-1574791938) (this example is about `Jupyter Desktop`)


- open `~/Library/jupyterlab-desktop/jlab_server`, find under it ‎⁨`share⁩/jupyter/⁩kernels/`
- install kernel
- copied installed files into kernel directory

```shell
 $ butuzov  ~/Library/jupyterlab-desktop/jlab_server/share/jupyter/kernels
 > tree
.
├── python3
│   ├── kernel.json
│   ├── logo-32x32.png
│   ├── logo-64x64.png
│   └── logo-svg.svg
└── rust
    ├── kernel.js
    ├── kernel.json
    ├── lint-LICENSE
    ├── lint.css
    ├── lint.js
    ├── logo-32x32.png
    ├── logo-64x64.png
    ├── logo-LICENSE.md
    └── version.txt

2 directories, 13 files
```


[list of kernels](https://github.com/jupyter/jupyter/wiki/Jupyter-kernels)

### `closure`

```python
(prn "Hello World")

> "Hello World"

result >>> nil
```

- https://clojure.org/
- https://s01blog.wordpress.com/2017/12/10/how-to-create-clojure-notebooks-in-jupyter/
- https://github.com/clojupyter/clojupyter


```shell
brew install clojure
brew install leiningen

git clone https://github.com/roryk/clojupyter
cd clojupyter
make
make install
>> ...
>> Created /Users/butuzov/Desktop/Projects/todayilearned/kaizen/writing/jupyter/clojupyter/target/clojupyter-0.3.7-SNAPSHOT.jar
>> Created /Users/butuzov/Desktop/Projects/todayilearned/kaizen/writing/jupyter/clojupyter/target/clojupyter-0.3.7-SNAPSHOT-standalone.jar
>> lein clojupyter install
>> If there are a lot of uncached dependencies this might take a while ...
>> WARNING: parse-boolean already refers to: #'clojure.core/parse-boolean in namespace: omniconf.core, being replaced by: #'omniconf.core/parse-boolean
>> Clojupyter v0.3.7-SNAPSHOT@v0.2.2-86-ge968 - Install local
>>
>>    Installed jar:	    ../clojupyter/target/clojupyter-0.3.7-SNAPSHOT-standalone.jar
>>    Install directory:	~/Library/Jupyter/kernels/clojupyter-0.3.7-snapshotv0.2.2-86-ge968
>>    Kernel identifier:	clojupyter-0.3.7-snapshotv0.2.2-86-ge968
>>
>>    Installation successful.
>>
>> exit(0)
# If required move into different location

mv ~/Library/Jupyter/kernels/clojupyter-0.3.7-snapshotv0.2.2-86-ge968  ~/Library/jupyterlab-desktop/jlab_server/share/jupyter/kernels/clj
code ~/Library/jupyterlab-desktop/jlab_server/share/jupyter/kernels/clj/kernel.json
```

### Rust

```python
println!("Hello World")

> Hello World

result >>> ()
```

- [evcxr](https://github.com/evcxr/evcxr)

```shell
cargo install --locked evcxr_jupyter
evcxr_jupyter --install
>> ...
>> Writing ~/Library/Jupyter/kernels/rust/kernel.json
>> ...
mv ~/Library/Jupyter/kernels/rust ~/Library/jupyterlab-desktop/jlab_server/share/jupyter/kernels/rust
```

### `go`

```python
import  "fmt"

fmt.Sprintln("hi")

result >>> hi
```

- (current) https://github.com/janpfeifer/gonb
- https://github.com/gopherdata/gophernotes

```shell
# https://github.com/gopherdata/gophernotes
go install github.com/gopherdata/gophernotes@v0.7.5

# default jupyter location.
mkdir -p ~/Library/Jupyter/kernels/gophernotes
cd ~/Library/Jupyter/kernels/gophernotes
cp "$(go env GOPATH)"/pkg/mod/github.com/gopherdata/gophernotes@v0.7.5/kernel/*  "."
chmod +w ./kernel.json # in case copied kernel.json has no write permission
mv ~/Library/Jupyter/kernels/gophernotes ~/Library/jupyterlab-desktop/jlab_server/share/jupyter/kernels/go
```

## Typescript

https://github.com/yunabe/tslab

```shell
npm install -g tslab
tslab install --version
tslab install [--python=python3]
# or tslab install
jupyter kernelspec list
```

## Magic (cells)

- https://ipython.org/ipython-doc/3/interactive/magics.html

### Autotime & Timeit

```python
#> pip install git+git://github.com/cpcloud/ipython-autotime -q
%load_ext autotime

> time: 591 µs (started: 2023-11-24 15:21:45 +02:00)
```

```python
from time import sleep
sleep(1)

> time: 1.01 s (started: 2023-11-24 15:21:45 +02:00)
```

```python
%timeit sleep(1)

> 1 s ± 1.45 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
> time: 8.03 s (started: 2023-11-24 15:21:46 +02:00)
```

```python
### top 4 magic commands

- % matplotlib. If you did an online course before, you probably recognize this magic command in combination with the inline parameter. ...
- %load_ext autoreload. This magic command allows you to load the most important extension: autoreload. ...
```

```python
%system

> time: 25.1 ms (started: 2023-11-24 15:22:37 +02:00)

result >>> []
```

```python
%who_ls

> time: 3.96 ms (started: 2023-11-24 15:22:18 +02:00)

result >>> ['sleep']
```

### IPython's 'magic' functions

- `%magic` in cell will show help on magic commands
- `%` short lines, `%%` cells.
- `%paste` paste (dowsn't work on mac)
- `%cpaste`
- `%run` runs script.
- `%time` time exceution
- `%timeit` time exceution
- `%prun` profiler run
- `%mrun` memory run
- `%memit` memory reports.



```
___currentDir___= %pwd
!mkdir -p /content/extensions/
%cd /content/extensions/

!wget -O <<plugin_name>>.py <<plugin_URL>>
%load_ext <<plugin_name>>
#%reload_ext <<plugin_name>> #usefull for plugin development
#%unload_ext <<plugin_name>> #usefull for plugin development

%cd {___currentDir___}
del ___currentDir___
```

## Extensions

```shell
(venv) > pip install jupyter jupyter_contrib_nbextensions
# autotime extension
(venv) > jupyter contrib nbextension install --sys-prefix
(venv) > jupyter nbextension enable codefolding/main
(venv) > jupyter nbextension disable codefolding/main
```


### Some other...


-  [Skip cell (True/False)](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/skip_ext.py)
-  [Push and Pop Directory: This extension allows you to modify the working directory internally to the cell](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/pNpDir_ext.py)
-  [Use variables within the %%writefile context](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/writefileE_ext.py)
-  [This magic command allows you to execute a Python script by specifying the script's name or path as part of the command line arguments. It identifies the script, checks if it exists in the system paths, and then runs it. Similar to %run but finds the script in the system path](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/sys_run_ext.py)
- [This magic command resizes the output cell by specifying height and/or width parameters. It uses argparse to parse the command line arguments, allowing users to control the size of the output cell](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/resize_output_ext.py)
- [Repo Manager: This extension allows to handle GIT repositories using a using a JSON file stored in Google Drive with access tokens](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/repo_ext.py)
- [Code Highlighting](https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/hightlighting_ext.py)

```
!wget -O skip_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/skip_ext.py
!wget -O pNpDir_ext.py  https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/pNpDir_ext.py
!wget -O writefileE_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/writefileE_ext.py
!wget -O sys_run_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/sys_run_ext.py
!wget -O resize_output_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/resize_output_ext.py
!wget -O repo_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/repo_ext.py
!wget -O hightlighting_ext.py https://gist.githubusercontent.com/Mr-McGL/661ae7e50a1cfe7dfd5c0b23216bf0c3/raw/hightlighting_ext.py

%load_ext skip_ext
%load_ext pNpDir_ext
%load_ext writefileE_ext
%load_ext sys_run_ext
%load_ext resize_output_ext
%load_ext repo_ext
%load_ext hightlighting_ext


```

## Customizations

#### Directory

Using custom directory intead `~/.jupyter/`

```shell
export JUPYTER_CONFIG_DIR="/location/of/the/jupyter/home"
```


#### Styles

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

## Software

- Using [Jupyter Desktop (Electron App)](https://github.com/jupyterlab/jupyterlab-desktop)
- VSCode (as editor)

### VSCode

<div class="alert alert-info center">Just in case if you need to open jupyter as json, right click at left sidebar and <b>Open With...</b> <b>Text editor</b></div>

## Fun With Jupyter

- Turn it into web app with https://github.com/LCL-CAVE/manganite/tree/master
