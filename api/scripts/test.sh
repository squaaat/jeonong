#!/bin/zsh


testData='
{
  "resource": "/",
  "path": "/swagger/doc.json",
  "httpMethod": "GET"
}
'
testDataFilePath=./temp

echo "${testData}" > ${testDataFilePath}


SLS_DEBUG=* sls invoke local -f api \
  --env J_ENV=alpha --env J_CICD=false \
  --path ${testDataFilePath}

#SLS_DEBUG=* sls invoke -f api \
#  --env J_ENV=alpha --env J_CICD=false \
#  --path ${testDataFilePath}

rm ${testDataFilePath}