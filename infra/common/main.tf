locals {
  meta = {
    team    = "nearsfeed"
    service = "nearsfeed"
    env     = "alpha"
  }
}

module "vpc" {
  source = "../modules/vpc"
  meta   = local.meta

  cidr_block = "10.128.0.0/16"
}

module "route53_zone" {
  source = "../modules/route53_zone"

  zone_name = "nearsfeed.com"
}

module "route53_record_github_validation" {
  source = "../modules/route53_record"

  zone_id = module.route53_zone.zone_id
  domain  = "_github-challenge-nearsfeed.${module.route53_zone.zone_name}"
  type    = "TXT"
  ttl     = "300"
  records = [
    "f85cb7b730",
  ]
}

module "rds" {
  source = "../modules/rds_instance"

  meta   = local.meta
  vpc_id = module.vpc.vpc_id
  sg_ids = [module.vpc.sg_basic_id]
  subnet_ids = [
    module.vpc.subnet_public_a_id,
    module.vpc.subnet_public_b_id,
  ]

  db_password = var.db_password
  db_meta = {
    az                  = data.aws_availability_zone.a.name,
    engine              = "mysql"
    engine_version      = "8.0.20"
    instance_class      = "db.t2.micro"
    volume_size         = 20
    maximum_volume_size = 100
    dbname              = "nearsfeed"
    username            = "grandcanyon"
  }
}

resource "aws_s3_bucket" "nearsfeed" {
  bucket = "nearsfeed-lambda"
  acl    = "private"

  tags = local.meta
}
