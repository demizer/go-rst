#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source tools/lib.sh

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

replacer() {
    /usr/bin/env python3 << EOF > $1
import re
tmp = """${TEMPLATE}"""
var = """${REPLACEMENT}"""
rxp = re.compile(r"...STATUS.START\n+(.*\n)*\n+...STATUS.END")
print(rxp.sub(r".. STATUS START\n\n{}\n\n.. STATUS END".format(var), tmp, 0))
EOF

}

go run tools/progress-dumper.go > "/tmp/go-rst-progress-table"

# REPLACE status in README.rst
cat /tmp/go-rst-progress-table | grep "README_STATUS" | cut -f2 -d: > "/tmp/go-rst-progress-readme-status"
TEMPLATE="$(cat README.rst)"
REPLACEMENT="$(cat /tmp/go-rst-progress-readme-status)"
replacer /tmp/go-rst-progress-new-README.rst
cat /tmp/go-rst-progress-new-README.rst > README.rst

# REPLACE text in doc/README.rst
cat /tmp/go-rst-progress-table | grep -v "README_STATUS" | sed '/^$/d' > "/tmp/go-rst-progress-readme-table"
TEMPLATE="$(cat doc/README.rst)"
REPLACEMENT="$(cat /tmp/go-rst-progress-readme-table)"
replacer /tmp/go-rst-progress-new-doc-README.rst
cat /tmp/go-rst-progress-new-doc-README.rst > doc/README.rst

# The END
rm -rf /tmp/go-rst-progress-*
