variable "az" {
  type = string
}

variable "subnet_ipv4_cidr_block" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "nat_id" {
  type = string
}

variable "meta" {
  type = object({
    team = string
  })
}
