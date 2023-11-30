#!/bin/sh

find ./day* -maxdepth 0 -type d -exec touch {}/input.txt \;

num_pass=$(go test -v ./... | grep -Ec "^--- PASS:")
num_tests=$(grep -r "func Test" ./day* | wc -l)

perc_passed=$(expr $num_pass / $num_tests)

if [ $perc_passed = 1 ]; then
    colour=success
elif [ $perc_passed > 0.5 ]; then
    colour=important
else
    colour=critical
fi

sed -i "s/[0-9]\+\/[0-9]\+-[a-z]\+/$num_pass\/$num_tests-$colour/" README.md
