terraform {
  required_version = "0.11.5"

  backend "s3" {
    bucket  = "envy-terraform-state"
    key     = "envy.tfstate"
    encrypt = true

    dynamodb_table = "envy-terraform-lock"

    profile = "envy/manage"
    region  = "eu-west-1"
  }
}

provider "aws" {
  version = "1.13.0"

  profile = "envy/manage"
  region  = "eu-west-1"
}
