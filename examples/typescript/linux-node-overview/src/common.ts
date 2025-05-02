import * as common from '@grafana/grafana-foundation-sdk/common';
import * as prometheus from '@grafana/grafana-foundation-sdk/prometheus';
import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as StatBuilder } from '@grafana/grafana-foundation-sdk/stat';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

export const queryVariable = (name: string, label: string, query: string): dashboard.QueryVariableBuilder => {
  return new dashboard.QueryVariableBuilder(name)
    .label(label)
    .query(query)
    .datasource({ uid: "$datasource" })
    .current({
      selected: true,
      text: "All",
      value: "$__all",
    })
    .refresh(dashboard.VariableRefresh.OnTimeRangeChanged)
    .sort(dashboard.VariableSort.AlphabeticalAsc)
    .multi(true)
    .includeAll(true);
};

export const prometheusQuery = (query: string, legend: string): prometheus.DataqueryBuilder => {
  return new prometheus.DataqueryBuilder()
    .datasource({ uid: "$datasource" })
    .expr(query)
    .legendFormat(legend);
};

export const tablePrometheusQuery = (query: string, ref: string): prometheus.DataqueryBuilder => {
  return new prometheus.DataqueryBuilder()
    .datasource({ uid: "$datasource" })
    .expr(query)
    .format(prometheus.PromQueryFormat.Table)
    .refId(ref);
};

export const defaultStat = (): StatBuilder => {
  return new StatBuilder()
    .height(2)
    .span(6)
    .decimals(1)
    .reduceOptions(
      new common.ReduceDataOptionsBuilder()
        .calcs(["lastNotNull"])
    )
    .colorMode(common.BigValueColorMode.Value)
    .graphMode(common.BigValueGraphMode.None);
};

export const unameStat = (panelTitle: string, panelDescription: string, field: string): StatBuilder => {
  return defaultStat()
    .title(panelTitle)
    .description(panelDescription)
    .withTarget(
      tablePrometheusQuery('node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'A')
    )
    .reduceOptions(
      new common.ReduceDataOptionsBuilder()
        .calcs(["lastNotNull"])
        .fields(`/^${field}$/`)
    )
    .colorMode(common.BigValueColorMode.None);
};

export const defaultTimeseries = (): TimeseriesBuilder => {
  return new TimeseriesBuilder()
    .lineWidth(2)
    .fillOpacity(30)
    .lineInterpolation(common.LineInterpolation.Smooth)
    .showPoints(common.VisibilityMode.Never)
    .drawStyle(common.GraphDrawStyle.Line)
    .gradientMode(common.GraphGradientMode.Scheme)
    .spanNulls(false)
    .axisBorderShow(false)
    .legend(
      new common.VizLegendOptionsBuilder()
        .showLegend(true)
        .placement(common.LegendPlacement.Bottom)
        .displayMode(common.LegendDisplayMode.List)
    )
    .tooltip(
      new common.VizTooltipOptionsBuilder()
        .mode(common.TooltipDisplayMode.Multi)
        .sort(common.SortOrder.Descending)
    )
    .thresholdsStyle(
      new common.GraphThresholdsStyleConfigBuilder()
        .mode(common.GraphThresholdsStyleMode.Off)
    );
};
