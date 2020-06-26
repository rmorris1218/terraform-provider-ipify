# terraform-provider-ipify

Terraform provider for [Ipify](https://ipify.org) to use your public IP in Terraform configurations.

## Install

```bash
git clone git@github.com:rmorris1218/terraform-provider-ipify
cd terraform-provider-ipify
mkdir -p ~/.terraform.d/plugins
go build -o ~/.terraform.d/plugins/terraform-provider-ipify
```

TODO: releases

## Using the Provider

Once the provider is installed on your system, you can use it in your Terraform configurations.

```hcl
provider "ipify" {}

data "ipify_current_ip" "runner_ip" {}

# print public IP to Terraform run outputs
output "current_ip" {
    value = data.ipify_current_ip.runner_ip.address
}

# use public IP CIDR notation for an AWS Security Group ingress rule
resource "aws_security_group" "jumphost" {
  # ...

  ingress {
    description = "ingress from workstation public ip"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [ data.ipify_current_ip.runner_ip.address_cidr ]
  }
  # ...
}
```

## Inputs

None.

## Outputs


| field | type | description |
| ----- | ---- | ----------- |
| address | string | public IP address from where terraform is running |
| address_cidr | string | the address field appended with single host CIDR notation |
