resource "aws_route53_zone" "zone" {
  name = var.zone_name

  comment = "Managed by Terraform"
}

variable "zone_name" {
  type = string
}

output "zone_id" {
  value = aws_route53_zone.zone.id
}

output "zone_name" {
  value = aws_route53_zone.zone.name
}

