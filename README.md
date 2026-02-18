# Grafana Foundation SDK

The Grafana Foundation SDK is a set of **_types_**, and **_builder libraries_**
that let you define Grafana dashboards and other resources using strongly typed
code. By writing your resources as code, you can:

* **Leverage strong typing**: Catch errors at compile time, ensuring more reliable configurations.
* **Enhance version control**: Track changes seamlessly using standard version control systems like Git.
* **Automate deployments**: Integrate dashboard provisioning into your CI/CD pipelines for consistent and repeatable setups.

The SDK supports multiple programming languages, including **Go**, **TypeScript**, **Python**, **PHP**, and **Java**, so you can choose the one that best fits your development environment.

> [!NOTE]
> This SDK is best suited for Grafana >= 12, but will work with Grafana >= 10.

## Overview

Here's a quick overview of how the SDK works:

* **Builder pattern**: The SDK implements the builder pattern to let you define dashboards fluently. You start with a `DashboardBuilder`, then add panels, queries, and other components step by step.
* **Strong typing**: Everything in the SDK is strongly typed. This gives you autocompletion in your IDE, catches mistakes early, and helps ensure you're always using valid configuration values.
* **Structured options**: When a configuration gets complex (like data reduction or display settings), the SDK uses typed option builders to keep things readable and predictable.
* **Resources**: Multiple resources are supported by the SDK. Dashboards, alerts, …

For example, here is how a dashboard can be built in Go:

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/grafana/grafana-foundation-sdk/go/common"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
    "github.com/grafana/grafana-foundation-sdk/go/prometheus"
    "github.com/grafana/grafana-foundation-sdk/go/timeseries"
    "github.com/grafana/grafana-foundation-sdk/go/units"
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
                Unit(units.BitsPerSecondSI).
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

> [!NOTE]
> More examples can be found in the [`./examples/`](https://github.com/grafana/grafana-foundation-sdk/tree/main/examples) folder.

## Publishing resources

After you've defined your resource as code, call the `build()` function (its
actual name might be slightly different depending on language choice) and
output the result as a JSON.

With the JSON payload, you can:

* Call [Grafana's API](https://grafana.com/docs/grafana/latest/developer-resources/api-reference/http-api/) to programmatically manage the resource.
* Use [Grafana CLI](https://grafana.com/docs/grafana/latest/as-code/observability-as-code/grafana-cli/) to publish the resource from CLI.

## Next steps

With the basics of using the Grafana Foundation SDK in mind, here are some possible next steps:

* **Explore more features**: Check out the full [API reference](https://grafana.github.io/grafana-foundation-sdk/) to learn more about what the SDK can do.
* **Version control your resources**: Store your resources code in Git to track changes over time.
* **Automate provisioning with CI/CD**: [Integrate the SDK into your CI/CD pipeline](https://grafana.com/docs/grafana/latest/as-code/observability-as-code/foundation-sdk/dashboard-automation/) to deploy dashboards and other resources automatically.
* **Explore a more [real-world example of using the SDK](https://www.youtube.com/watch?v=ZjWdGVsrCiQ)**

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
