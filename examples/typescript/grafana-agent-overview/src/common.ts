import * as common from '@grafana/grafana-foundation-sdk/common';
import * as prometheus from '@grafana/grafana-foundation-sdk/prometheus';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

export const prometheusQuery = (query: string, legend: string): prometheus.DataqueryBuilder => {
  return new prometheus.DataqueryBuilder()
    .expr(query)
    .legendFormat(legend);
};

export const tablePrometheusQuery = (query: string, ref: string): prometheus.DataqueryBuilder => {
  return new prometheus.DataqueryBuilder()
    .expr(query)
    .instant()
    .format(prometheus.PromQueryFormat.Table)
    .refId(ref);
};

export const defaultTimeseries = (): TimeseriesBuilder => {
  return new TimeseriesBuilder()
    .height(7)
    .span(12)
    .lineWidth(1)
    .fillOpacity(0)
    .pointSize(5)
    .showPoints(common.VisibilityMode.Auto)
    .drawStyle(common.GraphDrawStyle.Line)
    .gradientMode(common.GraphGradientMode.None)
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
        .mode(common.TooltipDisplayMode.Single)
        .sort(common.SortOrder.None)
    )
    .thresholdsStyle(
      new common.GraphThresholdsStyleConfigBuilder()
        .mode(common.GraphThresholdsStyleMode.Off)
    );
};
