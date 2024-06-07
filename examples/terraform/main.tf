terraform {
  required_providers {
    grafana = {
      source  = "grafana/grafana"
      version = "3.1.0"
    }
  }
}

provider "grafana" {
  url  = "http://localhost:3000"
  auth = "admin:admin"
}
