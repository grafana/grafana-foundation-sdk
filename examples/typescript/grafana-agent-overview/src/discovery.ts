import * as timeseries from "@grafana/grafana-foundation-sdk/timeseries";
import * as units from "@grafana/grafana-foundation-sdk/units";
import { defaultTimeseries, prometheusQuery } from "./common";

export const targetSyncTimeseries = (): timeseries.PanelBuilder => {
  return defaultTimeseries()
    .title("Target Sync")
    .description("Actual interval to sync the scrape pool.")
    .datasource({ uid: "$prometheus_datasource" })
    .unit(units.Seconds)
    .withTarget(prometheusQuery(
      'sum(rate(prometheus_target_sync_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval])) by (instance, scrape_job)',
      "{{instance}}/{{scrape_job}}"
    ));
};

export const targetsTimeseries = (): timeseries.PanelBuilder => {
  return defaultTimeseries()
    .title("Targets")
    .description("Discovered targets by prometheus service discovery.")
    .datasource({ uid: "$prometheus_datasource" })
    .unit(units.Short)
    .withTarget(prometheusQuery(
      'sum by (instance) (prometheus_sd_discovered_targets{job=~"$job", instance=~"$instance"})',
      "{{instance}}"
    ));
};
