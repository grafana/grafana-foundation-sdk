package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/expr"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func timeseriesPanel(query cog.Builder[variants.Dataquery]) *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		WithTarget(query)
}

func statPanel(query cog.Builder[variants.Dataquery]) *stat.PanelBuilder {
	return stat.NewPanelBuilder().
		WithTarget(query)
}

func main() {
	builder := dashboard.NewDashboardBuilder("[Example] Dashboard with expr queries").
		Editable().
		Uid("example-dashboard-with-expr-queries").
		WithPanel(timeseriesPanel(
			expr.NewExprBuilder().
				TypeSql(expr.NewTypeSqlBuilder().Expression("select * from A")),
		)).
		WithPanel(statPanel(
			expr.NewExprBuilder().
				TypeMath(expr.NewTypeMathBuilder().Expression("2 + 4")),
		))

	sampleDashboard, err := builder.Build()
	if err != nil {
		panic(err)
	}

	dashboardJson, err := json.MarshalIndent(sampleDashboard, "", "  ")

	fmt.Println(string(dashboardJson))
}
