import * as common from '@grafana/grafana-foundation-sdk/common';
import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as StatBuilder } from '@grafana/grafana-foundation-sdk/stat';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultStat, defaultTimeseries, prometheusQuery } from "./common";

export const totalStat = (): StatBuilder => {
  return defaultStat()
    .title("Memory total")
    .description("Amount of random-access memory (RAM) installed.\nIt represents the system's available working memory that applications and the operating system use to perform tasks.\nA higher memory total generally leads to better system performance and the ability to run more demanding applications and processes simultaneously.")
    .withTarget(
      prometheusQuery('node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
    )
    .decimals(2)
    .unit(units.BytesIEC)
    .colorMode(common.BigValueColorMode.None);
};

export const swapTotalStat = (): StatBuilder => {
  return defaultStat()
    .title("Total swap")
    .description("Total swap available.\n\nSwap is a space on a storage device (usually a dedicated swap partition or a swap file) \nused as virtual memory when the physical RAM (random-access memory) is fully utilized.\nSwap space helps prevent memory-related performance issues by temporarily transferring less-used data from RAM to disk,\nfreeing up physical memory for active processes and applications.")
    .withTarget(
      prometheusQuery('node_memory_SwapTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
    )
    .unit(units.BytesIEC)
    .colorMode(common.BigValueColorMode.None);
};

export const usageStat = (): StatBuilder => {
  const query = '100 -\n(\n  avg by (instance) (node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) /\n  avg by (instance) (node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"})\n* 100\n)';

  return defaultStat()
    .title("Memory usage")
    .description("RAM (random-access memory) currently in use by the operating system and running applications, in percent.")
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
  return defaultTimeseries()
    .title("Memory usage")
    .description("- Used: The amount of physical memory currently in use by the system.\n- Cached: The amount of physical memory used for caching data from disk. The Linux kernel uses available memory to cache data that is read from or written to disk. This helps speed up disk access times.\n- Free: The amount of physical memory that is currently not in use.\n- Buffers: The amount of physical memory used for temporary storage of data being transferred between devices or applications.\n- Available: The amount of physical memory that is available for use by applications. This takes into account memory that is currently being used for caching but can be freed up if needed.")
    .withTarget(prometheusQuery(
      '(\n  node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}\n-\n  node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}\n-\n  node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}\n-\n  node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}\n)',
      'Memory used',
    ))
    .withTarget(prometheusQuery(
      'node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'Memory cached',
    ))
    .withTarget(prometheusQuery(
      'node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'Memory available',
    ))
    .withTarget(prometheusQuery(
      'node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'Memory buffers',
    ))
    .withTarget(prometheusQuery(
      'node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'Memory free',
    ))
    .withTarget(prometheusQuery(
      'node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'Memory total',
    ))
    .decimals(2)
    .unit(units.BytesIEC)
    .gradientMode(common.GraphGradientMode.Opacity)
    .overrideByRegexp(".*(T|t)otal.*", [
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
