# Grafana Foundation SDK – Go

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in Go.

> [!NOTE]
> This branch contains **types and builders generated for Grafana {{ .Extra.GrafanaVersion }}.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Installing

```shell
go get github.com/grafana/grafana-foundation-sdk/go@{{ .Extra.ReleaseBranch }}
```

## Example usage

[More examples](https://github.com/grafana/grafana-foundation-sdk/tree/main/examples/go) can be found at the repository root.

### Building a dashboard

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
                        LegendFormat({{ `"{{ device }}"` }}),
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

### Unmarshaling a dashboard

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func main() {
    // Required to correctly unmarshal panels and dataqueries
    plugins.RegisterDefaultPlugins()

    dashboardJSON, err := os.ReadFile("dashboard.json")
    if err != nil {
        panic(err)
    }

    sampleDashboard := &dashboard.Dashboard{}
    if err := json.Unmarshal(dashboardJSON, sampleDashboard); err != nil {
        panic(fmt.Sprintf("%s", err))
    }

    fmt.Printf("%#v\n", sampleDashboard)
}
```

### Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type and a builder for the custom query:

```go
package main

import (
    "encoding/json"

    "github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type CustomQuery struct {
    // RefId and Hide are expected on all queries
    RefId        *string `json:"refId,omitempty"`
    Hide         *bool   `json:"hide,omitempty"`

    // Query is specific to the CustomQuery type
    Query        string  `json:"query,omitempty"`
}

// Let cog know that CustomQuery is a Dataquery variant
func (resource CustomQuery) ImplementsDataqueryVariant() {}

func CustomQueryVariantConfig() variants.DataqueryConfig {
    return variants.DataqueryConfig{
        Identifier: "custom", // datasource plugin ID
        DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
            dataquery := &CustomQuery{}

            if err := json.Unmarshal(raw, dataquery); err != nil {
                return nil, err
            }

            return dataquery, nil
        },
    }
}

// Compile-time check to ensure that CustomQueryBuilder indeed is
// a builder for variants.Dataquery
var _ cog.Builder[variants.Dataquery] = (*CustomQueryBuilder)(nil)

type CustomQueryBuilder struct {
    internal *CustomQuery
}

func NewCustomQueryBuilder(query string) *CustomQueryBuilder {
    return &CustomQueryBuilder{
        internal: &CustomQuery{Query: query},
    }
}

func (builder *CustomQueryBuilder) Build() (variants.Dataquery, error) {
    return *builder.internal, nil
}

func (builder *CustomQueryBuilder) RefId(refId string) *CustomQueryBuilder {
    builder.internal.RefId = &refId
    return builder
}

func (builder *CustomQueryBuilder) Hide(hide bool) *CustomQueryBuilder {
    builder.internal.Hide = &hide
    return builder
}
```

Register the type with cog, and use it as usual to build a dashboard:

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/grafana/grafana-foundation-sdk/go/cog"
    "github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
    "github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func main() {
    // Required to correctly unmarshal panels and dataqueries
    plugins.RegisterDefaultPlugins()

    // This lets cog know about the newly created query type and how to unmarshal it.
    cog.NewRuntime().RegisterDataqueryVariant(CustomQueryVariantConfig())

    sampleDashboard, err := dashboard.NewDashboardBuilder("Custom query type").
        Uid("test-custom-query-type").
        Refresh("1m").
        Time("now-30m", "now").
        WithRow(dashboard.NewRowBuilder("Overview")).
        WithPanel(
            timeseries.NewPanelBuilder().
                Title("Sample panel").
                WithTarget(
                    NewCustomQueryBuilder("query here"),
                ),
        ).
        Build()
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

### Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type and a builder for the custom panel's options:

```go
package main

import (
    "encoding/json"

    "github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type CustomPanelOptions struct {
    MakeBeautiful bool `json:"makeBeautiful"`
}

func CustomPanelVariantConfig() variants.PanelcfgConfig {
    return variants.PanelcfgConfig{
        Identifier: "custom-panel", // plugin ID
        OptionsUnmarshaler: func(raw []byte) (any, error) {
            options := CustomPanelOptions{}

            if err := json.Unmarshal(raw, &options); err != nil {
                return nil, err
            }

            return options, nil
        },
    }
}

// Compile-time check to ensure that CustomPanelBuilder indeed is
// a builder for a dashboard.Panel.
var _ cog.Builder[dashboard.Panel] = (*CustomPanelBuilder)(nil)

type CustomPanelBuilder struct {
    internal *dashboard.Panel
    errors   map[string]cog.BuildErrors
}

func NewCustomPanelBuilder() *CustomPanelBuilder {
    return &CustomPanelBuilder{
        internal: &dashboard.Panel{
            Type: "custom-panel",
        },
        errors: make(map[string]cog.BuildErrors),
    }
}

func (builder *CustomPanelBuilder) Build() (dashboard.Panel, error) {
    var errs cog.BuildErrors

    for _, err := range builder.errors {
        errs = append(errs, cog.MakeBuildErrors("CustomPanel", err)...)
    }

    if len(errs) != 0 {
        return dashboard.Panel{}, errs
    }

    return *builder.internal, nil
}

func (builder *CustomPanelBuilder) Title(title string) *CustomPanelBuilder {
    builder.internal.Title = &title

    return builder
}

func (builder *CustomPanelBuilder) WithTarget(targets cog.Builder[variants.Dataquery]) *CustomPanelBuilder {
    targetsResource, err := targets.Build()
    if err != nil {
        builder.errors["targets"] = err.(cog.BuildErrors)
        return builder
    }
    builder.internal.Targets = append(builder.internal.Targets, targetsResource)

    return builder
}

// [other panel options omitted for brevity]

func (builder *CustomPanelBuilder) MakeBeautiful() *CustomPanelBuilder {
    if builder.internal.Options == nil {
        builder.internal.Options = &CustomPanelOptions{}
    }
    builder.internal.Options.(*CustomPanelOptions).MakeBeautiful = true

    return builder
}
```

Register the type with cog, and use it as usual to build a dashboard:

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/grafana/grafana-foundation-sdk/go/cog"
    "github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func main() {
    // Required to correctly unmarshal panels and dataqueries
    plugins.RegisterDefaultPlugins()

    // This lets cog know about the newly created panel type and how to unmarshal it.
    cog.NewRuntime().RegisterPanelcfgVariant(CustomPanelVariantConfig())

    sampleDashboard, err := dashboard.NewDashboardBuilder("Custom panel type").
        Uid("test-custom-panel").
        Refresh("1m").
        Time("now-30m", "now").
        WithRow(dashboard.NewRowBuilder("Overview")).
        WithPanel(
            NewCustomPanelBuilder().
                Title("Sample panel").
                MakeBeautiful(),
        ).
        Build()
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

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
