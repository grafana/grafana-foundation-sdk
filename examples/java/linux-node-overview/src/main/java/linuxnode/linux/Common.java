package linuxnode.linux;

import com.grafana.foundation.common.*;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.prometheus.PromQueryFormat;
import com.grafana.foundation.timeseries.PanelBuilder;

import java.util.List;

public class Common {
    public static com.grafana.foundation.prometheus.DataqueryBuilder basicPrometheusQuery(String query, String legend) {
        return new com.grafana.foundation.prometheus.DataqueryBuilder().
                expr(query).
                legendFormat(legend);
    }

    public static com.grafana.foundation.loki.DataqueryBuilder basicLokiQuery(String query) {
        return new com.grafana.foundation.loki.DataqueryBuilder().expr(query);
    }

    public static com.grafana.foundation.prometheus.DataqueryBuilder tablePrometheusQuery(String query, String ref) {
        return new com.grafana.foundation.prometheus.DataqueryBuilder().
                expr(query).
                instant().
                legendFormat(PromQueryFormat.TABLE.Value()).
                refId(ref);
    }

    public static PanelBuilder defaultTimeSeries() {
        DataSourceRef ref = new DataSourceRef();
        ref.type = "loki";
        ref.uid = "grafana-cloud-logs";
        return new PanelBuilder().
                datasource(ref).
                lineWidth(1.0).
                fillOpacity(10.0).
                drawStyle(GraphDrawStyle.LINE).
                showPoints(VisibilityMode.NEVER).
                legend(new VizLegendOptionsBuilder().
                        showLegend(true).
                        placement(LegendPlacement.BOTTOM).
                        displayMode(LegendDisplayMode.LIST)
                );
    }

    public static com.grafana.foundation.logs.PanelBuilder defaultLogs() {
        return new com.grafana.foundation.logs.PanelBuilder().
                span(24).
                datasource(new DataSourceRef("loki", "grafana-cloud-logs")).
                showTime(true).
                enableLogDetails(true).
                sortOrder(LogsSortOrder.DESCENDING).
                wrapLogMessage(true);
    }

    public static com.grafana.foundation.gauge.PanelBuilder defaultGauge() {
        return new com.grafana.foundation.gauge.PanelBuilder().
                orientation(VizOrientation.AUTO).
                reduceOptions(new ReduceDataOptionsBuilder().
                        calcs(List.of("lastNotNull")).
                        values(false)
                );
    }
}
