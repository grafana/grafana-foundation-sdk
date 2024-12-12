# Defining a custom panel type

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

// Validate checks all the validation constraints that may be defined on `CustomPanelOptions` fields for violations and returns them.
func (resource CustomPanelOptions) Validate() error {
	return nil
}

func CustomPanelVariantConfig() variants.PanelcfgConfig {
    return variants.PanelcfgConfig{
        Identifier: "custom-panel", // plugin ID
        OptionsUnmarshaler: func(raw []byte) (any, error) {
            options := &CustomPanelOptions{}

            if err := json.Unmarshal(raw, options); err != nil {
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
    if err := builder.internal.Validate(); err != nil {
        return dashboard.Panel{}, err
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
