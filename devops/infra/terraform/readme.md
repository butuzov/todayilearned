# `terraform`

```shell
# Init & Init Update
terraform init
terraform init -upgrade

# Plan
terraform plan

# Apply
terraform apply -auto-approve

# Destroy
terraform destroy --auto-approve
```

## `.env`

Use `.auto.tfvars` to source variables.

```tfvars
FOOO = "bar"
```

## Debug

```shell
# On
export TF_LOG=trace
# Off
export TF_LOG=off
```


## Tooling

### `tenv`

Manages terraform installs (`tofu` (`opentofu`), `tf` (`terraform`), `tg` (`terragunt`), `atmos`).

```shell
# install tofu@1.6.0
tenv tofu install 1.6.0
tenv tf install "~> 1.6.0"
# reset
tenv tofu reset
# list
tenv tf list -v
```

### `terraform-ls`

Terraform Language Server to be used with [`hashicorp.terraform`](https://marketplace.visualstudio.com/items?itemName=hashicorp.terraform) and [`4ops.terraform`](https://marketplace.visualstudio.com/items?itemName=4ops.terraform) extensions.
