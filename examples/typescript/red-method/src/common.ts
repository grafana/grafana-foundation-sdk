import * as common from '@grafana/grafana-foundation-sdk/common';
import { PanelBuilder as TimeseriesBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

export const defaultTimeseries = (): TimeseriesBuilder => {
  return new TimeseriesBuilder()
    .lineWidth(1)
    .fillOpacity(30)
    .showPoints(common.VisibilityMode.Never)
    .drawStyle(common.GraphDrawStyle.Line)
    .gradientMode(common.GraphGradientMode.Opacity)
    .spanNulls(false)
    .axisBorderShow(false)
    .lineInterpolation(common.LineInterpolation.Smooth)
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
