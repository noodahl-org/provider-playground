terraform {
  required_providers {
    provider-playground = {
      source = "hashicorp/provider-playground"
    }
  }
}

provider "provider-playground" {
  os = "darwin"
  pg_data = "/opt/homebrew/var/postgresql@17"
}

resource "provider-playground_postgres" "postgres"{
}

data "provider-playground_postgres" "postgres" {}

output "postgres_status" {
  value = data.provider-playground_postgres.postgres.status
}