package builder

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func datasourceRef(uid string) dashboard.DataSourceRef {
	return dashboard.DataSourceRef{Uid: &uid}
}

func queryVariable(name string, label string, query string) *dashboard.QueryVariableBuilder {
	return dashboard.NewQueryVariableBuilder(name).
		Label(label).
		Query(dashboard.StringOrMap{String: cog.ToPtr(query)}).
		Datasource(datasourceRef("$datasource")).
		Current(dashboard.VariableOption{
			Selected: cog.ToPtr(true),
			Text:     dashboard.StringOrArrayOfString{ArrayOfString: []string{"All"}},
			Value:    dashboard.StringOrArrayOfString{ArrayOfString: []string{"$__all"}},
		}).
		Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
		Sort(dashboard.VariableSortAlphabeticalAsc).
		Multi(true).
		IncludeAll(true)
}

func prometheusQuery(query string, legend string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().
		Expr(query).
		LegendFormat(legend)
}

func tablePrometheusQuery(query string, refID string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().
		Expr(query).
		Format(prometheus.PromQueryFormatTable).
		RefId(refID)
}

func defaultStat() *stat.PanelBuilder {
	return stat.NewPanelBuilder().
		Datasource(datasourceRef("$datasource")).
		Height(2).
		Span(6).
		Decimals(1).
		ReduceOptions(common.NewReduceDataOptionsBuilder().Calcs([]string{"lastNotNull"})).
		ColorMode(common.BigValueColorModeValue).
		GraphMode(common.BigValueGraphModeNone)
}

func unameStat(panelTitle string, panelDescription string, field string) *stat.PanelBuilder {
	return defaultStat().
		Title(panelTitle).
		Description(panelDescription).
		WithTarget(
			tablePrometheusQuery(`node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "A"),
		).
		ReduceOptions(common.NewReduceDataOptionsBuilder().Calcs([]string{"lastNotNull"}).Fields(fmt.Sprintf("/^%s$/", field))).
		ColorMode(common.BigValueColorModeNone)
}

func defaultTimeseries() *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		Datasource(datasourceRef("$datasource")).
		LineWidth(2).
		FillOpacity(30).
		LineInterpolation(common.LineInterpolationSmooth).
		ShowPoints(common.VisibilityModeNever).
		DrawStyle(common.GraphDrawStyleLine).
		GradientMode(common.GraphGradientModeScheme).
		Legend(common.NewVizLegendOptionsBuilder().
			DisplayMode(common.LegendDisplayModeList).
			Placement(common.LegendPlacementBottom).
			ShowLegend(true),
		).
		Tooltip(common.NewVizTooltipOptionsBuilder().
			Mode(common.TooltipDisplayModeMulti).
			Sort(common.SortOrderDescending),
		).
		ThresholdsStyle(
			common.NewGraphThresholdsStyleConfigBuilder().
				Mode(common.GraphThresholdsStyleModeOff),
		).
		SpanNulls(common.BoolOrFloat64{Bool: cog.ToPtr(false)}).
		AxisBorderShow(false)
}
