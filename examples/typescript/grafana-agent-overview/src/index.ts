import {
  DashboardBuilder,
  DashboardCursorSync,
  DashboardLinkBuilder,
  DashboardLinkType,
  DatasourceVariableBuilder,
  QueryVariableBuilder, RowBuilder,
  TimePickerBuilder,
  VariableRefresh,
  VariableSort
} from '@grafana/grafana-foundation-sdk/dashboard';
import { runningInstancesTable } from "./overview";
import { targetsTimeseries, targetSyncTimeseries } from "./discovery";
import { appendedSamplesTimeseries, avgScrapeIntervalDurationTimeseries, scrapeFailuresTimeseries } from "./retrieval";

const builder = new DashboardBuilder('[Example] Grafana Agent Overview')
  .uid('example-integration-agent')
  .tags(["generated", "grafana-foundation-sdk", "grafana-agent-integration"])
  .editable()
  .tooltip(DashboardCursorSync.Off)
  .refresh("30s")
  .time({ from: "now-30m", to: "now" })
  .timezone("browser")
  .timepicker(
    new TimePickerBuilder()
      .refreshIntervals(["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"])
  )
  // Links to other agent-related dashboards
  .link(
    new DashboardLinkBuilder("Grafana Agent Dashboards")
      .type(DashboardLinkType.Dashboards)
      .icon("external link")
      .tags(["grafana-agent-integration"])
      .includeVars(true)
      .keepTime(true)
  )
  // "Data source" variable
  .withVariable(
    new DatasourceVariableBuilder("prometheus_datasource")
      .label("Data source")
      .type("prometheus")
      .regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+")
      .multi(false)
  )
  // "Job" variable
  .withVariable(
    new QueryVariableBuilder("job")
      .label("Job")
      .query("label_values(agent_build_info, job)")
      .datasource({ uid: "$prometheus_datasource" })
      .current({
        selected: true,
        text: "All",
        value: "$__all",
      })
      .refresh(VariableRefresh.OnTimeRangeChanged)
      .sort(VariableSort.AlphabeticalAsc)
      .multi(true)
      .includeAll(true)
  )
  // "Instance" variable
  .withVariable(
    new QueryVariableBuilder("instance")
      .label("Instance")
      .query("label_values(agent_build_info{job=~\"$job\"}, instance)")
      .datasource({ uid: "$prometheus_datasource" })
      .current({
        selected: true,
        text: "All",
        value: "$__all",
      })
      .refresh(VariableRefresh.OnTimeRangeChanged)
      .sort(VariableSort.AlphabeticalAsc)
      .multi(true)
      .includeAll(true)
  )
  // Panels
  .withRow(new RowBuilder("Overview"))
  .withPanel(runningInstancesTable())
  .withRow(new RowBuilder("Prometheus Discovery"))
  .withPanel(targetSyncTimeseries())
  .withPanel(targetsTimeseries())
  .withRow(new RowBuilder("Prometheus Retrieval"))
  .withPanel(avgScrapeIntervalDurationTimeseries())
  .withPanel(scrapeFailuresTimeseries())
  .withPanel(appendedSamplesTimeseries());

console.log(JSON.stringify(builder.build(), null, 2));
