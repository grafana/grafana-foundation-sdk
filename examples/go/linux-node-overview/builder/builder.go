package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func Build() (dashboard.Dashboard, error) {
	builder := dashboard.NewDashboardBuilder("[Example] Linux node / overview").
		Uid("example-integration-linux-node").
		Tags([]string{"generated", "grafana-foundation-sdk", "linux-node-integration"}).
		Editable().
		Tooltip(dashboard.DashboardCursorSyncOff).
		Refresh("30s").
		Time("now-30m", "now").
		Timezone(common.TimeZoneBrowser).
		Timepicker(dashboard.NewTimePickerBuilder().
			RefreshIntervals([]string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}),
		).
		// "Back to fleet" link
		Link(dashboard.NewDashboardLinkBuilder("Back to Linux node / fleet").
			Type("link").
			Url("/d/node-fleet").
			KeepTime(true),
		).
		// Links to other linux-node-related dashboards
		Link(dashboard.NewDashboardLinkBuilder("All Linux node /  dashboards").
			Type("dashboards").
			Tags([]string{"linux-node-integration"}).
			IncludeVars(true).
			AsDropdown(true).
			KeepTime(true),
		).
		// "Data source" variable
		WithVariable(dashboard.NewDatasourceVariableBuilder("datasource").
			Label("Data source").
			Type("prometheus").
			Regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+").
			Multi(false),
		).
		// "Cluster" variable
		WithVariable(
			queryVariable("cluster", "Cluster", `label_values(node_uname_info{job=~"integrations/(node_exporter|unix)"}, cluster)`).
				AllValue(".*"),
		).
		// "Job" variable
		WithVariable(
			queryVariable("job", "Job", `label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster"}, job)`).
				AllValue(".+"),
		).
		// "Instance" variable
		WithVariable(
			queryVariable("instance", "Instance", `label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job"}, instance)`).
				AllValue(".+"),
		).
		// "Loki data source" variable
		WithVariable(dashboard.NewDatasourceVariableBuilder("loki_datasource").
			Label("Loki data source").
			Type("loki").
			Regex("(?!grafanacloud.+usage-insights|grafanacloud.+alert-state-history).+").
			Multi(false).
			Hide(dashboard.VariableHideHideVariable),
		).
		// Panels
		WithRow(dashboard.NewRowBuilder("Overview")).
		WithPanel(uptimeStat()).
		WithPanel(hostnameStat()).
		WithPanel(kernelVersionStat()).
		WithPanel(osStat()).
		WithPanel(cpuCountStat()).
		WithPanel(memoryTotalStat()).
		WithPanel(swapTotalStat()).
		WithPanel(rootMountSizeStat()).
		WithRow(dashboard.NewRowBuilder("CPU")).
		WithPanel(cpuUsageStat().Height(6)).
		WithPanel(cpuUsageTimeseries().Height(6)).
		WithPanel(loadAverageTimeseries().Height(6).Span(6)).
		WithRow(dashboard.NewRowBuilder("Memory")).
		WithPanel(memoryUsageStat().Height(6)).
		WithPanel(memoryUsageTimeseries().Height(6)).
		WithRow(dashboard.NewRowBuilder("Disk")).
		WithPanel(diskReadWriteTimeseries().Height(8)).
		WithPanel(diskSpaceUsageTable().Height(8)).
		WithRow(dashboard.NewRowBuilder("Network")).
		WithPanel(networkTrafficTimeseries().Height(8)).
		WithPanel(networkErrorsTimeseries().Height(8))

	return builder.Build()
}
