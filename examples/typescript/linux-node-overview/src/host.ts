import * as common from '@grafana/grafana-foundation-sdk/common';
import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as StatBuilder } from '@grafana/grafana-foundation-sdk/stat';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultStat, prometheusQuery, tablePrometheusQuery, unameStat } from "./common";

export const uptimeStat = (): StatBuilder => {
  return defaultStat()
    .title("Uptime")
    .description("The duration of time that has passed since the last reboot or system start.")
    .withTarget(
      prometheusQuery('time() - node_boot_time_seconds{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
    )
    .unit(units.DurationSeconds)
    .thresholds(
      new dashboard.ThresholdsConfigBuilder()
        .mode(dashboard.ThresholdsMode.Absolute)
        .steps([
          { value: null, color: "orange" },
          { value: 600, color: "text" },
        ])
    );
};

export const hostnameStat = (): StatBuilder => {
  return unameStat(
    "Hostname",
    "System's hostname.",
    "nodename",
  );
};

export const kernelVersionStat = (): StatBuilder => {
  return unameStat(
    "Kernel version",
    "Kernel version of linux host.",
    "release",
  );
};

export const osStat = (): StatBuilder => {
  return defaultStat()
    .title("OS")
    .description("Operating system.")
    .withTarget(
      tablePrometheusQuery('node_os_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'A')
    )
    .reduceOptions(
      new common.ReduceDataOptionsBuilder()
        .calcs(["lastNotNull"])
        .fields("/^pretty_name$/")
    )
    .colorMode(common.BigValueColorMode.None);
};
