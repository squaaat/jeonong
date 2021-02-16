resource "aws_s3_bucket" "infrastructure_bucket" {
  bucket = "nearsfeed-infrastructure"
  acl    = "private"
}

resource "aws_dynamodb_table" "terraform-lock" {
  name = "nearsfeed-terraform-lock"

  read_capacity  = 1
  write_capacity = 1

  hash_key = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
