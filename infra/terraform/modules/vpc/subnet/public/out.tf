output "subnet_id" {
  value = aws_subnet.subnet.id
}

output "rtb_id" {
  value = aws_route_table.rtb.id
}

//// for subnet 'private_nat'
//output "nat_id" {
//  value = aws_nat_gateway.nat.id
//}
