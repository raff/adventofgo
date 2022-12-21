#!/bin/sh
next=$((`ls done | sort -n -r | head -1` + 1))

echo "next: day $next"
echo

code="code${next}.go"
data="data${next}.txt"
test="test${next}.txt"

if [[ "$1" == "" ]]; then
    if [[ -f $code ]]; then
        echo $code already exists.
        echo
    else
        cp advlib/template_go $code
        touch $data
        touch $test
    fi

    ls -l *.*

elif [[ "$1" == "done" ]]; then
    mkdir done/$next

    if [[ $? == 0 ]]; then
        mv $code $data $test done/$next
        ls -l done
    fi

else
    echo "invalid command: $1"
fi
