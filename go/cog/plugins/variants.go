// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package plugins

import (
	annotationslist "github.com/grafana/grafana-foundation-sdk/go/annotationslist"
	azuremonitor "github.com/grafana/grafana-foundation-sdk/go/azuremonitor"
	barchart "github.com/grafana/grafana-foundation-sdk/go/barchart"
	bargauge "github.com/grafana/grafana-foundation-sdk/go/bargauge"
	candlestick "github.com/grafana/grafana-foundation-sdk/go/candlestick"
	canvas "github.com/grafana/grafana-foundation-sdk/go/canvas"
	cloudwatch "github.com/grafana/grafana-foundation-sdk/go/cloudwatch"
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardlist "github.com/grafana/grafana-foundation-sdk/go/dashboardlist"
	datagrid "github.com/grafana/grafana-foundation-sdk/go/datagrid"
	debug "github.com/grafana/grafana-foundation-sdk/go/debug"
	elasticsearch "github.com/grafana/grafana-foundation-sdk/go/elasticsearch"
	gauge "github.com/grafana/grafana-foundation-sdk/go/gauge"
	geomap "github.com/grafana/grafana-foundation-sdk/go/geomap"
	googlecloudmonitoring "github.com/grafana/grafana-foundation-sdk/go/googlecloudmonitoring"
	grafanapyroscope "github.com/grafana/grafana-foundation-sdk/go/grafanapyroscope"
	heatmap "github.com/grafana/grafana-foundation-sdk/go/heatmap"
	histogram "github.com/grafana/grafana-foundation-sdk/go/histogram"
	logs "github.com/grafana/grafana-foundation-sdk/go/logs"
	loki "github.com/grafana/grafana-foundation-sdk/go/loki"
	news "github.com/grafana/grafana-foundation-sdk/go/news"
	nodegraph "github.com/grafana/grafana-foundation-sdk/go/nodegraph"
	parca "github.com/grafana/grafana-foundation-sdk/go/parca"
	piechart "github.com/grafana/grafana-foundation-sdk/go/piechart"
	prometheus "github.com/grafana/grafana-foundation-sdk/go/prometheus"
	stat "github.com/grafana/grafana-foundation-sdk/go/stat"
	statetimeline "github.com/grafana/grafana-foundation-sdk/go/statetimeline"
	statushistory "github.com/grafana/grafana-foundation-sdk/go/statushistory"
	table "github.com/grafana/grafana-foundation-sdk/go/table"
	tempo "github.com/grafana/grafana-foundation-sdk/go/tempo"
	text "github.com/grafana/grafana-foundation-sdk/go/text"
	timeseries "github.com/grafana/grafana-foundation-sdk/go/timeseries"
	trend "github.com/grafana/grafana-foundation-sdk/go/trend"
	xychart "github.com/grafana/grafana-foundation-sdk/go/xychart"
)

func RegisterDefaultPlugins() {
	runtime := cog.NewRuntime()

	// Panelcfg variants
	runtime.RegisterPanelcfgVariant(annotationslist.VariantConfig())
	runtime.RegisterPanelcfgVariant(barchart.VariantConfig())
	runtime.RegisterPanelcfgVariant(bargauge.VariantConfig())
	runtime.RegisterPanelcfgVariant(candlestick.VariantConfig())
	runtime.RegisterPanelcfgVariant(canvas.VariantConfig())
	runtime.RegisterPanelcfgVariant(dashboardlist.VariantConfig())
	runtime.RegisterPanelcfgVariant(datagrid.VariantConfig())
	runtime.RegisterPanelcfgVariant(debug.VariantConfig())
	runtime.RegisterPanelcfgVariant(gauge.VariantConfig())
	runtime.RegisterPanelcfgVariant(geomap.VariantConfig())
	runtime.RegisterPanelcfgVariant(heatmap.VariantConfig())
	runtime.RegisterPanelcfgVariant(histogram.VariantConfig())
	runtime.RegisterPanelcfgVariant(logs.VariantConfig())
	runtime.RegisterPanelcfgVariant(news.VariantConfig())
	runtime.RegisterPanelcfgVariant(nodegraph.VariantConfig())
	runtime.RegisterPanelcfgVariant(piechart.VariantConfig())
	runtime.RegisterPanelcfgVariant(stat.VariantConfig())
	runtime.RegisterPanelcfgVariant(statetimeline.VariantConfig())
	runtime.RegisterPanelcfgVariant(statushistory.VariantConfig())
	runtime.RegisterPanelcfgVariant(table.VariantConfig())
	runtime.RegisterPanelcfgVariant(text.VariantConfig())
	runtime.RegisterPanelcfgVariant(timeseries.VariantConfig())
	runtime.RegisterPanelcfgVariant(trend.VariantConfig())
	runtime.RegisterPanelcfgVariant(xychart.VariantConfig())

	// Dataquery variants
	runtime.RegisterDataqueryVariant(azuremonitor.VariantConfig())
	runtime.RegisterDataqueryVariant(cloudwatch.VariantConfig())
	runtime.RegisterDataqueryVariant(elasticsearch.VariantConfig())
	runtime.RegisterDataqueryVariant(googlecloudmonitoring.VariantConfig())
	runtime.RegisterDataqueryVariant(grafanapyroscope.VariantConfig())
	runtime.RegisterDataqueryVariant(loki.VariantConfig())
	runtime.RegisterDataqueryVariant(parca.VariantConfig())
	runtime.RegisterDataqueryVariant(prometheus.VariantConfig())
	runtime.RegisterDataqueryVariant(tempo.VariantConfig())
}
