#!/usr/bin/env bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color


REVIVE_CURRENT_RULES=$(\
  curl -s https://raw.githubusercontent.com/mgechev/revive/master/RULES_DESCRIPTIONS.md | \
    grep  -Eoi '^## ([a-zA-Z0-9\-]{1,})' | \
    sed "s/^##//" | \
    sort | \
    uniq)


REVIVE_CURRENT_FILES=$( \
  ls -1 testdata-revive | \
  xargs -I {} sh -c "basename {} .go" | \
  sort )

# new, not exists in examples
for RULE in $REVIVE_CURRENT_RULES; do
  example=$(printf "testdata-revive/%s.go" $RULE)
  if [[ ! -f $example ]]; then
    printf "${YELLOW} (need): %s${NC}\n" $RULE
  fi
done

# old, no more relevant
for RULE in $REVIVE_CURRENT_FILES; do
  w=$(echo $REVIVE_CURRENT_RULES | grep -w $RULE)
  if [[ -z $w ]]; then
    printf "${RED}(remove): %s${NC}\n" $RULE
  fi
done

if [[ $1 == "-verbose" ]]; then
    for RULE in $REVIVE_CURRENT_FILES; do
        w=$(echo $REVIVE_CURRENT_RULES | grep -w $RULE)
        if [[ ! -z $w ]]; then
            printf "${GREEN}(exits): %s${NC}\n" $RULE
        fi
    done
fi