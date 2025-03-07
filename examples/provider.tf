terraform {
  required_providers {
    provider-playground = {
      source = "hashicorp/provider-playground"
    }
  }
}

provider "provider-playground" {
  os = "darwin"
  
}

resource "provider-playground_postgres" "postgres"{
  version = "17"
}

data "provider-playground_postgres" "postgres" {}

output "postgres_status" {
  value = data.provider-playground_postgres.postgres.status
}