locals {
  meta = {
      service = "nearsfeed"
      env     = "alpha"
      app     = "admin"
  }
}

module "admin" {
  source = "./admin"

  meta = local.meta
}

output "admin" {
  value = module.admin
}