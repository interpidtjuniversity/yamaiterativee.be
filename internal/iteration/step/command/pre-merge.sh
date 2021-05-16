#!/bin/bash

#./pre-merge.sh git_repo targetBranch sourceBranch repoName
#./pre-merge.sh http://localhost:3002/interpidtjuniversity/miniselfop.git 5defe44db1_2021_5_10 5defe44db1_2021_5_10_dev miniselfop

function error_result {
    conflict=true
}

if [ ! -d $4 ]; then
    git clone $1 >> /dev/null 2>&1
fi
cd $4
git checkout $3 >> /dev/null 2>&1
git checkout $2 >> /dev/null 2>&1
git merge $3 --no-ff --no-commit >> /dev/null 2>&1 || error_result
git merge --abort >> /dev/null 2>&1

if [ "$conflict" = true ]; then
    echo false
    exit 1
fi

echo true
exit 0
