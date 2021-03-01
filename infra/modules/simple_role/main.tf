variable "meta" {
  type = object({
    service = string
    app = string
    env     = string
  })
}

variable "assume_role_policy" {
  type = string
}

variable "policy" {
  type = string
}

output "id" {
  value = aws_iam_role.role.id
}

output "arn" {
  value = aws_iam_role.role.arn
}

output "policy_arn" {
  value = aws_iam_policy.policy.arn
}

resource "aws_iam_role" "role" {
  name               = "role-${var.meta.service}-${var.meta.app}-${var.meta.env}"
  assume_role_policy = var.assume_role_policy
}

resource "aws_iam_policy" "policy" {
  name               = "policy-${var.meta.service}-${var.meta.app}-${var.meta.env}"
  description        = "policy-${var.meta.service}-${var.meta.app}-${var.meta.env}"

  policy = var.policy
}

resource "aws_iam_role_policy_attachment" "task_execution_role_policy" {
  role       = aws_iam_role.role.id
  policy_arn = aws_iam_policy.policy.arn
}


