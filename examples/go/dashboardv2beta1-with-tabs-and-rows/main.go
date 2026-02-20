package main

import (
	"encoding/json"
	"fmt"

	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
	"github.com/grafana/grafana-foundation-sdk/go/testdata"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func randomWalkPanel(panelId int, seriesCount int64) *dashboard.PanelBuilder {
	return dashboard.NewPanelBuilder().
		Title(fmt.Sprintf("New panel %d", panelId)).
		Visualization(timeseries.NewVisualizationBuilder()).
		Data(
			dashboard.NewQueryGroupBuilder().
				Target(
					dashboard.NewTargetBuilder().Query(
						testdata.NewQueryBuilder().
							ScenarioId("random_walk").
							SeriesCount(seriesCount),
					),
				),
		)
}

func main() {
	builder := dashboard.NewDashboardBuilder("[Example] Dashboard with tabs and rows").
		Description("Dashboard with tabs and rows generated with grafana-foundation-sdk").
		// Available panels
		Panel("panel-1", randomWalkPanel(1, 4)).
		Panel("panel-2", randomWalkPanel(2, 5)).
		Panel("panel-3", randomWalkPanel(3, 1)).
		Panel("panel-4", randomWalkPanel(4, 1)).
		// Layout building
		TabsLayout(
			dashboard.Tabs().
				Tab(dashboard.Tab("Tab without rows").GridLayout(
					dashboard.Grid().
						Item(dashboard.GridItem("panel-1").Width(12).Height(8)),
				)).
				Tab(dashboard.Tab("Tab With Rows").RowsLayout(
					dashboard.Rows().
						Row(
							dashboard.Row("Row without tabs").
								Collapse(false).
								AutoGridLayout(
									dashboard.AutoGrid().Item(dashboard.AutoGridItem("panel-2")),
								),
						).
						Row(
							dashboard.Row("Row with tabs").
								Collapse(true).
								TabsLayout(
									dashboard.Tabs().
										Tab(
											dashboard.Tab("First tab").
												AutoGridLayout(
													dashboard.AutoGrid().Item(dashboard.AutoGridItem("panel-3")),
												),
										).
										Tab(
											dashboard.Tab("Second tab").
												AutoGridLayout(
													dashboard.AutoGrid().Item(dashboard.AutoGridItem("panel-4")),
												),
										),
								),
						).
						Row(
							dashboard.Row("Empty row").
								AutoGridLayout(dashboard.AutoGrid()),
						).
						Row(
							dashboard.Row("Hide header row").
								HideHeader(true).
								AutoGridLayout(dashboard.AutoGrid()),
						),
				)),
		)

	manifestBuilder := dashboard.Manifest("example-dashboard-with-tabs-and-rows", builder)
	manifest, err := manifestBuilder.Build()
	if err != nil {
		panic(err)
	}

	manifestJson, err := json.MarshalIndent(manifest, "", "  ")

	fmt.Println(string(manifestJson))
}
