import * as common from '@grafana/grafana-foundation-sdk/common';
import * as table from "@grafana/grafana-foundation-sdk/table";
import { tablePrometheusQuery } from "./common";

export const runningInstancesTable = (): table.PanelBuilder => {
  return new table.PanelBuilder()
    .title("Running Instances")
    .description("General statistics of running grafana agent instances.")
    .height(7)
    .span(24)
    .footer(
      new common.TableFooterOptionsBuilder()
        .countRows(false)
        .reducer(["sum"])
    )
    .datasource({ uid: "$prometheus_datasource" })
    .withTarget(
      tablePrometheusQuery('count by (instance, version) (agent_build_info{job=~"$job", instance=~"$instance"})', "A"),
    )
    .withTarget(
      tablePrometheusQuery('max by (instance) (time() - process_start_time_seconds{job=~"$job", instance=~"$instance"})', "B"),
    )
    // Transformations
    .withTransformation({
      id: "merge",
      options: {}
    })
    .withTransformation({
      id: "organize",
      options: {
        excludeByName: {
          "Time": true,
          "Value #A": true,
        },
        "renameByName": {
          "Value #B": "Uptime",
        },
      },
    })
    // Overrides
    .overrideByName("Value #B", [
      { id: "unit", value: "s" },
    ]);
};
