variable "meta" {
  type = object({
    service = string
    env     = string
  })
}

variable "sg_ids" {
  type = list(string)
}

variable "subnet_ids" {
  type = list(string)
}

variable "db_password" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable db_meta {
  type = object({
    az                  = string
    engine              = string
    engine_version      = string
    volume_size         = number
    maximum_volume_size = number
    instance_class      = string
    dbname              = string
    username            = string
    publicly_accessible = bool
  })
}
