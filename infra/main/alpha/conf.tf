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
    key            = "terraform/main/alpha"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "nearsfeed-terraform-lock"
  }
}

provider "aws" {
  region = "ap-northeast-2"
}
