#!/bin/bash


function putParameter() {
  region=$1
  name=$2
  value="${3}"

  aws ssm put-parameter \
    --region $region \
    --name $name \
    --value "${value}" \
    --type "String" \
    --overwrite | jq -crM ""
}

name=$1
value=$2

putParameter ap-northeast-2 $name "${value}"
putParameter us-east-1 $name "${value}"