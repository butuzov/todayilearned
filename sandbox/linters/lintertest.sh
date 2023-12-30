#!/usr/bin/env bash

export GOPATH=$(pwd)/gopath

for d in repo-*/; do
  pushd $d;
      printf "\n ==== [%s] ========= \n" $(basename $d);
      go mod download -x
      $@
  popd;
done


