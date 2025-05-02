package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func memoryTotalStat() *stat.PanelBuilder {
	return defaultStat().
		Title("Memory total").
		Description("Amount of random-access memory (RAM) installed.\nIt represents the system's available working memory that applications and the operating system use to perform tasks.\nA higher memory total generally leads to better system performance and the ability to run more demanding applications and processes simultaneously.").
		WithTarget(
			prometheusQuery(`node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, ""),
		).
		Decimals(2).
		Unit(units.BytesIEC).
		ColorMode(common.BigValueColorModeNone)
}

func swapTotalStat() *stat.PanelBuilder {
	return defaultStat().
		Title("Total swap").
		Description("Total swap available.\n\nSwap is a space on a storage device (usually a dedicated swap partition or a swap file) \nused as virtual memory when the physical RAM (random-access memory) is fully utilized.\nSwap space helps prevent memory-related performance issues by temporarily transferring less-used data from RAM to disk,\nfreeing up physical memory for active processes and applications.").
		WithTarget(
			prometheusQuery(`node_memory_SwapTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, ""),
		).
		Unit(units.BytesIEC).
		ColorMode(common.BigValueColorModeNone)
}

func memoryUsageStat() *stat.PanelBuilder {
	query := `100 -
(
  avg by (instance) (node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) /
  avg by (instance) (node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"})
* 100
)`

	return defaultStat().
		Title("Memory usage").
		Description("RAM (random-access memory) currently in use by the operating system and running applications, in percent.").
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

func memoryUsageTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Memory usage").
		Description("- Used: The amount of physical memory currently in use by the system.\n- Cached: The amount of physical memory used for caching data from disk. The Linux kernel uses available memory to cache data that is read from or written to disk. This helps speed up disk access times.\n- Free: The amount of physical memory that is currently not in use.\n- Buffers: The amount of physical memory used for temporary storage of data being transferred between devices or applications.\n- Available: The amount of physical memory that is available for use by applications. This takes into account memory that is currently being used for caching but can be freed up if needed.").
		Span(18).
		WithTarget(prometheusQuery(`(
  node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
  node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
)`, "Memory used")).
		WithTarget(prometheusQuery(`node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "Memory cached")).
		WithTarget(prometheusQuery(`node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "Memory available")).
		WithTarget(prometheusQuery(`node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "Memory buffers")).
		WithTarget(prometheusQuery(`node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "Memory free")).
		WithTarget(prometheusQuery(`node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "Memory total")).
		Decimals(2).
		Unit(units.BytesIEC).
		GradientMode(common.GraphGradientModeOpacity).
		OverrideByRegexp(".*(T|t)otal.*", []dashboard.DynamicConfigValue{
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
