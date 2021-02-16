locals {
  subnet_public = {
    a = {
      az         = data.aws_availability_zone.a.name,
      cidr_block = "${substr(var.cidr_block, 0, 6)}.0.0/23"
    },
    b = {
      az         = data.aws_availability_zone.b.name,
      cidr_block = "${substr(var.cidr_block, 0, 6)}.8.0/23"
    },
    //    c = {
    //      az         = data.aws_availability_zone.c.name,
    //      cidr_block = "${substr(var.cidr_block, 0, 6)}.16.0/23"
    //    },
    //    d = {
    //      az         = data.aws_availability_zone.d.name,
    //      cidr_block = "${substr(var.cidr_block, 0, 6)}.24.0/23"
    //    },
  }

  subnet_private = {
    a = {
      az         = data.aws_availability_zone.a.name,
      cidr_block = "${substr(var.cidr_block, 0, 6)}.32.0/23"
    },
    b = {
      az         = data.aws_availability_zone.b.name,
      cidr_block = "${substr(var.cidr_block, 0, 6)}.40.0/23"
    },
    //    c = {
    //      az         = data.aws_availability_zone.c.name,
    //      cidr_block = "${substr(var.cidr_block, 0, 6)}.48.0/23"
    //    },
    //    d = {
    //      az         = data.aws_availability_zone.d.name,
    //      cidr_block = "${substr(var.cidr_block, 0, 6)}.56.0/23"
    //    },
  }

  //  subnet_private_nat = {
  //    a = {
  //      az         = data.aws_availability_zone.a.name,
  //      nat_id     = module.subnet_public[a].nat_id
  //      cidr_block = "${substr(var.cidr_block, 0, 6)}.64.0/23"
  //    },
  //    b = {
  //      az         = data.aws_availability_zone.b.name,
  //      nat_id     = module.subnet_public[b].nat_id
  //      cidr_block = "${substr(var.cidr_block, 0, 6)}.72.0/23"
  //    },
  //    c = {
  //      az         = data.aws_availability_zone.c.name,
  //      nat_id     = module.subnet_public[c].nat_id
  //      cidr_block = "${substr(var.cidr_block, 0, 6)}.80.0/23"
  //    },
  //    d = {
  //      az         = data.aws_availability_zone.d.name,
  //      nat_id     = module.subnet_public[d].nat_id
  //      cidr_block = "${substr(var.cidr_block, 0, 6)}.88.0/23"
  //    },
  //  }
}

module "subnet_public" {
  for_each = tomap(local.subnet_public)

  source = "./subnet/public"

  az                     = each.value.az
  subnet_ipv4_cidr_block = each.value.cidr_block
  vpc_id                 = aws_vpc.nearsfeed.id
  igw_main_id            = aws_internet_gateway.nearsfeed.id
  meta                   = var.meta
}

module "subnet_private" {
  for_each = tomap(local.subnet_private)

  source = "./subnet/private"

  az                     = each.value.az
  subnet_ipv4_cidr_block = each.value.cidr_block
  vpc_id                 = aws_vpc.nearsfeed.id
  meta                   = var.meta
}


//module "subnet_private_nat" {
//  for_each = local.subnet_private_nat
//
//  source = "./subnet/private_nat"
//
//  az = each.value.az
//  subnet_ipv4_cidr_block = each.value.cidr_block
//  vpc_id = aws_vpc.nearsfeed.id
//  igw_main_id = aws_internet_gateway.nearsfeed.id
//  meta = var.meta
//}




