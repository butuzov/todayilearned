#!/usr/bin/env python3
'''
> Looks for readme.ipynb and convert them to readme.md
'''

import os
import json
import hashlib
from sys import exit
from pathlib import Path

skip_dirs = [".ipynb_checkpoints", ".git"]
target_ext = ".ipynb"


# return pairs (notebook that "suppose to be there" markdown file)
def pairs(root, files):
    return ((
        os.path.join(root, f),
        os.path.join(root, f.replace(target_ext, ".md")),
    ) for f in files if f.endswith(target_ext))


def can_be_converted(files):
    markdown_file = files[1]
    original_file = files[0]
    if not os.path.exists(markdown_file):
        return True

    t1 = os.path.getmtime(original_file)
    t2 = os.path.getmtime(markdown_file)

    return (t2 - t1) > 1 or (t1 - t2) > 1


def filter_out(lines):
    # todo(butuzov): move rejected externaly
    rejected = ["postgresql:"]

    return (l for l in lines
            if not (any((pattern in l for pattern in rejected))))


def cell_markdown(cell) -> str:

    return "".join(cell["source"])


def code_cell(language, cell) -> str:
    code = "".join(cell['source']).rstrip()

    stderr, stdout = "", ""
    data_text, data_html = "", ""
    for out in cell["outputs"]:
        named = "name" in out
        datad = "data" in out

        if named and out["name"] == "stderr":
            stderr = "stderr >>> {0}".format("stderr >>> ".join(
                filter_out(out["text"])))
            stderr = stderr.strip()

        if named and out["name"] == "stdout":
            stdout = "stdout >>> {0}".format("stdout >>> ".join(
                filter_out(out["text"])))
            stdout = stdout.strip()

        if datad and "text/html" in out["data"]:
            data_html = "\n\n" + "".join(filter_out(out["data"]["text/html"]))

        if datad and "image/png" in out["data"]:
            data_html += "\n\n<img src='data:image/png;base64,{0}' />".format(
                out["data"]["image/png"][0:-1])

        if datad and "text/plain" in out["data"]:
            data_text = "result >>> {0}\n".format("result >>> ".join(
                filter_out(out["data"]["text/plain"])))
            data_text = data_text.strip()

    if code == "":
        return ""

    data_text = "" if data_html else data_text

    out_list = []
    if stderr: out_list.append(stderr)
    if stdout: out_list.append(stdout)
    if data_text: out_list.append(data_text)
    out = "\n".join(out_list).strip()
    out = "" if not out else f"\n{out}"

    if '%%sql' in code:
        language = 'sql'
        code = code.replace("%%sql", "")

    if '%sql' in code:
        language = 'sql'
        code = code.replace("%sql", "")

    if out and not out.endswith("\n"):
        out += "\n"

    return "```{0}```{1}".format("\n".join([language,
                                            code.lstrip(), out]), data_html)


def convert(src: str, dst: str):

    old_hash: str = ""
    if Path(dst).is_file():
        old_hash = hashlib.md5(open(dst, 'rb').read()).hexdigest()

    result = []

    with open(src) as f:
        data = json.load(f)
        for cell in data["cells"]:
            c = ""

            if cell["cell_type"] == "markdown":
                c = cell_markdown(cell)
            elif cell["cell_type"] == "code":
                c = code_cell(data["metadata"]["kernelspec"]["language"], cell)
            else:
                print(cell)

            if c != "":
                result.append(c)

    if len(result) == 0:
        return False

    with open(dst, "w") as f:
        f.write("\n\n".join(result))

    return old_hash != hashlib.md5(open(dst, 'rb').read()).hexdigest()


def main():
    """ """

    changed = []

    for root, _, files in os.walk("."):
        if any((d in root for d in skip_dirs)):
            continue

        for pair in (p for p in pairs(root, files) if can_be_converted(p)):
            markdown_file = pair[1]
            original_file = pair[0]

            changed.append(convert(original_file, markdown_file))

    return 1 if any(changed) else 0


if __name__ == "__main__":
    print("Converting Jupyter Notebook's")
    exit(main())
