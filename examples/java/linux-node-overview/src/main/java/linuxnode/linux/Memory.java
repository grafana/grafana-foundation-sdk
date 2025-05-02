package linuxnode.linux;

import com.grafana.foundation.common.StackingConfigBuilder;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.dashboard.Threshold;
import com.grafana.foundation.dashboard.ThresholdsConfigBuilder;
import com.grafana.foundation.dashboard.ThresholdsMode;
import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.units.Constants;

import java.util.List;

public class Memory {
    public static PanelBuilder memoryUsageTimeseries() {
        String memUsedQuery = "(" +
                "  node_memory_MemTotal_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}" +
                "-" +
                "  node_memory_MemFree_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}" +
                "-" +
                "  node_memory_Buffers_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}" +
                "-" +
                "  node_memory_Cached_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}" +
                ")";

        return Common.defaultTimeSeries().
                title("Memory Usage").
                span(18).
                stacking(new StackingConfigBuilder().mode(StackingMode.NORMAL)).
                thresholds(new ThresholdsConfigBuilder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(
                                new Threshold(null, "green"),
                                new Threshold(80.0, "red")
                        ))
                ).
                min(0.0).
                unit(Constants.BytesIEC).
                decimals(2.0).
                withTarget(Common.basicPrometheusQuery(memUsedQuery, "Used")).
                withTarget(
                        Common.basicPrometheusQuery("node_memory_Buffers_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "Buffers")
                ).
                withTarget(
                        Common.basicPrometheusQuery("node_memory_Cached_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "Cached")
                ).
                withTarget(
                        Common.basicPrometheusQuery("node_memory_MemFree_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "Free")
                );
    }

    public static com.grafana.foundation.gauge.PanelBuilder memoryUsageGauge() {
        String query = "100 - (" +
                "  avg(node_memory_MemAvailable_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}) /" +
                "  avg(node_memory_MemTotal_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\"})" +
                "* 100)";

        return Common.defaultGauge().
                title("Memory Usage").
                span(6).
                min(30.0).
                max(100.0).
                unit(Constants.Percent).
                thresholds(new ThresholdsConfigBuilder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(
                                new Threshold(null, "rgba(50, 172, 45, 0.97)"),
                                new Threshold(80.0, "rgba(237, 129, 40, 0.89)"),
                                new Threshold(90.0, "rgba(245, 54, 54, 0.9)")
                        ))
                ).
                withTarget(Common.basicPrometheusQuery(query, ""));
    }
}
