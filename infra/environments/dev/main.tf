locals {
  product_name = "kns23_catch_up"
  env = "dev"
}

# tfstateをS3で管理する
terraform {
  backend "s3" {
    region  = "ap-northeast-1"
    encrypt = false

    # 以下は実行時引数から指定
    # bucket  = "XXXXXXXXXXXXXXX"
    # key     = "XXXXXXXXXXXXXXX"
  }
}

terraform {
  # Terraformのバージョンを指定
  required_version = ">=1.7"

  # プロバイダーを指定
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      env = "${local.env}"
      service = "${local.product_name}"
    }
  }
}

module "api" {
  source = "../../modules/api"

  product_name = "${local.product_name}"
  env = "${local.env}"
}

module "database" {
  source = "../../modules/database"
}