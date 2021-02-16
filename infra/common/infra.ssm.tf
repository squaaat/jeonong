resource "aws_ssm_parameter" "aws_route_53_nearsfeed_com_zone_id" {
  name      = "/nearsfeed/infra/aws_route_53/nearsfeed.com/zone_id"
  type      = "String"
  value     = module.route53_zone.zone_id
  overwrite = true
}


// ------
resource "aws_ssm_parameter" "aws_acm_certificate_nearsfeed_com_domain" {
  name      = "/nearsfeed/infra/aws_acm_certificate/nearsfeed.com/domain"
  type      = "String"
  value     = data.aws_acm_certificate.nearsfeed_com.domain
  overwrite = true
}

data "aws_acm_certificate" "nearsfeed_com" {
  domain   = "nearsfeed.com"
  statuses = ["ISSUED"]
}

output "infra_ssm" {
  value = {
    aws_route_53_nearsfeed_com_zone_id       = aws_ssm_parameter.aws_route_53_nearsfeed_com_zone_id.value
    aws_acm_certificate_nearsfeed_com_domain = aws_ssm_parameter.aws_acm_certificate_nearsfeed_com_domain.value
  }
}
