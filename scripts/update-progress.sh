#!/bin/bash

shopt -s nullglob

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

unset OFF BOLD BLUE GREEN RED YELLOW
OFF="\e[1;0m"
BOLD="\e[1;1m"
BLUE="${BOLD}\e[1;34m"
GREEN="${BOLD}\e[1;32m"
RED="${BOLD}\e[1;31m"
YELLOW="${BOLD}\e[1;33m"
readonly OFF BOLD BLUE GREEN RED YELLOW

plain() {
    local mesg=$1; shift
    printf "${mesg}\n" "$@" >&2
}

msg() {
    local mesg=$1; shift
    printf "${GREEN}####${OFF}${BOLD} ${mesg}${OFF}\n" "$@" >&2
}

msg2() {
    local mesg=$1; shift
    printf "${BLUE}  ##${OFF} ${mesg}\n" "$@" >&2
}

warning() {
    local mesg=$1; shift
    printf "${YELLOW}#### WARNING:${OFF}${BOLD} ${mesg}${OFF}\n" "$@" >&2
}

error() {
    local mesg=$1; shift
    printf "${RED}#### ERROR:${OFF}${BOLD} ${mesg}${OFF}\n" "$@" >&2
}

cleanup() {
    [[ $1 ]] && exit $1
}

abort() {
    echo
    msg 'Aborted'
    cleanup 0
}

trap_abort() {
    trap - EXIT INT QUIT TERM HUP
    abort
}

trap_exit() {
    trap - EXIT INT QUIT TERM HUP
    cleanup
}

trap 'trap_abort' INT QUIT TERM HUP
trap 'trap_exit' EXIT

while true; do
    case "$1" in
        -h|--help) usage; exit 0;;
        *) break;;
    esac
    shift
done

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

cd $DIR/../tools/progress-dump/

go run main.go --progress-yml ../../progress.yml > /tmp/go-rst-progress-table

cd - > /dev/null

cat README.rst.template | sed -e '/%%PROGRESS_TABLE%%/ {
    r /tmp/go-rst-progress-table
    d
}' | sed -e "s/%%MODIFIED_DATE%%/:Modified: `date +\"%a %b %d %H:%M %Y\"`/g" > README.rst

