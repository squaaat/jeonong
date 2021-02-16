resource "aws_route53_record" "record" {
  zone_id = var.zone_id
  name    = var.domain
  type    = var.type
  ttl     = var.ttl
  records = var.records
}

variable "zone_id" {
  type = string
}

variable "domain" {
  type = string
}

variable "type" {
  type = string
}

variable "ttl" {
  type = number
}

variable "records" {
  type = list(string)
}

output "record" {
  value = aws_route53_record.record
}