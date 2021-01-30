#!/bin/zsh

# DD=/squaaat/squaaat-api/alpha/env

region="ap-northeast-2"
project="jeonong"
app="api"
environment="alpha"

usage() {
  echo "
Description: Create AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: api)
  -e environment (default: alpha)
  [-h help]

Example:
  ./scripts/secrets/create.sh -r ap-northeast-2 -a squaaat-api -e alpha
"
exit 1;
}

while getopts 'r:a:e:h' optname; do
  case "${optname}" in
    h) usage;;
    r) region=${OPTARG};;
    a) app=${OPTARG};;
    e) environment=${OPTARG};;
    *) usage;;
  esac
done

[ -z "${app}" ] && >&2 echo "Error: -n app required" && usage
[ -z "${environment}" ] && >&2 echo "Error: -m environment required" && usage

echo "/${project}/${app}/${environment}/application.yml"

echo "- Output -------------------------------"

aws ssm put-parameter \
  --region ${region} \
  --name "/${project}/${app}/${environment}/application.yml" \
  --type "SecureString" \
  --value "version: 1" \
  --tags \
    Key=project,Value=${project} \
    Key=app,Value=${app} \
    Key=project,Value=${environment} \
  --no-overwrite | jq



#aws ssm delete-parameter \
#  --region ${region} \
#  --name "/${project}/${app}/${environment}/env" | jq

#aws ssm delete-parameter \
#  --region ap-northeast-2 \
#  --name "/squaaat/squaaat-api/alpha/env" | jq