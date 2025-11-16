##!/usr/bin/env -S uv run --script

# /// script
# dependencies = [ 'rich' ]
# ///

import asyncio
import pathlib
import re

from rich.console import Console
from rich.table import Table

go_version = re.compile(r"go (.*?)$", re.MULTILINE | re.DOTALL)

def get_go_mod_path(path: pathlib.Path) -> str | None:
    if not path.exists():
        return None

    return str(path)


def get_go_mod_version(path: pathlib.Path) -> str | None:
    if m := go_version.search(path.read_text()):
        return m.group(1)

    return None


async def main():
    root = pathlib.Path(".")
    home = pathlib.Path.home() / "go/pkg/mod"

    table = Table(title="Star Wars Movies")
    table.add_column("Package", justify="right", style="magenta")
    table.add_column("Version", style="cyan")
    table.add_column("Path go.mod", style="cyan")
    table.add_column("Version go.mod", style="cyan")

    goMod = root / "go.mod"
    pattern = re.compile(r"require \((.*?)\)$", re.MULTILINE | re.DOTALL)

    for req in pattern.findall(goMod.read_text()):
        for req in req.strip().splitlines():
            parts = req.split("//")[0].split()

            if len(parts) == 2:
                package, version = parts

                suffixes = package.split("/")

                pathToGoMod = home / f"{'/'.join(suffixes[:-1])}/{suffixes[-1]}@{version}/go.mod"

                path, goVersion = None, None
                if _path := get_go_mod_path(pathToGoMod):
                    path = _path
                    if _version := get_go_mod_version(pathToGoMod):
                        goVersion = _version

                table.add_row(package, version, path, goVersion)

    console = Console()
    console.print(table)


if __name__ == "__main__":
    asyncio.run(main())
