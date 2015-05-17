#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source ./lib.sh

trap 'trap_abort' INT QUIT TERM HUP
trap 'trap_exit' EXIT

usage() {
cat << EOF

Updates README.rst with a table showing the level of completion of the
reStructuredText parser in go-rst.

Usage:
    ${0##*/}
    ${0##*/} -h | --help

Options:
    -h       --help               Display this help and exit.

EOF
}

while true; do
    case "$1" in
        -h|--help) usage; exit 0;;
        *) break;;
    esac
    shift
done

cd $DIR/../tools/progress-dump/

go run main.go --progress-yml ../../progress.yml > /tmp/go-rst-progress-table

cd - > /dev/null

cat README.rst.template | sed -e '/%%PROGRESS_TABLE%%/ {
    r /tmp/go-rst-progress-table
    d
}' | sed -e "s/%%MODIFIED_DATE%%/:Modified: `date +\"%a %b %d %H:%M %Y\"`/g" > README.rst

