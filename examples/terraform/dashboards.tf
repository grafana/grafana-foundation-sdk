locals {
  go_builders = toset([
    "${path.module}/../go/custom-panel",
    "${path.module}/../go/custom-query",
    "${path.module}/../go/grafana-agent-overview",
    "${path.module}/../go/linux-node-overview",
    "${path.module}/../go/red-method",
  ])
}

// List and run the go dashboard builders. Any of the languages would also work.
// We are using jq to convert the output to a map of strings, as expected by the external data source.
// The dashboard building could be done with a single Go program that outputs the expected TF format,
// but this example does it in multiple Go calls in order to reuse the existing examples.
data "external" "dashboards" {
  for_each    = local.go_builders
  working_dir = each.value
  program = ["bash", "-c", <<EOF
  go run *.go | jq '{"dashboard": . | tostring}'
EOF
  ]
}

resource "grafana_dashboard" "examples" {
  for_each    = local.go_builders
  config_json = data.external.dashboards[each.key].result.dashboard
}
