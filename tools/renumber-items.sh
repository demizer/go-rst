#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source $DIR/lib.sh

trap 'trap_abort' INT QUIT TERM HUP
trap 'trap_exit' EXIT

usage() {
cat << EOF

Used to renumber IDs in items.json files.

Usage:
    ${0##*/} (-f <file | --file <file>)
    ${0##*/} -h | --help

Options:
    -f       --file               Absolute path to the file.
    -h       --help               Display this help and exit.

EOF
}

INPUT_FILE=""

if [[ $# -lt 1 ]]; then
    usage;
    exit 0;
fi

while true; do
    case "$1" in
        -f|--file) shift; INPUT_FILE=$1;;
        -h|--help) usage; exit 0;;
        *) break;;
    esac
    shift
done

if [[ -f "$INPUT_FILE" ]]; then
    awk 'BEGIN{count = 1;}{
    if ($1 == "\"id\":") {
        print "        "$1, count",";
        count++;
    } else
        print $0;
    }' "$INPUT_FILE" > tmp
    mv tmp "$INPUT_FILE"
else
    error "Input file not found!"
fi
