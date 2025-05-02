import {
  DashboardBuilder,
  DashboardCursorSync,
  DashboardLinkBuilder,
  DashboardLinkType,
  DatasourceVariableBuilder, RowBuilder,
  TimePickerBuilder,
  VariableHide,
} from '@grafana/grafana-foundation-sdk/dashboard';
import { queryVariable } from "./common";
import * as cpu from "./cpu";
import * as disk from "./disk";
import * as host from "./host";
import * as memory from "./memory";
import * as network from "./network";

export const builder = new DashboardBuilder('[Example] Linux node / overview')
  .uid('example-integration-linux-node')
  .tags(["generated", "grafana-foundation-sdk", "linux-node-integration"])
  .editable()
  .tooltip(DashboardCursorSync.Off)
  .refresh("30s")
  .time({ from: "now-30m", to: "now" })
  .timezone("browser")
  .timepicker(
    new TimePickerBuilder()
      .refreshIntervals(["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"])
  )
  // "Back to fleet" link
  .link(
    new DashboardLinkBuilder("Grafana Agent Dashboards")
      .type(DashboardLinkType.Link)
      .url("/d/node-fleet")
      .keepTime(true)
  )
  // Links to other linux-node-related dashboards
  .link(
    new DashboardLinkBuilder("All Linux node /  dashboards")
      .type(DashboardLinkType.Dashboards)
      .tags(["linux-node-integration"])
      .includeVars(true)
      .asDropdown(true)
      .keepTime(true)
  )
  // "Data source" variable
  .withVariable(
    new DatasourceVariableBuilder("datasource")
      .label("Data source")
      .type("prometheus")
      .regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+")
      .multi(false)
  )
  // "Loki data source" variable
  .withVariable(
    new DatasourceVariableBuilder("loki_datasource")
      .label("Loki data source")
      .type("loki")
      .regex("(?!grafanacloud.+usage-insights|grafanacloud.+alert-state-history).+")
      .multi(false)
      .hide(VariableHide.HideVariable)
  )
  // "Cluster" variable
  .withVariable(
    queryVariable("cluster", "Cluster", 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)"}, cluster)')
      .allValue('.*')
  )
  // "Job" variable
  .withVariable(
    queryVariable("job", "Job", 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster"}, job)')
      .allValue('.+')
  )
  // "Instance" variable
  .withVariable(
    queryVariable("instance", "Instance", 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job"}, instance)')
      .allValue('.+')
  )
  // Panels
  .withRow(new RowBuilder("Overview"))
  .withPanel(host.uptimeStat())
  .withPanel(host.hostnameStat())
  .withPanel(host.kernelVersionStat())
  .withPanel(host.osStat())
  .withPanel(cpu.countStat())
  .withPanel(memory.totalStat())
  .withPanel(memory.swapTotalStat())
  .withPanel(disk.rootMountSizeStat())
  .withRow(new RowBuilder("CPU"))
  .withPanel(cpu.usageStat().height(6))
  .withPanel(cpu.usageTimeseries().height(6))
  .withPanel(cpu.loadAverageTimeseries().height(6).span(6))
  .withRow(new RowBuilder("Memory"))
  .withPanel(memory.usageStat().height(6))
  .withPanel(memory.usageTimeseries().height(6).span(18))
  .withRow(new RowBuilder("Disk"))
  .withPanel(disk.readWriteTimeseries().height(8))
  .withPanel(disk.usageTable().height(8))
  .withRow(new RowBuilder("Network"))
  .withPanel(network.trafficTimeseries().height(8))
  .withPanel(network.errorsTimeseries().height(8))
  ;

console.log(JSON.stringify(builder.build(), null, 2));
