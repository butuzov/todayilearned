#!/usr/bin/env python3

'''
> Looks for readme.ipynb and convert them to readme.md
'''

import sys

from os import walk
from os.path import abspath
import re

from subprocess import run


def main():
    """ General crawler """
    root = abspath(".")
    for parent, _, files in walk(root):
        for readme in [f for f in files if f == "readme.ipynb"]:
            readme_dir = parent.replace(root+'/', "");
            # args = [ "git", "ls-files", "--error-unmatch",
            #                 f"{readme_dir}/{readme}"]

            # cp = run(args, capture_output=True)ipn

            # # skip untracked files
            # if cp.returncode == 1:
            #     continue


            # # exit if error?
            # if cp.returncode > 1 :
            #     print(cp.stderr.decode(), file=sys.stderr)
            #     sys.exit(cp.returncode)

            # parse readme.ipynb into readme.md
            cp = run( ["jupyter", "nbconvert", f"{readme_dir}/{readme}", "--to", \
                    "markdown"], capture_output=True)

            # exit if error?
            if cp.returncode >= 1 :
                print(cp.stderr.decode(), file=sys.stderr)
                sys.exit(cp.returncode)


            print(f"Converting {readme_dir}/{readme}")

            readme = "readme.md";
            contents = ""
            with open(f"{parent}/{readme}") as r:
                contents = "".join(r.readlines())


            with open(f"{readme_dir}/{readme}", "w") as r:
                r.write(converter(contents))
                print(f"Converted {readme_dir}/{readme}")



### Converter functionality description

def clean_python_hide_block(c):
    """
        Removes python code block with first line #hide
    """
    pattern = re.compile('(```python\n(.*?)```)', re.MULTILINE|re.DOTALL)
    for m in re.findall(pattern, c):
        if m[1].startswith('#hide'):
            c = c.replace(m[0], "", 1)
    return c



def clean_value(v):
    """
        We ned to deal with multiline values, we can avoid thism but jupyter
        will fomat it to multiline too.
    """
    if "\n" not in v:
        return v

    new_v = str()
    for line in v.split("\n"):
        if line.strip() == "":
            continue

        if line[0:8] == " "*8:
            new_v +="\n"+ line[7:]
        else:
            new_v = line

    return new_v


def clean_icecream_normilize(c):
    ic_group  = re.findall(re.compile('^(\s*ic\((.*)\))', re.MULTILINE), c)
    ic_output = re.findall(re.compile('^(\s{1,}(ic>) (.*?):(.+?))\n\n', re.MULTILINE|re.DOTALL), c)

    matches = []
    for m in ic_output:
        c = c.replace(m[0], "\n")
        for s in m[0].split("\n"):
            if (s != "" and len(matches) == 0) or s[4:7] == 'ic>':
                matches.append(s)
            elif len(matches) > 0:
                matches[-1] += "\n" + s

    matches_filtred = []
    pattern = re.compile('^\s{1,}(ic>) (.*?): (.+)', re.MULTILINE|re.DOTALL)
    for m in matches:
        _, key, val = re.findall(pattern, m)[0]
        matches_filtred.append((key, clean_value(val)))


    # actual cleanup > last execution output
    for i in matches_filtred:
        c = c.replace(" "*3+i[1], "")

    # print(i)\n>output instead of ic>
    for k, p in enumerate(ic_group):

        pre = re.match(re.compile("( {1,})", re.DOTALL), p[0])
        pre = pre[0] if pre is not None else ""
        v = str()
        for n in matches_filtred[k:]:
            if n[0] == p[1]:
                r = n[1].replace("\n" + " "*5, " ")
                v += "{2}print({0})\n{2}output > {1}\n".format(p[1], r, pre)

        c = c.replace(p[0], v[:-1], 1)

    return c


def clean_empty_code_blocks(c):
    """
    Removes empty code blocks
    """
    regexp = re.compile('(```(.*?)\n(.*?)```)', re.MULTILINE|re.DOTALL)
    for m in regexp.findall(c):
        if len(m[2].strip()) == 0:
            c = c.replace(m[0]+"\n", "")
    return c

def clean_3_lines(c):
    """ simple. simple and silly code. """

    if "\n \n" in c:
        return clean_3_lines(c.replace("\n \n", "\n\n", 1))

    if "\n\n\n" in c:
        return clean_3_lines(c.replace("\n\n\n", "\n\n", 1))
    return c


def clean_empty_block(c):
    """
        Cleans empty blocks (### headings i am leaving for placeholding)
    """
    code = []
    for m in re.findall(re.compile("(```.*?```)", re.DOTALL), c):
        code.append(m)
        c = c.replace(m, f"__code[{len(code)-1}]__")

    # print(c)

    p = re.compile("(#{1,6}) (.*?)\n(.*?)(#|\Z)", re.MULTILINE|re.DOTALL)
    headers = []
    for h in p.findall(c):
        headers.append({
            'header': "{2}{1} {0}".format( h[1], h[0], ('', '#')[len(headers)>0] ),
            'content': h[2],
            'size': len(h[0]) + (0, 1)[len(headers)>0]
        })

    Replaced = False

    for k, h in enumerate(headers):

        if h['content'].strip() != "":
            continue

        if k < len(headers)-1 and h['size'] >= headers[k+1]['size']:
            c = c.replace( f"{h['header']}\n{h['content']}", "" )
            Replaced = True

        if k == len(headers)-1:
            c = c.replace( f"{h['header']}\n{h['content']}", "" )
            Replaced = True

    for m in re.findall(re.compile("(__code\[(\d{1,})\]__)", re.DOTALL), c):
        c = c.replace(m[0], code[int(m[1])])

    return c if not Replaced else clean_empty_block(c)

def converter(c):
    """ Converts IPython Notebook generated readme to nice readme.md file"""

    c = clean_python_hide_block(c)
    c = clean_empty_code_blocks(c)
    c = clean_icecream_normilize(c)
    c = clean_3_lines(c)
    c = clean_empty_block(c)

    return c


if __name__ == "__main__":
    print("Converting Jupyter Notebook's readmes")
    main()
