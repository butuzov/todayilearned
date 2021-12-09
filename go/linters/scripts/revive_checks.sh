#!/usr/bin/env bash

set -f

# save linter results
make check-revive 1> tmp.revive.txt


REVIVE_CURRENT_RULES=$(\
  curl -s https://raw.githubusercontent.com/mgechev/revive/master/RULES_DESCRIPTIONS.md | \
    grep  -Eoi '^## ([a-zA-Z0-9\-]{1,})' | \
    sed "s/^##//" | \
    sort | \
    uniq)


RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color



find_in_file() {
    CHECK=$1
    FILE=tmp.revive.txt

    # not found example
    if [[ -z $(grep "$CHECK.go" "$FILE") ]]; then
        echo 3;
        return;
    fi


    if [[ -z $(grep "$CHECK.go" "$FILE"  | grep -i "$CHECK:") ]]; then
        echo 2;
        return
    fi

    echo $(grep "$CHECK.go" "$FILE" | grep -i "$CHECK:" | head -1 );
    return;
}


for RULE in $REVIVE_CURRENT_RULES; do

    RESULT="$(find_in_file $RULE)"
    if [[ $RESULT == 1 ]]; then
    #     # not found in results
        echo "not found in results"
    elif [[ $RESULT == 3 ]]; then
        printf "${RED}[404] %s${NC}\n" $RULE
    elif [[ $RESULT == 2 ]]; then
    #     # not found in file
        printf "${YELLOW}[410] %s${NC}\n" $RULE
    else
        # printf "${GREEN}[200] %s: %s${NC}\n" $RULE "$(echo "$RESULT" | awk 'BEGIN { FS=":" } { print $5 }')"
        foo="1"
    fi

#   example=$(printf "testdata-gocritic/%s.go" $RULE)
#   if [[ ! -f $example ]]; then
#     printf "${YELLOW} (need): %s${NC}\n" $RULE
#   fi
done


if [[ $FAILED == "YES" ]]; then
    printf "%s\n" "----------------------------"
    printf "${RED}  (trigger): %s${NC}\n" "Linter isn't triggered"
    printf "${YELLOW}  (example): %s${NC}\n" "Example isn't working"
    exit 1
fi
