variable "meta" {
  type = object({
    service = string
    env     = string
  })
}

variable "sg_ids" {
  type = list(string)
}

variable "subnet_ids" {
  type = list(string)
}

variable "db_password" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable db_meta {
  type = object({
    az                  = string
    engine              = string
    engine_version      = string
    volume_size         = number
    maximum_volume_size = number
    instance_class      = string
    dbname              = string
    username            = string
  })
}

resource "aws_db_instance" "db" {
  identifier = "${var.meta.service}-db-mysql"

  allocated_storage     = var.db_meta.volume_size
  max_allocated_storage = var.db_meta.maximum_volume_size
  storage_type          = "gp2"
  engine                = var.db_meta.engine
  engine_version        = var.db_meta.engine_version
  instance_class        = var.db_meta.instance_class
  name                  = var.db_meta.dbname
  username              = var.db_meta.username
  password              = var.db_password

  allow_major_version_upgrade = true

  availability_zone = var.db_meta.az

  final_snapshot_identifier = false # for test
  skip_final_snapshot       = true  # for test
  deletion_protection       = false # for test
  enabled_cloudwatch_logs_exports = [
    "error", "general", "slowquery",
  ]

  vpc_security_group_ids = [aws_security_group.sg.id]
  db_subnet_group_name   = aws_db_subnet_group.subnets.id

  tags = {
    Name        = "${var.meta.service}-db-mysql"
    Service     = var.meta.service
    Environment = var.meta.env
  }
}

output "endpoint" {
  value = aws_db_instance.db.endpoint
}

output "arn" {
  value = aws_db_instance.db.arn
}

output "id" {
  value = aws_db_instance.db.id
}

output "identifier" {
  value = aws_db_instance.db.identifier
}

output "password" {
  value = aws_db_instance.db.password
}

output "resource_id" {
  value = aws_db_instance.db.resource_id
}

output "tags" {
  value = aws_db_instance.db.tags
}

resource "aws_db_subnet_group" "subnets" {
  name = "${var.meta.service}-db-mysql"

  subnet_ids = var.subnet_ids

  tags = {
    Name        = "${var.meta.service}-db-mysql"
    Service     = var.meta.service
    Environment = var.meta.env
  }
}

resource "aws_security_group" "sg" {
  name        = "${var.meta.service}-db-mysql"
  description = "access for nearsfeed crews"
  vpc_id      = var.vpc_id

  ingress {
    from_port = 0
    to_port   = 0
    protocol  = -1
    self      = true
  }

  ingress {
    security_groups = var.sg_ids
    from_port       = 0
    to_port         = 0
    protocol        = -1
  }


  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }
}
