resource "aws_route_table" "rtb" {
  vpc_id = var.vpc_id

  tags = {
    Name = "private_${var.meta.team}_${var.az}",
  }
}

resource "aws_subnet" "subnet" {
  vpc_id            = var.vpc_id
  availability_zone = var.az

  cidr_block = var.subnet_ipv4_cidr_block
  # ipv6_cidr_block                 = var.subnet_ipv6_cidr_block
  # assign_ipv6_address_on_creation = true

  tags = {
    Name          = "private_nat_${var.meta.team}_${var.az}",
    Team          = var.meta.team,
    AvailableZone = var.az,
  }
}

resource "aws_route" "nat" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = aws_route_table.rtb.id
  nat_gateway_id         = var.nat_id
}

resource "aws_route_table_association" "associate" {
  subnet_id      = aws_subnet.subnet.id
  route_table_id = aws_route_table.rtb.id
}
