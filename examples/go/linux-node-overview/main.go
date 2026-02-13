package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/examples/go/linux-node-overview/builder"
	"github.com/grafana/grafana-foundation-sdk/go/resource"
)

func main() {
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
