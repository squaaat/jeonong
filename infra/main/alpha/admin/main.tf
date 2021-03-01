variable "meta" {
  type = object({
    service = string
    env = string
    app = string
  })
}

// 참고1[policies]: https://www.serverless.com/plugins/serverless-nextjs-plugin
// 참고2[assume_role]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html
module "role" {
  source="../../../modules/simple_role"

  meta = var.meta
  assume_role_policy = <<JSON
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    },
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "edgelambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
JSON
  policy = <<JSON
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "ssm:Get*"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:CreateBucket",
        "s3:GetAccelerateConfiguration",
        "s3:GetObject",
        "s3:HeadBucket",
        "s3:ListBucket",
        "s3:PutAccelerateConfiguration",
        "s3:PutBucketPolicy",
        "s3:PutObject"
      ],
      "Resource": "arn:aws:s3:::nearsfeed-admin/*"
    }
  ]
}
JSON
}

output "role" {
  value = module.role
}

// global
module "ssm_admin_role_arn" {
  source = "../../../modules/global_ssm"

  name      = "/nearsfeed/infra/aws_iam_role/admin/arn"
  type      = "String"
  value     = module.role.arn
  overwrite = true
}
module "ssm_admin_role_id" {
  source = "../../../modules/global_ssm"

  name      = "/nearsfeed/infra/aws_iam_role/admin/id"
  type      = "String"
  value     = module.role.id
  overwrite = true
}
module "ssm_admin_role_policy_arn" {
  source = "../../../modules/global_ssm"

  name      = "/nearsfeed/infra/aws_iam_role/admin/policy_arn"
  type      = "String"
  value     = module.role.policy_arn
  overwrite = true
}
