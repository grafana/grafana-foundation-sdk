import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';
import * as testdata from '@grafana/grafana-foundation-sdk/testdata';
import * as timeseries from '@grafana/grafana-foundation-sdk/timeseries';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultTimeseries } from "./common";

export interface redConfig {
  dashboardTitle: string;
  serviceIds: string[];
}

export const red = (config: redConfig): dashboard.DashboardBuilder => {
  let builder = new dashboard.DashboardBuilder(`[Example] ${config.dashboardTitle}`)
    .uid("example-red-method")
    .tags(["generated", "red"])
    .editable()
    .tooltip(dashboard.DashboardCursorSync.Crosshair)
    .refresh("30s")
    .time({ from: "now-30m", to: "now" })
    .timezone("browser")
    .timepicker(
      new dashboard.TimePickerBuilder()
        .refreshIntervals(["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"])
    )
    // More info about the RED method
    .link(
      new dashboard.DashboardLinkBuilder("Grafana Agent Dashboards")
        .type(dashboard.DashboardLinkType.Link)
        .icon("question")
        .targetBlank(true)
        .url("https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method")
    );

  for (const serviceID of config.serviceIds) {
    builder = builder
      .withRow(new dashboard.RowBuilder(serviceID))
      .withPanel(requestRateTimeseries().span(8).height(8))
      .withPanel(errorRateTimeseries().span(8).height(8))
      .withPanel(durationTimeseries().span(8).height(8));
  }

  return builder;
};

export const requestRateTimeseries = (): timeseries.PanelBuilder => {
  return defaultTimeseries()
    .title("Request rate")
    .description("Number of requests handled by the service, per second.")
    .datasource({ uid: "grafana", type: "grafana" })
    .unit(units.RequestsPerSecond)
    .withTarget(
      new testdata.DataqueryBuilder().queryType("randomWalk")
    );
};

export const errorRateTimeseries = (): timeseries.PanelBuilder => {
  return defaultTimeseries()
    .title("Error rate")
    .description("Number of failed requests, per second.")
    .datasource({ uid: "grafana", type: "grafana" })
    .unit(units.RequestsPerSecond)
    .withTarget(
      new testdata.DataqueryBuilder().queryType("randomWalk")
    );
};

export const durationTimeseries = (): timeseries.PanelBuilder => {
  return defaultTimeseries()
    .title("Duration")
    .description("Time taken to process the requests.")
    .datasource({ uid: "grafana", type: "grafana" })
    .unit(units.Seconds)
    .withTarget(
      new testdata.DataqueryBuilder().queryType("randomWalk")
    );
};
