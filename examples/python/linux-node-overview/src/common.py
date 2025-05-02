from grafana_foundation_sdk.builders import (
    dashboard,
    timeseries,
    common as common_builder,
    prometheus,
    stat,
)
from grafana_foundation_sdk.models import common, prometheus as prom
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    VariableRefresh,
    VariableOption,
    VariableSort,
)


def query_variable(name: str, label: str, query: str) -> dashboard.QueryVariable:
    return (
        dashboard.QueryVariable(name)
        .label(label)
        .query(query)
        .datasource(DataSourceRef(uid="$datasource"))
        .current(VariableOption(selected=True, text="All", value="$__all"))
        .refresh(VariableRefresh.ON_TIME_RANGE_CHANGED)
        .sort(VariableSort.ALPHABETICAL_ASC)
        .multi(True)
        .include_all(True)
    )


def prometheus_query(query: str, legend: str = "") -> prometheus.Dataquery:
    return prometheus.Dataquery().expr(query).legend_format(legend)


def table_prometheus_query(query: str, ref_id: str) -> prometheus.Dataquery:
    return (
        prometheus.Dataquery()
        .expr(query)
        .format(prom.PromQueryFormat.TABLE)
        .ref_id(ref_id)
    )


def default_stat() -> stat.Panel:
    return (
        stat.Panel()
        .datasource(DataSourceRef(uid="$datasource"))
        .height(2)
        .span(6)
        .decimals(1)
        .reduce_options(common_builder.ReduceDataOptions().calcs(["lastNotNull"]))
        .color_mode(common.BigValueColorMode.VALUE)
        .graph_mode(common.BigValueGraphMode.NONE)
    )


def uname_stat(title: str, description: str, field: str) -> stat.Panel:
    return (
        default_stat()
        .title(title)
        .description(description)
        .with_target(
            table_prometheus_query(
                query='node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}',
                ref_id="A",
            )
        )
        .reduce_options(
            common_builder.ReduceDataOptions()
            .calcs(["lastNotNull"])
            .fields("/^%s$/" % field)
        )
        .color_mode(common.BigValueColorMode.NONE)
    )


def default_timeseries() -> timeseries.Panel:
    return (
        timeseries.Panel()
        .datasource(DataSourceRef(uid="$datasource"))
        .line_width(2)
        .fill_opacity(30)
        .line_interpolation(common.LineInterpolation.SMOOTH)
        .show_points(common.VisibilityMode.NEVER)
        .draw_style(common.GraphDrawStyle.LINE)
        .gradient_mode(common.GraphGradientMode.SCHEME)
        .span_nulls(False)
        .axis_border_show(False)
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
