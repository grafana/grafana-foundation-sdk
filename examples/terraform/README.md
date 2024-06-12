# Examples

This folder contains examples of the Foundation SDK being used in combination with [Terraform](https://www.terraform.io/) and the [Grafana Provider](https://registry.terraform.io/providers/grafana/grafana/latest/docs). 

While the Foundation SDK does not integrate directly with TF/the HCL language, it can be used in Terraform "external datasources". Building dashboards is a complex problem that is hard to model/reason with in the HCL language.

## Running this example

```console
$ terraform init
$ terraform apply
```
