provider "aws" {
}

# Aliased Providers (for doing things different region/account)
provider "aws" {
  alias = "czi-si"
}

terraform {
  required_version = "=1.2.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.45.0"
    }
  }
}

module "test" {
  source = "../"

  subdomain = "test"
  env       = "test"
  tags      = {}

  providers = {
    aws.czi-si = aws.czi-si
  }
}
