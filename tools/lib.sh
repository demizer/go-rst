shopt -s nullglob

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
