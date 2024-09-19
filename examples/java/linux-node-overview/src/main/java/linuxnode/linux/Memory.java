package linuxnode.linux;

import com.grafana.foundation.common.StackingConfig;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.dashboard.Threshold;
import com.grafana.foundation.dashboard.ThresholdsConfig;
import com.grafana.foundation.dashboard.ThresholdsMode;
import com.grafana.foundation.timeseries.PanelBuilder;

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

        Threshold th1 = new Threshold();
        th1.color = "green";

        Threshold th2 = new Threshold();
        th2.color = "red";
        th2.value = 80.0;

        return Common.defaultTimeSeries().
                title("Memory Usage").
                span(18).
                stacking(new StackingConfig.Builder().mode(StackingMode.NORMAL)).
                thresholds(new ThresholdsConfig.Builder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(th1, th2))
                ).
                min(0.0).
                unit("bytes").
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

        Threshold th1 = new Threshold();
        th1.color = "rgba(50, 172, 45, 0.97)";

        Threshold th2 = new Threshold();
        th2.value = 80.0;
        th2.color = "rgba(237, 129, 40, 0.89)";

        Threshold th3 = new Threshold();
        th3.value = 90.0;
        th3.color = "rgba(245, 54, 54, 0.9)";

        return Common.defaultGauge().
                title("Memory Usage").
                span(6).
                min(30.0).
                max(100.0).
                unit("percent").
                thresholds(new ThresholdsConfig.Builder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(th1, th2, th3))
                ).
                withTarget(Common.basicPrometheusQuery(query, ""));
    }
}
