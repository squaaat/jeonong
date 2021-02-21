#!/bin/zsh

region="ap-northeast-2"
project="nearsfeed"
app="nearsfeed-api"
environment="alpha"
inDir="./"

usage() {
  echo "
Description: Update AWS System Store Manager Parameter
Usage: $(basename $0)
  -r region (default: ap-northeast-2)
  -a app (default: nearsfeed-api)
  -e environment (default: alpha)
  -i inDir (default: ./)
  [-h help]

Example:
  ./scripts/secrets/update.sh -r ap-northeast-2 -a nearsfeed-api -e alpha
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

YML="$(cat ${inDir}./application.${app}.${environment}.yml)"
echo "${YML}"

echo "- Output -------------------------------"
echo "region:${region} | path: /${project}/${app}/${environment}/application.yml"
aws ssm put-parameter \
  --region ${region} \
  --name "/${project}/${app}/${environment}/application.yml" \
  --type "SecureString" \
  --value "${YML}" \
  --overwrite | jq
