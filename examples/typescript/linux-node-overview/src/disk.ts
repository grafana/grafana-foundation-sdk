import * as common from '@grafana/grafana-foundation-sdk/common';
import { PanelBuilder as StatBuilder } from '@grafana/grafana-foundation-sdk/stat';
import { PanelBuilder as TableBuilder } from '@grafana/grafana-foundation-sdk/table';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';
import * as units from '@grafana/grafana-foundation-sdk/units';
import { defaultStat, defaultTimeseries, prometheusQuery, tablePrometheusQuery } from "./common";

export const rootMountSizeStat = (): StatBuilder => {
  return defaultStat()
    .title("Root mount size")
    .description("Total capacity on the primary mount point /.")
    .withTarget(
      prometheusQuery('node_filesystem_size_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mountpoint="/", fstype!="rootfs"}', '')
    )
    .unit(units.BytesIEC)
    .colorMode(common.BigValueColorMode.None);
};

export const readWriteTimeseries = (): TimeseriesBuilder => {
  return defaultTimeseries()
    .title("Disk reads/writes")
    .description("Disk read/writes in bytes per second.")
    .withTarget(prometheusQuery(
      'irate(node_disk_read_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
      '{{ device }} read',
    ))
    .withTarget(prometheusQuery(
      'irate(node_disk_written_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
      '{{ device }} written',
    ))
    .withTarget(prometheusQuery(
      'irate(node_disk_io_time_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])',
      '{{ device }} io util',
    ))
    .decimals(0)
    .unit(units.BytesPerSecondSI)
    .fillOpacity(1)
    .gradientMode(common.GraphGradientMode.Opacity)
    .overrideByRegexp("/time|used|busy|util/", [
      { id: "custom.axisSoftMax", value: 100 },
      { id: "custom.drawStyle", value: "points" },
      { id: "unit", value: "percent" },
    ]);
};

export const usageTable = (): TableBuilder => {
  return new TableBuilder()
    .title("Disk space usage")
    .description("Disk utilisation in percent, by mountpoint. Some duplication can occur if the same filesystem is mounted in multiple locations.")
    .footer(
      new common.TableFooterOptionsBuilder()
        .countRows(false)
        .reducer(["sum"])
    )
    .withTarget(tablePrometheusQuery(
      'node_filesystem_size_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'TOTAL',
    ).instant())
    .withTarget(tablePrometheusQuery(
      'node_filesystem_avail_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
      'FREE',
    ).instant().legendFormat("{{ mountpoint }} free"))
    .unit(units.BytesIEC)
    // Transformations
    .withTransformation({
      id: "groupBy",
      options: {
        fields: {
          "Value #FREE": {
            aggregations: ["lastNotNull"],
            operation: "aggregate"
          },
          "Value #TOTAL": {
            aggregations: ["lastNotNull"],
            operation: "aggregate"
          },
          mountpoint: {
            aggregations: [],
            operation: "groupby"
          }
        }
      }
    })
    .withTransformation({
      id: "merge",
      options: {}
    })
    .withTransformation({
      id: "calculateField",
      options: {
        alias: "Used",
        binary: {
          left: "Value #TOTAL (lastNotNull)",
          operator: "-",
          reducer: "sum",
          right: "Value #FREE (lastNotNull)"
        },
        mode: "binary",
        reduce: { reducer: "sum" }
      }
    })
    .withTransformation({
      id: "calculateField",
      options: {
        alias: "Used, %",
        binary: {
          left: "Used",
          operator: "/",
          reducer: "sum",
          right: "Value #TOTAL (lastNotNull)"
        },
        mode: "binary",
        reduce: { "reducer": "sum" }
      }
    })
    .withTransformation({
      id: "organize",
      options: {
        excludeByName: {},
        indexByName: {
          "Used": 3,
          "Used, %": 4,
          "Value #FREE (lastNotNull)": 2,
          "Value #TOTAL (lastNotNull)": 1,
          "mountpoint": 0,
        },
        renameByName: {
          "Value #FREE (lastNotNull)": "Available",
          "Value #TOTAL (lastNotNull)": "Size",
          "mountpoint": "Mounted on",
        }
      }
    })
    .withTransformation({
      id: "sortBy",
      options: {
        fields: {},
        sort: [
          { field: "Mounted on", desc: false }
        ]
      }
    })
    // Overrides
    .overrideByName("Mounted on", [{ id: "custom.width", value: "260" }])
    .overrideByName("Size", [{ id: "custom.width", value: "80" }])
    .overrideByName("Used", [{ id: "custom.width", value: "80" }])
    .overrideByName("Available", [{ id: "custom.width", value: "80" }])
    .overrideByName("Used, %", [
      {
        id: "custom.cellOptions",
        value: { mode: "basic", type: "gauge" }
      },
      { id: "min", value: 0 },
      { id: "max", value: 1 },
      { id: "unit", value: units.PercentUnit },
    ]);
};
