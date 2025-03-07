# provider-playground

trimmed scaffolding for custom terraform providers

## setup

### prerequisites

This guide assumes the user has the following dependencies installed: 

- [golang](https://go.dev/doc/install)
  - A golang environment variable of `$GOBIN` set 
- [terraoform-cli](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

### mac os (only supported)

Since we're not uploading this project to the terraform registry have to have dev override specified. 
In your root (~/) directory, create a `.terraformrc ` file with the following contents

```
provider_installation {
  dev_overrides {
    "hashicorp/provider-playground" = "/Users/$USERNAME/.terraform.d/plugins/registry.terraform.io/hashicorp/provider-playground/dev/darwin_arm64"
  }
  direct {}
}
```
_Replace the `$USERNAME` environment variable with the active user_

## debugging

### vscode

create a `launch.json` file at `/.vscode/launch.json`:
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go"
        }
    ]
}
```

When debugging `main.go` an environment variable used by terraform is output:

```
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/provider-playground":{"Protocol":"grpc","ProtocolVersion":6,"Pid":30800,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/bn/_g99llgx12gfcq9dmt96t2c00000gn/T/plugin3760811943"}}}'
```

Export the environment variable before attempting to run a `terraform plan`
`export TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/provider-playground":{"Protocol":"grpc","ProtocolVersion":6,"Pid":30800,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/bn/_g99llgx12gfcq9dmt96t2c00000gn/T/plugin3760811943"}}}'`

Navigate to the `examples` folder and run `terraform plan`. If everything is configured correctly you should see:
```
...
Terraform used the selected providers to generate the following execution plan.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # provider-playground_postgres.postgres will be created
  + resource "provider-playground_postgres" "postgres" {
      + id     = (known after apply)
      + status = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.
```

### output 

```
terraform apply

Terraform used the selected providers to generate the following execution plan.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # provider-playground_postgres.postgres will be created
  + resource "provider-playground_postgres" "postgres" {
      + id     = (known after apply)
      + status = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + postgres_status = ""

...
Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
Outputs:

postgres_status = <<EOT
started
```

terraform destroy
```
Terraform will perform the following actions:

  # provider-playground_postgres.postgres will be destroyed
  - resource "provider-playground_postgres" "postgres" {
      - id     = "/REq+9laxsaO2Ar3tTTPzpo0G1iUel4EvEh7vKPeM+k=" -> null
      - status = <<-EOT
            started
        EOT -> null
    }

Plan: 0 to add, 0 to change, 1 to destroy.

Changes to Outputs:
  - postgres_status = <<-EOT
        started
    EOT -> null
...
Destroy complete! Resources: 1 destroyed.
```
