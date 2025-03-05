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

