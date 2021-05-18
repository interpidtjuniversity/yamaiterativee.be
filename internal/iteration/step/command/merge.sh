#!/bin/bash
# grpcurl -plaintext  -d '{"appOwner":"interpidtjuniversity","appName":"miniselfop"}' localhost:8000 proto.YaMaHubBranchService.QueryAppAllBranch

grpcurl -plaintext -d $1 $2 $3

