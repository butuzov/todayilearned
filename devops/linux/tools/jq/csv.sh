#!/usr/bin/env bash
echo "1231";

# Example 1:
# jq --slurp --raw-input --raw-output \
#    'split("\n") | .[1:] | map(split(",")) |
#     map({
#         "PassengerId": .[0],
#         "Survived": "\(.[1])",
#         "Pclass": .[2],
#       })' \
#   example_titanic.csv


