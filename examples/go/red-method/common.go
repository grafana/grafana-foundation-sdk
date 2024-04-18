package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func grafanaDatasourceRef() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Uid:  cog.ToPtr("grafana"),
		Type: cog.ToPtr("grafana"),
	}
}

func defaultTimeseries() *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		LineWidth(1).
		FillOpacity(30).
		ShowPoints(common.VisibilityModeNever).
		DrawStyle(common.GraphDrawStyleLine).
		GradientMode(common.GraphGradientModeOpacity).
		SpanNulls(common.BoolOrFloat64{Bool: cog.ToPtr[bool](false)}).
		AxisBorderShow(false).
		LineInterpolation(common.LineInterpolationSmooth).
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
		)
}
