#!/usr/bin/env bash

set -eu

function usage() {
    echo "$0 folder old_expression new_string"
}

if [ "$#" -ne 3 ]; then
    usage
    exit 1
fi

FOLDER="${1}"
OLD_EXP="${2}"
NEW_STR="${3}"

echo "--- Check occurrences ---"
for f in $(ls "${FOLDER}"); do
    if grep "${OLD_EXP}" "${f}" > /dev/null; then
        printf "\033[1;32m${f}\033[0m\n"
        grep --color "${OLD_EXP}" "${f}"
    fi
done

read -p "Press any key to confirm"

for f in $(ls "${FOLDER}"); do
    if grep "${OLD_EXP}" "${f}" > /dev/null; then
        sed -E -i ".old" "s/$OLD_EXP/$NEW_STR/g" "${f}"
    fi
done
