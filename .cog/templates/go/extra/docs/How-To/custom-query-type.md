# Defining a custom query type

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
    RefId string `json:"refId"`
    Hide  *bool  `json:"hide,omitempty"`

    // Query is specific to the CustomQuery type
    Query string `json:"query,omitempty"`
}

func (resource CustomQuery) Equals(otherCandidate variants.Dataquery) bool {
    if otherCandidate == nil {
        return false
    }

    other, ok := otherCandidate.(CustomQuery)
    if !ok {
        return false
    }

    if resource.RefId != other.RefId {
        return false
    }

    if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
        return false
    }
    if resource.Hide != nil && *resource.Hide != *other.Hide {
        return false
    }

    return resource.Query == other.Query
}

func (resource CustomQuery) Validate() error {
    return nil
}

// Let cog know that CustomQuery is a Dataquery variant
func (resource CustomQuery) ImplementsDataqueryVariant() {}

func (resource CustomQuery) DataqueryType() string {
    return "custom"
}

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
    if err := builder.internal.Validate(); err != nil {
        return CustomQuery{}, err
    }

    return *builder.internal, nil
}

func (builder *CustomQueryBuilder) RefId(refId string) *CustomQueryBuilder {
    builder.internal.RefId = refId
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
