package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/testdata"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

type redConfig struct {
	DashboardTitle string
	ServiceIDs     []string
}

func red(config redConfig) *dashboard.DashboardBuilder {
	builder := dashboard.NewDashboardBuilder(fmt.Sprintf("[Example] %s", config.DashboardTitle)).
		Uid("example-red-method").
		Tags([]string{"generated", "red"}).
		Editable().
		Tooltip(dashboard.DashboardCursorSyncCrosshair).
		Refresh("30s").
		Time("now-30m", "now").
		Timezone(common.TimeZoneBrowser).
		Timepicker(dashboard.NewTimePickerBuilder().
			RefreshIntervals([]string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}),
		).
		// More info about the RED method
		Link(dashboard.NewDashboardLinkBuilder("RED method").
			Type("link").
			Icon("question").
			TargetBlank(true).
			Url("https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method"),
		)

	for _, serviceID := range config.ServiceIDs {
		builder = builder.WithRow(dashboard.NewRowBuilder(serviceID)).
			WithPanel(requestRateTimeseries().Span(8).Height(8)).
			WithPanel(errorRateTimeseries().Span(8).Height(8)).
			WithPanel(durationTimeseries().Span(8).Height(8))
	}

	return builder
}

func requestRateTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Request rate").
		Description("Number of requests handled by the service, per second.").
		Unit(units.RequestsPerSecond).
		WithTarget(
			testdata.NewDataqueryBuilder().
				QueryType("randomWalk").
				Datasource(grafanaDatasourceRef()),
		)
}

func errorRateTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Error rate").
		Description("Number of failed requests, per second.").
		Unit(units.RequestsPerSecond).
		WithTarget(
			testdata.NewDataqueryBuilder().
				QueryType("randomWalk").
				Datasource(grafanaDatasourceRef()),
		)
}

func durationTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Duration").
		Description("Time taken to process the requests.").
		Datasource(grafanaDatasourceRef()).
		Unit(units.Seconds).
		WithTarget(
			testdata.NewDataqueryBuilder().
				QueryType("randomWalk").
				Datasource(grafanaDatasourceRef()),
		)
}
