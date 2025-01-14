locals {
  standard_secrets = {
    kind               = "k8s"
    vpc_id             = var.cloud-env.vpc_id
    zone_id            = var.base_zone_id
    external_zone_name = local.base_domain
    internal_zone_name = local.env_domain

    cloud_env             = var.cloud-env
    eks_cluster           = var.eks-cluster
    certificate_arn       = module.cert.arn
    oauth_certificate_arn = length(var.oauth_dns_prefix) == 0 ? module.cert.arn : module.cert_oauth.arn
    tags                  = var.tags

    proxy_service_name = module.proxy.proxy_service_name

    ecrs = { for name, ecr in module.ecrs : name => { "url" : ecr.repository_url } }
    dbs = {
      for name, db in module.dbs :
      name => {
        "database_user" : db.master_username,
        "database_password" : db.master_password,
        "database_host" : db.endpoint,
        "database_name" : db.database_name,
        "database_port" : db.port,
      }
    }
  }

  # TODO: this only works if all additional_secrets values are maps!
  merged_secrets = { for key, value in var.additional_secrets : key => merge(lookup(local.standard_secrets, key, {}), value) }
  secret_string = merge(
    local.standard_secrets,
    local.merged_secrets,
  )
}

resource "kubernetes_secret" "happy_env_secret" {
  metadata {
    name      = "integration-secret"
    namespace = kubernetes_namespace.happy.id
  }
  data = {
    "integration_secret" = jsonencode(local.secret_string)
  }
}
