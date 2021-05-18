#!/bin/bash

# pmd.sh miniselfop/src
/root/pmd/bin/run.sh pmd -d $1 -R rulesets/java/quickstart.xml -f text >> pmd.log 2>&1 -no-cache
echo "["
while read record
do
  arr=(${record//:/ })
  file=`eval echo ${arr[0]}`
  line=`eval echo ${arr[1]}`
  key=`eval echo ${arr[2]}`
  info=`eval echo ${arr[@]:3}`
  code=`sed -n ${line}p ${file}`
  data=$(jq -n --arg var1 "$file" --arg var2 "$line" --arg var3 "$key" --arg var4 "$info" --arg var5 "$code" '{file: $var1, line: $var2, key: $var3, info: $var4, code: $var5}')
  echo ${data},
done < pmd.log
echo "{}]"
