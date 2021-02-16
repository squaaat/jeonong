resource "aws_route_table" "rtb" {
  vpc_id = var.vpc_id

  tags = {
    Name = "public_${var.meta.team}_${var.az}",
  }
}

resource "aws_route" "igw" {
  route_table_id         = aws_route_table.rtb.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = var.igw_main_id
}

resource "aws_subnet" "subnet" {
  vpc_id            = var.vpc_id
  availability_zone = var.az

  cidr_block = var.subnet_ipv4_cidr_block
  # ipv6_cidr_block                 = var.subnet_ipv6_cidr_block
  # assign_ipv6_address_on_creation = true

  tags = {
    Name          = "public_${var.meta.team}_${var.az}",
    Team          = var.meta.team,
    AvailableZone = var.az,
  }
}

# resource "aws_eip" "eip" {
#   vpc = true
# }

# resource "aws_nat_gateway" "nat" {
#   allocation_id = aws_eip.eip.id
#   subnet_id     = aws_subnet.subnet.id

#   tags = {
#     Name = "${var.meta.crew}-NAT-GW-${var.az}"
#   }
# }

resource "aws_route_table_association" "associate" {
  subnet_id      = aws_subnet.subnet.id
  route_table_id = aws_route_table.rtb.id
}
