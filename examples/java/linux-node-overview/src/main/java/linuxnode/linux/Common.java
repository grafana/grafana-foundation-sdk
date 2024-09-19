package linuxnode.linux;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.common.*;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.prometheus.PromQueryFormat;
import com.grafana.foundation.timeseries.PanelBuilder;

import java.util.List;

public class Common {

    // ??
    public static Builder<Dataquery> basicPrometheusQuery(String query, String legend) {
        return new com.grafana.foundation.prometheus.Dataquery.Builder().
                expr(query).
                legendFormat(legend);
    }

    public static Builder<Dataquery> basicLokiQuery(String query) {
        return new com.grafana.foundation.loki.Dataquery.Builder().expr(query);
    }

    // ??
    public static Builder<Dataquery> tablePrometheusQuery(String query, String ref) {
        return new com.grafana.foundation.prometheus.Dataquery.Builder().
                expr(query).
                instant().
                legendFormat(PromQueryFormat.TABLE.Value()).
                refId(ref);
    }

    public static PanelBuilder defaultTimeSeries() {
        return new PanelBuilder().
                lineWidth(1.0).
                fillOpacity(10.0).
                drawStyle(GraphDrawStyle.LINE).
                showPoints(VisibilityMode.NEVER).
                legend(new VizLegendOptions.Builder().
                        showLegend(true).
                        placement(LegendPlacement.BOTTOM).
                        displayMode(LegendDisplayMode.LIST)
                );
    }

    public static com.grafana.foundation.logs.PanelBuilder defaultLogs() {
        DataSourceRef ref = new DataSourceRef();
        ref.type = "loki";
        ref.uid = "grafana-cloud-logs";
        return new com.grafana.foundation.logs.PanelBuilder().
                span(24).
                datasource(ref).
                showTime(true).
                enableLogDetails(true).
                sortOrder(LogsSortOrder.DESCENDING).
                wrapLogMessage(true);
    }

    public static com.grafana.foundation.gauge.PanelBuilder defaultGauge() {
        return new com.grafana.foundation.gauge.PanelBuilder().
                orientation(VizOrientation.AUTO).
                reduceOptions(new ReduceDataOptions.Builder().
                        calcs(List.of("lastNotNull")).
                        values(false)
                );
    }
}
