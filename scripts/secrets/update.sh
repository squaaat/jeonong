#!/bin/zsh

DD=/squaaat/squaaat-api/alpha/env

region="ap-northeast-2"
project="jeonong"
app="api"
environment="alpha"
inDir="./"

usage() {
  echo "
Description: Update AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: api)
  -e environment (default: alpha)
  -i inDir (default: ./)
  [-h help]

Example:
  ./scripts/secrets/update.sh -r ap-northeast-2 -a api -e alpha
"
exit 1;
}

while getopts 'r:a:e:h' optname; do
  case "${optname}" in
    h) usage;;
    r) region=${OPTARG};;
    a) app=${OPTARG};;
    e) environment=${OPTARG};;
    i) inDir=${OPTARG};;
    *) usage;;
  esac
done

[ -z "${app}" ] && >&2 echo "Error: -n app required" && usage
[ -z "${environment}" ] && >&2 echo "Error: -m environment required" && usage

echo "/${project}/${app}/${environment}/env"

YML="$(cat ${inDir}./application.${environment}.yml)"
echo "${YML}"

echo "- Output -------------------------------"

echo "${inDir}./application.${environment}.yml)"
aws ssm put-parameter \
  --region ${region} \
  --name "/${project}/${app}/${environment}/application.yml" \
  --type "SecureString" \
  --value "${YML}" \
  --overwrite | jq
