
Member=$1

MyIP=$(curl -s ifconfig.co)

# example
# ./scripts/add.sg.myip.sh ${github_nickname}

function add() {
  SecurityGroupName=$1

  SecurityGroupMeta=$(aws ec2 describe-security-groups --filter Name=group-name,Values=${SecurityGroupName} | jq -crM .)
  SecurityGroupId=$(echo "${SecurityGroupMeta}" | jq -crM '.SecurityGroups[0].GroupId')
  SecurityGroupIpPermissions=$(echo $SecurityGroupMeta | jq -crM ".SecurityGroups[0].IpPermissions[] | select( .IpRanges[] | select( .Description == \"${Member}\"))")

  if [ "${SecurityGroupIpPermissions}" == "" ]; then
    echo "이미 제거되었거나, 없었습니다. 새로운 ingess를 등록합니다."
  else
    aws ec2 revoke-security-group-ingress \
      --group-id ${SecurityGroupId} \
      --ip-permissions "${SecurityGroupIpPermissions}" | jq -crM
  fi

  IngressFormat='{"FromPort":0,"ToPort":60000,"IpProtocol":"tcp","IpRanges":[{"CidrIp":"","Description":""}],"Ipv6Ranges":[],"PrefixListIds":[],"UserIdGroupPairs":[]}'

  Ingress=$(echo "${IngressFormat}" | jq -crM ".IpRanges[0].Description=\"${Member}\"")
  Ingress=$(echo "${Ingress}" | jq -crM ".IpRanges[0].CidrIp = \"${MyIP}/32\"")

  aws ec2 authorize-security-group-ingress \
    --group-id ${SecurityGroupId} \
    --ip-permissions ${Ingress} \
    | jq -crM

}

add members
add nearsfeed-alpha