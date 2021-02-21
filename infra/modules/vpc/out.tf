# vpc
output "vpc_id" {
  value = aws_vpc.nearsfeed.id
}

output "vpc_arn" {
  value = aws_vpc.nearsfeed.arn
}

output "vpc_ipv4_cidr_block" {
  value = aws_vpc.nearsfeed.cidr_block
}

output "vpc_ipv6_association_id" {
  value = aws_vpc.nearsfeed.ipv6_association_id
}

output "vpc_ipv6_cidr_block" {
  value = aws_vpc.nearsfeed.ipv6_cidr_block
}

output "vpc_main_route_table_id" {
  value = aws_vpc.nearsfeed.main_route_table_id
}

output "vpc_owner_id" {
  value = aws_vpc.nearsfeed.owner_id
}

output "sg_basic_id" {
  value = aws_security_group.basic.id
}

output "sg_members_id" {
  value = aws_security_group.members.id
}

output "igw_main_id" {
  value = aws_internet_gateway.nearsfeed.id
}

output "subnet_public_a_id" {
  value = module.subnet_public["a"].subnet_id
}

output "subnet_public_b_id" {
  value = module.subnet_public["b"].subnet_id
}

output "subnet_private_a_id" {
  value = module.subnet_private["a"].subnet_id
}
output "subnet_private_b_id" {
  value = module.subnet_private["b"].subnet_id
}


output "subnet_private_nat_a_id" {
  value = module.subnet_private_nat["a"].subnet_id
}

output "subnet_private_nat_b_id" {
  value = module.subnet_private_nat["b"].subnet_id
}
