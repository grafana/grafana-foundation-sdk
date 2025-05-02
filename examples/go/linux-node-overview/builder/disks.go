package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func rootMountSizeStat() *stat.PanelBuilder {
	return defaultStat().
		Title("Root mount size").
		Description("Total capacity on the primary mount point /.").
		WithTarget(
			prometheusQuery(`node_filesystem_size_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mountpoint="/", fstype!="rootfs"}`, ""),
		).
		Unit(units.BytesIEC).
		ColorMode(common.BigValueColorModeNone)
}

func diskReadWriteTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Disk reads/writes").
		Description("Disk read/writes in bytes per second.").
		WithTarget(prometheusQuery(`irate(node_disk_read_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])`, "{{ device }} read")).
		WithTarget(prometheusQuery(`irate(node_disk_written_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])`, "{{ device }} written")).
		WithTarget(prometheusQuery(`irate(node_disk_io_time_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])`, "{{ device }} io util")).
		Decimals(0).
		Unit(units.BytesPerSecondSI).
		FillOpacity(1).
		GradientMode(common.GraphGradientModeOpacity).
		OverrideByRegexp("/time|used|busy|util/", []dashboard.DynamicConfigValue{
			{Id: "custom.axisSoftMax", Value: 100},
			{Id: "custom.drawStyle", Value: "points"},
			{Id: "unit", Value: units.Percent},
		})
}

func diskSpaceUsageTable() *table.PanelBuilder {
	return table.NewPanelBuilder().
		Title("Disk space usage").
		Description("Disk utilisation in percent, by mountpoint. Some duplication can occur if the same filesystem is mounted in multiple locations.").
		Footer(common.NewTableFooterOptionsBuilder().CountRows(false).Reducer([]string{"sum"})).
		Datasource(datasourceRef("$datasource")).
		WithTarget(
			tablePrometheusQuery(`node_filesystem_size_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "TOTAL").
				Instant(),
		).
		WithTarget(
			tablePrometheusQuery(`node_filesystem_avail_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}`, "FREE").
				LegendFormat("{{ mountpoint }} free").
				Instant(),
		).
		Unit(units.BytesIEC).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().
				Mode(dashboard.ThresholdsModeAbsolute).
				Steps([]dashboard.Threshold{
					{Value: nil, Color: "light-blue"},
					{Value: cog.ToPtr(0.8), Color: "light-yellow"},
					{Value: cog.ToPtr(0.9), Color: "light-red"},
				}),
		).
		// Transformations
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "groupBy",
			Options: map[string]any{
				"fields": map[string]any{
					"Value #FREE": map[string]any{
						"aggregations": []string{"lastNotNull"},
						"operation":    "aggregate",
					},
					"Value #TOTAL": map[string]any{
						"aggregations": []string{"lastNotNull"},
						"operation":    "aggregate",
					},
					"mountpoint": map[string]any{
						"aggregations": []string{},
						"operation":    "groupby",
					},
				},
			},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id:      "merge",
			Options: map[string]any{},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "calculateField",
			Options: map[string]any{
				"alias": "Used",
				"binary": map[string]any{
					"left":     "Value #TOTAL (lastNotNull)",
					"operator": "-",
					"reducer":  "sum",
					"right":    "Value #FREE (lastNotNull)",
				},
				"mode": "binary",
				"reduce": map[string]any{
					"reducer": "sum",
				},
			},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "calculateField",
			Options: map[string]any{
				"alias": "Used, %",
				"binary": map[string]any{
					"left":     "Used",
					"operator": "/",
					"reducer":  "sum",
					"right":    "Value #TOTAL (lastNotNull)",
				},
				"mode": "binary",
				"reduce": map[string]any{
					"reducer": "sum",
				},
			},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "organize",
			Options: map[string]any{
				"excludeByName": map[string]any{},
				"indexByName": map[string]any{
					"Used":                       3,
					"Used, %":                    4,
					"Value #FREE (lastNotNull)":  2,
					"Value #TOTAL (lastNotNull)": 1,
					"mountpoint":                 0,
				},
				"renameByName": map[string]any{
					"Value #FREE (lastNotNull)":  "Available",
					"Value #TOTAL (lastNotNull)": "Size",
					"mountpoint":                 "Mounted on",
				},
			},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "sortBy",
			Options: map[string]any{
				"fields": map[string]any{},
				"sort": []map[string]any{
					{"field": "Mounted on", "desc": false},
				},
			},
		}).
		// Overrides
		OverrideByName("Mounted on", []dashboard.DynamicConfigValue{
			{Id: "custom.width", Value: "260"},
		}).
		OverrideByName("Size", []dashboard.DynamicConfigValue{
			{Id: "custom.width", Value: "80"},
		}).
		OverrideByName("Used", []dashboard.DynamicConfigValue{
			{Id: "custom.width", Value: "80"},
		}).
		OverrideByName("Available", []dashboard.DynamicConfigValue{
			{Id: "custom.width", Value: "80"},
		}).
		OverrideByName("Used, %", []dashboard.DynamicConfigValue{
			{
				Id: "custom.cellOptions",
				Value: map[string]any{
					"mode": "basic",
					"type": "gauge",
				},
			},
			{Id: "min", Value: 0},
			{Id: "max", Value: 1},
			{Id: "unit", Value: units.PercentUnit},
		})
}
