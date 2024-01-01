#!/usr/bin/env bash

# cat "o.json" | jq 'length'
# jq 'length' o.json

printf "# --- %s ----------------\n" "$(date)"
# jq --null-input --argjson test null '.test=$test' | jq -r tostring
# cat o.json | jq -M '.[]'

export j='-M o.json'
# jq '.' $j
echo 1 | jq '[., tojson, tostring]'
