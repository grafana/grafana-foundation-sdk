package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
	"github.com/grafana/grafana-foundation-sdk/go/expr"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func timeseriesPanel(panelId int, query cog.Builder[dashboard.DataQueryKind]) *dashboard.PanelBuilder {
	return dashboard.NewPanelBuilder().
		Id(float64(panelId)).
		Visualization(timeseries.NewVisualizationBuilder()).
		Data(dashboard.NewQueryGroupBuilder().
			Target(
				dashboard.NewTargetBuilder().
					RefId(fmt.Sprintf("target-%d", panelId)).
					Query(query),
			),
		)
}

func statPanel(panelId int, query cog.Builder[dashboard.DataQueryKind]) *dashboard.PanelBuilder {
	return dashboard.NewPanelBuilder().
		Id(float64(panelId)).
		Visualization(stat.NewVisualizationBuilder()).
		Data(dashboard.NewQueryGroupBuilder().
			Target(
				dashboard.NewTargetBuilder().
					RefId(fmt.Sprintf("target-%d", panelId)).
					Query(query),
			),
		)
}

func main() {
	builder := dashboard.NewDashboardBuilder("[Example] Dashboard with expr queries").
		Editable(true).
		Panel("expr-sql", timeseriesPanel(
			1,
			expr.NewQueryBuilder().
				Sql(expr.NewTypeSqlBuilder().Expression("select * from A")),
		)).
		Panel("expr-math", statPanel(
			2,
			expr.NewQueryBuilder().
				Math(expr.NewTypeMathBuilder().Expression("2 + 4")),
		)).
		AutoGridLayout(
			dashboard.AutoGrid().
				Item(dashboard.AutoGridItem("expr-sql")).
				Item(dashboard.AutoGridItem("expr-math")),
		)

	manifestBuilder := dashboard.Manifest("example-dashboard-with-expr-queries", builder)
	manifest, err := manifestBuilder.Build()
	if err != nil {
		panic(err)
	}

	manifestJson, err := json.MarshalIndent(manifest, "", "  ")

	fmt.Println(string(manifestJson))
}
