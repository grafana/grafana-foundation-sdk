# Grafana Foundation SDK – Go

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in Go.

> ℹ️ This branch contains **types and builders generated for Grafana v10.4.x.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Maturity

> _The code in this repository should be considered experimental. Documentation is only
available alongside the code. It comes with no support, but we are keen to receive
feedback on the product and suggestions on how to improve it, though we cannot commit
to resolution of any particular issue. No SLAs are available. It is not meant to be used
in production environments, and the risks are unknown/high._

Grafana Labs defines experimental features as follows:

> Projects and features in the Experimental stage are supported only by the Engineering
teams; on-call support is not available. Documentation is either limited or not provided
outside of code comments. No SLA is provided.
>
> Experimental projects or features are primarily intended for open source engineers who
want to participate in ensuring systems stability, and to gain consensus and approval
for open source governance projects.
>
> Projects and features in the Experimental phase are not meant to be used in production
environments, and the risks are unknown/high.

## Installing

```shell
go get github.com/grafana/grafana-foundation-sdk/go@v10.4.x+cog-v0.0.x
```

## Example usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func main() {
	builder := dashboard.NewDashboardBuilder("Sample dashboard").
		Uid("generated-from-go").
		Tags([]string{"generated", "from", "go"}).
		Refresh("1m").
		Time("now-30m", "now").
		Timezone(common.TimeZoneBrowser).
		WithRow(dashboard.NewRowBuilder("Overview")).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Network Received").
				Unit("bps").
				Min(0).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr(`rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8`).
						LegendFormat("{{ device }}"),
				),
		)

	sampleDashboard, err := builder.Build()
	if err != nil {
		panic(err)
	}
	dashboardJson, err := json.MarshalIndent(sampleDashboard, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dashboardJson))
}
```

## License

[Apache 2.0 License](./LICENSE)
