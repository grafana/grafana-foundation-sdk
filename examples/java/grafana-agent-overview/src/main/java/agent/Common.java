package agent;

import com.grafana.foundation.common.*;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.prometheus.DataqueryBuilder;
import com.grafana.foundation.prometheus.PromQueryFormat;
import com.grafana.foundation.timeseries.PanelBuilder;

public class Common {
    static DataSourceRef datasourceRef() {
        return new DataSourceRef("prometheus", "$prometheus_datasource");
    }

    static DataqueryBuilder prometheusQuery(String query, String refId) {
        return new DataqueryBuilder().expr(query).refId(refId);
    }

    static DataqueryBuilder tablePrometheusQuery(String query, String refId) {
        return new DataqueryBuilder().
                expr(query).
                format(PromQueryFormat.TABLE).
                instant().
                intervalFactor(2.0).
                refId(refId);
    }

    static PanelBuilder defaultTimeSeries() {
        return new PanelBuilder().
                height(7).
                span(12).
                lineWidth(1.0).
                fillOpacity(0.0).
                pointSize(5.0).
                showPoints(VisibilityMode.AUTO).
                drawStyle(GraphDrawStyle.LINE).
                gradientMode(GraphGradientMode.NONE).
                legend(new VizLegendOptionsBuilder().
                        displayMode(LegendDisplayMode.LIST).
                        placement(LegendPlacement.BOTTOM).
                        showLegend(true)
                ).
                tooltip(new VizTooltipOptionsBuilder().
                        mode(TooltipDisplayMode.SINGLE).
                        sort(SortOrder.NONE)
                ).
                thresholdsStyle(new GraphThresholdsStyleConfigBuilder().
                        mode(GraphThresholdsStyleMode.OFF)
                ).
                spanNulls(BoolOrFloat64.createBool(false))
                .axisBorderShow(false);
    }
}
