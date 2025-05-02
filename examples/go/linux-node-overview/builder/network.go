package builder

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func networkTrafficTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Network traffic").
		Description("Network traffic (bits per sec) measures data transmitted and received.").
		AxisLabel("out(-) | in(+)").
		AxisPlacement(common.AxisPlacementAuto).
		WithTarget(prometheusQuery(`irate(node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
# only show interfaces that had traffic change at least once during selected dashboard interval:
and
increase(
    node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
    ) > 0`, "{{ device }} received")).
		WithTarget(prometheusQuery(`irate(node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
# only show interfaces that had traffic change at least once during selected dashboard interval:
and
increase(
    node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
    ) > 0`, "{{ device }} transmitted")).
		Decimals(1).
		Unit(units.BitsPerSecondSI).
		GradientMode(common.GraphGradientModeOpacity).
		OverrideByRegexp("/transmit|tx|out/", []dashboard.DynamicConfigValue{
			{Id: "custom.transform", Value: "negative-Y"},
		})
}

func networkErrorsTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Network errors and dropped packets").
		Description(`**Network errors**:

Network errors refer to issues that occur during the transmission of data across a network. 

These errors can result from various factors, including physical issues, jitter, collisions, noise and interference.

Monitoring network errors is essential for diagnosing and resolving issues, as they can indicate problems with network hardware or environmental factors affecting network quality.

**Dropped packets**:

Dropped packets occur when data packets traveling through a network are intentionally discarded or lost due to congestion, resource limitations, or network configuration issues. 

Common causes include network congestion, buffer overflows, QoS settings, and network errors, as corrupted or incomplete packets may be discarded by receiving devices.

Dropped packets can impact network performance and lead to issues such as degraded voice or video quality in real-time applications.
`).
		AxisLabel("out(-) | in(+)").
		AxisPlacement(common.AxisPlacementAuto).
		WithTarget(prometheusQuery(`irate(node_network_transmit_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} errors transmitted")).
		WithTarget(prometheusQuery(`irate(node_network_receive_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} errors received")).
		WithTarget(prometheusQuery(`irate(node_network_transmit_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} transmitted dropped")).
		WithTarget(prometheusQuery(`irate(node_network_receive_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} received dropped")).
		Decimals(1).
		Unit(units.PacketsPerSecond).
		GradientMode(common.GraphGradientModeOpacity).
		OverrideByRegexp("/transmit|tx|out/", []dashboard.DynamicConfigValue{
			{Id: "custom.transform", Value: "negative-Y"},
		})
}
