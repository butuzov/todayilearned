#!/usr/bin/env bash

for repo in $(cat repos.txt)
    do
        name=$(basename $repo);
        echo $name;
        if [[ ! -d "repo-$(basename $repo)" ]]; then
            git clone -q $repo "repo-$(basename $repo)";
        fi
    done
