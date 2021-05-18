#!/bin/bash

cd $1
git checkout $3
git merge $2 -m $4
mvn compile