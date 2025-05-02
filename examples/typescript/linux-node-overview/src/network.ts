import * as common from '@grafana/grafana-foundation-sdk/common';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultTimeseries, prometheusQuery } from "./common";

export const trafficTimeseries = (): TimeseriesBuilder => {
  return defaultTimeseries()
    .title("Network traffic")
    .description("Network traffic (bits per sec) measures data transmitted and received.")
    .axisLabel("out(-) | in(+)")
    .axisPlacement(common.AxisPlacement.Auto)
    .withTarget(prometheusQuery(
      'irate(node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8\nand\nincrease(\n    node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]\n    ) > 0',
      '{{ device }} received',
    ))
    .withTarget(prometheusQuery(
      'irate(node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8\nand\nincrease(\n    node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]\n    ) > 0',
      '{{ device }} transmitted',
    ))
    .decimals(1)
    .unit(units.BitsPerSecondSI)
    .gradientMode(common.GraphGradientMode.Opacity)
    .overrideByRegexp("/transmit|tx|out/", [
      { id: "custom.transform", value: "negative-Y" },
    ]);
};

export const errorsTimeseries = (): TimeseriesBuilder => {
  return defaultTimeseries()
    .title("Network errors and dropped packets")
    .description("**Network errors**:\n\nNetwork errors refer to issues that occur during the transmission of data across a network. \n\nThese errors can result from various factors, including physical issues, jitter, collisions, noise and interference.\n\nMonitoring network errors is essential for diagnosing and resolving issues, as they can indicate problems with network hardware or environmental factors affecting network quality.\n\n**Dropped packets**:\n\nDropped packets occur when data packets traveling through a network are intentionally discarded or lost due to congestion, resource limitations, or network configuration issues. \n\nCommon causes include network congestion, buffer overflows, QoS settings, and network errors, as corrupted or incomplete packets may be discarded by receiving devices.\n\nDropped packets can impact network performance and lead to issues such as degraded voice or video quality in real-time applications.\n")
    .axisLabel("out(-) | in(+)")
    .axisPlacement(common.AxisPlacement.Auto)
    .withTarget(prometheusQuery(`irate(node_network_transmit_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} errors transmitted"))
    .withTarget(prometheusQuery(`irate(node_network_receive_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} errors received"))
    .withTarget(prometheusQuery(`irate(node_network_transmit_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} transmitted dropped"))
    .withTarget(prometheusQuery(`irate(node_network_receive_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0`, "{{ device }} received dropped"))
    .decimals(1)
    .unit(units.PacketsPerSecond)
    .gradientMode(common.GraphGradientMode.Opacity)
    .overrideByRegexp("/transmit|tx|out/", [
      { id: "custom.transform", value: "negative-Y" },
    ]);
};
