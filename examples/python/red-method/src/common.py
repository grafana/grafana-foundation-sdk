from grafana_foundation_sdk.builders import timeseries, common as common_builder
from grafana_foundation_sdk.models import common


def default_timeseries() -> timeseries.Panel:
    return (
        timeseries.Panel()
        .line_width(1)
        .fill_opacity(30)
        .show_points(common.VisibilityMode.NEVER)
        .draw_style(common.GraphDrawStyle.LINE)
        .gradient_mode(common.GraphGradientMode.OPACITY)
        .span_nulls(False)
        .line_interpolation(common.LineInterpolation.SMOOTH)
        .legend(
            common_builder.VizLegendOptions()
            .display_mode(common.LegendDisplayMode.LIST)
            .placement(common.LegendPlacement.BOTTOM)
            .show_legend(True)
        )
        .tooltip(
            common_builder.VizTooltipOptions()
            .mode(common.TooltipDisplayMode.MULTI)
            .sort(common.SortOrder.DESCENDING)
        )
        .thresholds_style(
            common_builder.GraphThresholdsStyleConfig().mode(
                common.GraphThresholdsStyleMode.OFF
            )
        )
    )
