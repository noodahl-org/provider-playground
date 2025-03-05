# provider-playground

trimmed scaffolding for custom terraform providers

## setup

### prerequisites

This guide assumes the user has the following dependencies installed: 
- [golang](https://go.dev/doc/install)
  - A golang environment variable of `$GOBIN` set 
- [terraoform-cli](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
- 

### mac os

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
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp/provider-playground in /Users/tachi/.terraform.d/plugins/registry.terraform.io/hashicorp/provider-playground/dev/darwin_arm64
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause
│ the state to become incompatible with published releases.
╵

No changes. Your infrastructure matches the configuration.

Terraform has compared your real infrastructure against your configuration and found no differences, so no
changes are needed.
```