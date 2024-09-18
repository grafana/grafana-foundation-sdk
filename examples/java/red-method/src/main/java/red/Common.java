package red;

import com.grafana.foundation.common.*;
import com.grafana.foundation.testdata.Datasource;
import com.grafana.foundation.timeseries.PanelBuilder;

public class Common {

    static PanelBuilder defaultTimeSeries() {
        return new PanelBuilder().
                lineWidth(1.0).
                fillOpacity(30.0).
                showPoints(VisibilityMode.NEVER).
                drawStyle(GraphDrawStyle.LINE).
                gradientMode(GraphGradientMode.OPACITY).
                spanNulls(BoolOrFloat64.createBool(false)).
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
