#!/bin/bash
grpcurl -plaintext  -d '{"appOwner":"interpidtjuniversity","appName":"miniselfop"}' localhost:8000 proto.YaMaHubBranchService.QueryAppAllBranch
