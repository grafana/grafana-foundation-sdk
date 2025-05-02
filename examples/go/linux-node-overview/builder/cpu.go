package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func cpuCountStat() *stat.PanelBuilder {
	return defaultStat().
		Title("CPU count").
		Description("CPU count is the number of processor cores or central processing units (CPUs) in a computer,\ndetermining its processing capability and ability to handle tasks concurrently.").
		WithTarget(
			prometheusQuery(`count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})`, "Cores"),
		).
		Decimals(0).
		ColorMode(common.BigValueColorModeNone)
}

func cpuUsageStat() *stat.PanelBuilder {
	query := `(((count by (instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))) 
- 
avg by (instance) (sum by (instance, mode)(irate(node_cpu_seconds_total{mode='idle',job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])))) * 100) 
/ 
count by(instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))`

	return defaultStat().
		Title("CPU usage").
		Description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.").
		WithTarget(prometheusQuery(query, "")).
		Min(0).
		Max(100).
		Unit(units.Percent).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().
				Mode(dashboard.ThresholdsModeAbsolute).
				Steps([]dashboard.Threshold{
					{Value: nil, Color: "green"},
					{Value: cog.ToPtr[float64](80), Color: "red"},
				}),
		).
		GraphMode(common.BigValueGraphModeArea)
}

func cpuUsageTimeseries() *timeseries.PanelBuilder {
	query := `(
  (1 - sum without (mode) (rate(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode=~"idle|iowait|steal"}[$__rate_interval])))
/ ignoring(cpu) group_left
  count without (cpu, mode) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})
) * 100`

	return defaultTimeseries().
		Title("CPU usage").
		Description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.").
		WithTarget(prometheusQuery(query, "CPU {{cpu}}")).
		Min(0).
		Max(100).
		Decimals(1).
		Unit(units.Percent).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().
				Mode(dashboard.ThresholdsModeAbsolute).
				Steps([]dashboard.Threshold{
					{Value: nil, Color: "green"},
					{Value: cog.ToPtr[float64](80), Color: "red"},
				}),
		)
}

func loadAverageTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Load average").
		Description("System load average over the previous 1, 5, and 15 minute ranges.\n\nA measurement of how many processes are waiting for CPU cycles. The maximum number is the number of CPU cores for the node.").
		WithTarget(
			prometheusQuery(`node_load1{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "1m"),
		).
		WithTarget(
			prometheusQuery(`node_load5{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "5m"),
		).
		WithTarget(
			prometheusQuery(`node_load15{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "15m"),
		).
		WithTarget(
			prometheusQuery(`count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})`, "Cores"),
		).
		Min(0).
		Decimals(2).
		Unit(units.Short).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().
				Mode(dashboard.ThresholdsModeAbsolute).
				Steps([]dashboard.Threshold{
					{Value: nil, Color: "green"},
					{Value: cog.ToPtr[float64](80), Color: "red"},
				}),
		).
		OverrideByRegexp("Cores", []dashboard.DynamicConfigValue{
			{
				Id: "color",
				Value: map[string]any{
					"fixedColor": "light-orange",
					"mode":       "fixed",
				},
			},
			{Id: "custom.fillOpacity", Value: 0},
			{
				Id: "custom.lineStyle",
				Value: map[string]any{
					"dash": []int{10, 10},
					"fill": "dash",
				},
			},
		})
}
