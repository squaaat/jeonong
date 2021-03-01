
variable "name" {
  type=string
}

variable "type" {
  type=string
}

variable "value" {
  type=string
}

variable "overwrite" {
  type=bool
}

# // ap-northeast-2
# provider "aws" {
#   alias = "apnortheast2"
#   region = "ap-northeast-2"
# }

resource "aws_ssm_parameter" "apnortheast2" {
  # provider = aws.apnortheast2

  name      = var.name
  type      = var.type
  value     = var.value
  overwrite = var.overwrite
}

// us-east-1
provider "aws" {
  alias = "useast1"
  region = "us-east-1"
}
resource "aws_ssm_parameter" "useast1" {
  provider = aws.useast1

  name      = var.name
  type      = var.type
  value     = var.value
  overwrite = var.overwrite
}
