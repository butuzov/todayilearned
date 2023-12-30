#!/bin/bash

for service in services/*/;
  do for v in $service*;
    do
      pushd $v;
      tag=butuzov/demoapp-$(basename $service):$(basename $v)
      docker build -t $tag .;
      docker push $tag;
      popd;
  done
done
