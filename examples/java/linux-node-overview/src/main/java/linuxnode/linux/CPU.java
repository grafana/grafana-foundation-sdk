package linuxnode.linux;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.common.StackingConfig;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.dashboard.Panel;
import com.grafana.foundation.dashboard.Threshold;
import com.grafana.foundation.dashboard.ThresholdsConfig;
import com.grafana.foundation.dashboard.ThresholdsMode;

import java.util.List;

public class CPU {
    public static Builder<Panel> cpuUsageTimeseries() {
        String query = "((1 - sum without (mode) (rate(node_cpu_seconds_total{job=\"integrations/raspberrypi-node\", mode=~\"idle|iowait|steal\", instance=\"$instance\"}[$__rate_interval]))) " +
                "/ ignoring(cpu) group_left " +
                "count without (cpu, mode) (node_cpu_seconds_total{job=\"integrations/raspberrypi-node\", mode=\"idle\", instance=\"$instance\"}))";

        Threshold th1 = new Threshold();
        th1.color = "green";

        Threshold th2 = new Threshold();
        th2.color = "red";
        th2.value = 80.0;

        return Common.defaultTimeSeries().
                title("CPU Usage").
                span(18).
                stacking(new StackingConfig.Builder().mode(StackingMode.NORMAL)).
                thresholds(new ThresholdsConfig.Builder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(th1, th2))).
                max(1.0).
                min(0.0).
                withTarget(Common.basicPrometheusQuery(query, "{{ cpu }}"));
    }

    public static Builder<Panel> loadAverageTimeseries() {
        Threshold th1 = new Threshold();
        th1.color = "green";

        Threshold th2 = new Threshold();
        th2.value = 80.0;
        th2.color = "red";
        return Common.defaultTimeSeries().title("Load Average").
                span(18).
                thresholds(
                        new ThresholdsConfig.Builder().
                                mode(ThresholdsMode.ABSOLUTE).
                                steps(List.of(th1, th2))
                ).
                min(0.0).
                unit("short").
                withTarget(Common.basicPrometheusQuery("node_load1{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "1m load average")).
                withTarget(Common.basicPrometheusQuery("node_load5{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "5m load average")).
                withTarget(Common.basicPrometheusQuery("node_load15{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", "15m load average")).
                withTarget(Common.basicPrometheusQuery("count(node_cpu_seconds_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", mode=\"idle\"})", "logical cores"));
    }

    public static Builder<Panel> cpuTemperatureGauge() {
        Threshold th1 = new Threshold();
        th1.color = "rgba(50, 172, 45, 0.97)";

        Threshold th2 = new Threshold();
        th2.value = 65.0;
        th2.color = "rgba(237, 129, 40, 0.89)";

        Threshold th3 = new Threshold();
        th3.value = 85.0;
        th3.color = "rgba(245, 54, 54, 0.9)";
        return Common.defaultGauge().
                title("CPU Temperature").
                span(6).
                min(30.0).
                max(100.0).
                unit("celsius").
                thresholds(new ThresholdsConfig.Builder().
                        mode(ThresholdsMode.ABSOLUTE).
                        steps(List.of(th1, th2, th3))).
                withTarget(Common.basicPrometheusQuery("avg(node_hwmon_temp_celsius{job=\"integrations/raspberrypi-node\", instance=\"$instance\"}", ""));
    }
}
