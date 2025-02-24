#!/usr/bin/env bash

if [ ! -f "example_titanic.csv" ]; then
  curl -Ls  https://raw.githubusercontent.com/datasciencedojo/datasets/master/titanic.csv -o example_titanic.csv ;
fi
