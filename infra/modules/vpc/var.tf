variable "meta" {
  type = object({
    team = string
  })
}

variable "cidr_block" {
  type        = string
  description = "vpc's maximum range for ipv4"
}

data "aws_availability_zone" "a" {
  name = "ap-northeast-2a"
}

data "aws_availability_zone" "b" {
  name = "ap-northeast-2b"
}

data "aws_availability_zone" "c" {
  name = "ap-northeast-2c"
}

data "aws_availability_zone" "d" {
  name = "ap-northeast-2d"
}

