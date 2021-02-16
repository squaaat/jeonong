terraform {
  required_version = "0.14.3"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.22.0"
    }
  }

  backend "s3" {
    bucket         = "nearsfeed-infrastructure"
    key            = "terraform/common"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "nearsfeed-terraform-lock"
  }
}

provider "aws" {
  region = "ap-northeast-2"
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
