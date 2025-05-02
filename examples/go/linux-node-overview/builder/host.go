package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func uptimeStat() *stat.PanelBuilder {
	return defaultStat().
		Title("Uptime").
		Description("The duration of time that has passed since the last reboot or system start.").
		WithTarget(
			prometheusQuery(`time() - node_boot_time_seconds{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, ""),
		).
		Unit(units.DurationSeconds).
		Thresholds(dashboard.NewThresholdsConfigBuilder().
			Mode(dashboard.ThresholdsModeAbsolute).
			Steps([]dashboard.Threshold{
				{Value: nil, Color: "orange"},
				{Value: cog.ToPtr[float64](600), Color: "text"},
			}),
		)
}

func hostnameStat() *stat.PanelBuilder {
	return unameStat(
		"Hostname",
		"System's hostname.",
		"nodename",
	)
}

func kernelVersionStat() *stat.PanelBuilder {
	return unameStat(
		"Kernel version",
		"Kernel version of linux host.",
		"release",
	)
}

func osStat() *stat.PanelBuilder {
	return defaultStat().
		Title("OS").
		Description("Operating system.").
		WithTarget(
			tablePrometheusQuery(`node_os_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "A"),
		).
		ReduceOptions(common.NewReduceDataOptionsBuilder().Calcs([]string{"lastNotNull"}).Fields("/^pretty_name$/")).
		ColorMode(common.BigValueColorModeNone)
}
