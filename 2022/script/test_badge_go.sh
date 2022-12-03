#!/bin/sh

 touch ./day{0,1,2}{0,1,2,3,4,5,6,7,8,9}/input.txt

num_pass=$(go test -v ./... | grep -c PASS:)
num_tests=$(grep -r --exclude-dir=script "func Test" . | wc -l)

perc_passed=$(expr $num_pass / $num_tests)

if [ $perc_passed = 1 ]; then
    colour=success
elif [ $perc_passed > 0.5 ]; then
    colour=important
else
    colour=critical
fi

sed -i "s/[0-9]\+\/[0-9]\+-[a-z]\+/$num_pass\/$num_tests-$colour/" README.md