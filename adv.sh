#!/bin/sh
year=${YEAR:-`date "+%Y"`}
if [[ ! -d $year ]]; then
	mkdir $year 2> /dev/null
fi

next=$((`ls $year | sort -n -r | head -1` + 1))

echo "year: $year"
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
    mkdir $year/$next

    if [[ $? == 0 ]]; then
        mv $code $data $test $year/$next
        ls -l $year
    fi

else
    echo "invalid command: $1"
fi
