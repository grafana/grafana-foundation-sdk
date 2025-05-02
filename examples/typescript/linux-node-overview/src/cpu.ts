import * as common from '@grafana/grafana-foundation-sdk/common';
import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as StatBuilder } from '@grafana/grafana-foundation-sdk/stat';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultStat, defaultTimeseries, prometheusQuery } from "./common";

export const countStat = (): StatBuilder => {
  return defaultStat()
    .title("CPU count")
    .description("CPU count is the number of processor cores or central processing units (CPUs) in a computer,\ndetermining its processing capability and ability to handle tasks concurrently.")
    .withTarget(
      prometheusQuery('count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})', 'Cores')
    )
    .decimals(0)
    .colorMode(common.BigValueColorMode.None);
};

export const usageStat = (): StatBuilder => {
  const query = '(((count by (instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))) \n- \navg by (instance) (sum by (instance, mode)(irate(node_cpu_seconds_total{mode=\'idle\',job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])))) * 100) \n/ \ncount by(instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))';

  return defaultStat()
    .title("CPU usage")
    .description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.")
    .withTarget(prometheusQuery(query, ''))
    .min(0)
    .max(100)
    .unit(units.Percent)
    .graphMode(common.BigValueGraphMode.Area)
    .thresholds(
      new dashboard.ThresholdsConfigBuilder()
        .mode(dashboard.ThresholdsMode.Absolute)
        .steps([
          { value: null, color: "green" },
          { value: 80, color: "red" },
        ])
    );
};

export const usageTimeseries = (): TimeseriesBuilder => {
  const query = '(\n  (1 - sum without (mode) (rate(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode=~"idle|iowait|steal"}[$__rate_interval])))\n/ ignoring(cpu) group_left\n  count without (cpu, mode) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})\n) * 100';

  return defaultTimeseries()
    .title("CPU usage")
    .description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.")
    .withTarget(prometheusQuery(query, 'CPU {{cpu}}'))
    .min(0)
    .max(100)
    .decimals(1)
    .unit(units.Percent)
    .thresholds(
      new dashboard.ThresholdsConfigBuilder()
        .mode(dashboard.ThresholdsMode.Absolute)
        .steps([
          { value: null, color: "green" },
          { value: 80, color: "red" },
        ])
    );
};

export const loadAverageTimeseries = (): TimeseriesBuilder => {
  return defaultTimeseries()
    .title("Load average")
    .description("System load average over the previous 1, 5, and 15 minute ranges.\n\nA measurement of how many processes are waiting for CPU cycles. The maximum number is the number of CPU cores for the node.")
    .withTarget(prometheusQuery(
      'node_load1{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      '1m',
    ))
    .withTarget(prometheusQuery(
      'node_load5{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      '5m',
    ))
    .withTarget(prometheusQuery(
      'node_load15{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      '15m',
    ))
    .withTarget(prometheusQuery(
      'count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})',
      'Cores',
    ))
    .min(0)
    .decimals(2)
    .unit(units.Short)
    .thresholds(
      new dashboard.ThresholdsConfigBuilder()
        .mode(dashboard.ThresholdsMode.Absolute)
        .steps([
          { value: null, color: "green" },
          { value: 80, color: "red" },
        ])
    )
    .overrideByRegexp("Cores", [
      {
        id: "color",
        value: {
          "fixedColor": "light-orange",
          "mode": "fixed",
        },
      },
      { id: "custom.fillOpacity", value: 0 },
      {
        id: "custom.lineStyle",
        value: {
          "dash": [10, 10],
          "fill": "dash",
        },
      },
    ]);
};
