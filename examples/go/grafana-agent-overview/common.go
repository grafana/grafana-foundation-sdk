package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func datasourceRef(uid string) dashboard.DataSourceRef {
	return dashboard.DataSourceRef{Uid: &uid}
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
		Instant().
		IntervalFactor(2).
		RefId(refID)
}

func defaultTimeseries() *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		Height(7).
		Span(12).
		LineWidth(1).
		FillOpacity(0).
		PointSize(5).
		ShowPoints(common.VisibilityModeAuto).
		DrawStyle(common.GraphDrawStyleLine).
		GradientMode(common.GraphGradientModeNone).
		Legend(common.NewVizLegendOptionsBuilder().
			DisplayMode(common.LegendDisplayModeList).
			Placement(common.LegendPlacementBottom).
			ShowLegend(true),
		).
		Tooltip(common.NewVizTooltipOptionsBuilder().
			Mode(common.TooltipDisplayModeSingle).
			Sort(common.SortOrderNone),
		).
		ThresholdsStyle(
			common.NewGraphThresholdsStyleConfigBuilder().
				Mode(common.GraphThresholdsStyleModeOff),
		).
		SpanNulls(common.BoolOrFloat64{Bool: cog.ToPtr(false)}).
		AxisBorderShow(false)
}
