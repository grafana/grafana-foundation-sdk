package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/resource"
)

func main() {
	redDashboard, err := red(redConfig{
		DashboardTitle: "RED method",
		ServiceIDs:     []string{"sample-service", "payments", "front-gateway"},
	}).Build()
	if err != nil {
		panic(err)
	}
	manifest := resource.Manifest{
		ApiVersion: "dashboard.grafana.app/v1beta1",
		Kind:       resource.DashboardKind,
		Metadata: resource.Metadata{
			Name: *redDashboard.Uid,
		},
		Spec: redDashboard,
	}

	manifestJson, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(manifestJson))
}
