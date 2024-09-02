package red;

import com.grafana.foundation.common.*;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.timeseries.PanelBuilder;

public class Common {

    static DataSourceRef datasourceRef() {
        DataSourceRef ref = new DataSourceRef();
        ref.uid = "grafana";
        ref.type = "grafana";
        return ref;
    }

    static PanelBuilder defaultTimeSeries() {
        BoolOrFloat64 spanNulls = new BoolOrFloat64();
        spanNulls.bool = false;

        return new PanelBuilder().
                lineWidth(1.0).
                fillOpacity(30.0).
                showPoints(VisibilityMode.NEVER).
                drawStyle(GraphDrawStyle.LINE).
                gradientMode(GraphGradientMode.OPACITY).
                spanNulls(spanNulls).
                axisBorderShow(false).
                lineInterpolation(LineInterpolation.SMOOTH).
                legend(new VizLegendOptions.Builder().
                        displayMode(LegendDisplayMode.LIST).
                        placement(LegendPlacement.BOTTOM).
                        showLegend(true)
                ).
                tooltip(new VizTooltipOptions.Builder().
                        mode(TooltipDisplayMode.MULTI).
                        sort(SortOrder.DESCENDING)
                ).
                thresholdsStyle(new GraphThresholdsStyleConfig.Builder().
                        mode(GraphThresholdsStyleMode.OFF)
                );
    }
}
