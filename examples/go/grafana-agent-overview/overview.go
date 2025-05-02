package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func runningInstancesTable() *table.PanelBuilder {
	return table.NewPanelBuilder().
		Title("Running Instances").
		Description("General statistics of running grafana agent instances.").
		Height(7).
		Span(24).
		Footer(common.NewTableFooterOptionsBuilder().CountRows(false).Reducer([]string{"sum"})).
		Datasource(datasourceRef("$prometheus_datasource")).
		WithTarget(
			tablePrometheusQuery(`count by (instance, version) (agent_build_info{job=~"$job", instance=~"$instance"})`, "A"),
		).
		WithTarget(
			tablePrometheusQuery(`max by (instance) (time() - process_start_time_seconds{job=~"$job", instance=~"$instance"})`, "B"),
		).
		// Transformations
		WithTransformation(dashboard.DataTransformerConfig{
			Id:      "merge",
			Options: map[string]any{},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "organize",
			Options: map[string]any{
				"excludeByName": map[string]any{
					"Time":     true,
					"Value #A": true,
				},
				"renameByName": map[string]any{
					"Value #B": "Uptime",
				},
			},
		}).
		// Overrides
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Seconds},
		})
}
