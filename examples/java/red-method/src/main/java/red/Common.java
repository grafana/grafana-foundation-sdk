package red;

import com.grafana.foundation.common.*;
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
                legend(new VizLegendOptionsBuilder().
                        displayMode(LegendDisplayMode.LIST).
                        placement(LegendPlacement.BOTTOM).
                        showLegend(true)
                ).
                tooltip(new VizTooltipOptionsBuilder().
                        mode(TooltipDisplayMode.MULTI).
                        sort(SortOrder.DESCENDING)
                ).
                thresholdsStyle(new GraphThresholdsStyleConfigBuilder().
                        mode(GraphThresholdsStyleMode.OFF)
                );
    }
}
