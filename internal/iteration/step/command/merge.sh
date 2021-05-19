#!/bin/bash -e

# '{"userName":"interpidtjuniversity","repository":"miniselfop","sourceBranch":"5defe44db1_2021_5_10_dev","targetBranch":"5f6191c9b6_2021_5_16","mergeInfo":"gergergerg"}' localhost:8000 proto.YaMaHubBranchService/Merge2Branch

cd $8
git checkout $4
git merge $3 -m "$5"
cd ..


data=$(jq -n --arg var1 $1 --arg var2 $2 --arg var3 $3 --arg var4 $4 --arg var5 "$5" '{userName:$var1,repository:$var2,sourceBranch:$var3,targetBranch:$var4,mergeInfo:$var5}')

grpcurl -plaintext -d "$data" $6 $7

