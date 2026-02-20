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
			dashboard.NewTabsLayoutBuilder().
				Tab(dashboard.NewTabsLayoutTabBuilder("Tab without rows").GridLayout(
					dashboard.NewGridLayoutBuilder().
						Item(dashboard.NewGridLayoutItemBuilder("panel-1").Width(12).Height(8)),
				)).
				Tab(dashboard.NewTabsLayoutTabBuilder("Tab With Rows").RowsLayout(
					dashboard.NewRowsLayoutBuilder().
						Row(
							dashboard.NewRowsLayoutRowBuilder("Row without tabs").
								Collapse(false).
								AutoGridLayout(
									dashboard.NewAutoGridLayoutBuilder().Item(dashboard.NewAutoGridLayoutItemBuilder("panel-2")),
								),
						).
						Row(
							dashboard.NewRowsLayoutRowBuilder("Row with tabs").
								Collapse(true).
								TabsLayout(
									dashboard.NewTabsLayoutBuilder().
										Tab(
											dashboard.NewTabsLayoutTabBuilder("First tab").
												AutoGridLayout(
													dashboard.NewAutoGridLayoutBuilder().Item(dashboard.NewAutoGridLayoutItemBuilder("panel-3")),
												),
										).
										Tab(
											dashboard.NewTabsLayoutTabBuilder("Second tab").
												AutoGridLayout(
													dashboard.NewAutoGridLayoutBuilder().Item(dashboard.NewAutoGridLayoutItemBuilder("panel-4")),
												),
										),
								),
						).
						Row(
							dashboard.NewRowsLayoutRowBuilder("Empty row").
								AutoGridLayout(dashboard.NewAutoGridLayoutBuilder()),
						).
						Row(
							dashboard.NewRowsLayoutRowBuilder("Hide header row").
								HideHeader(true).
								AutoGridLayout(dashboard.NewAutoGridLayoutBuilder()),
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
