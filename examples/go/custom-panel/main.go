package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/resource"
)

func main() {
	// Required to correctly unmarshal panels and dataqueries
	plugins.RegisterDefaultPlugins()

	// This lets cog know about the newly created query type and how to unmarshal it.
	cog.NewRuntime().RegisterPanelcfgVariant(CustomPanelVariantConfig())

	builder := dashboard.NewDashboardBuilder("[Example] Custom panel").
		Uid("example-custom-panel").
		WithPanel(
			NewCustomPanelBuilder().
				Title("Sample custom panel").
				MakeBeautiful(),
		)

	sampleDashboard, err := builder.Build()
	if err != nil {
		panic(err)
	}

	manifest := resource.Manifest{
		ApiVersion: "dashboard.grafana.app/v1beta1",
		Kind:       resource.DashboardKind,
		Metadata: resource.Metadata{
			Name: *sampleDashboard.Uid,
		},
		Spec: sampleDashboard,
	}

	manifestJson, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(manifestJson))
}
