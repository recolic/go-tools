#!/bin/bash

mountp="$1"

[[ "$mountp" = "" ]] && echo "Usage: $0 path/to/mountpoint" && exit 1

_self_bin_name="$0"
function where_is_him () {
    SOURCE="$1"
    while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
        DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
        SOURCE="$(readlink "$SOURCE")"
        [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
    done
    DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
    echo -n "$DIR"
}

function where_am_i () {
    _my_path=`type -p ${_self_bin_name}`
    [[ "$_my_path" = "" ]] && where_is_him "$_self_bin_name" || where_is_him "$_my_path"
}

targets=`cat "$(where_am_i)/files" | sed 's/\\\\/\//g' | sed 's/^C://g'`

echo "$targets" | while read fl; do
    echo FUCKING $fl...
done


