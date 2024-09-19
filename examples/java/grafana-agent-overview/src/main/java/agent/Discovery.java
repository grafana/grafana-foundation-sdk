package agent;

import com.grafana.foundation.timeseries.PanelBuilder;

import static agent.Common.*;

public class Discovery {
    static PanelBuilder targetSyncTimeSeries() {
        return defaultTimeSeries().
                title("Target Sync").
                description("Actual interval to sync the scrape pool.").
                datasource(datasourceRef()).
                unit("s").
                withTarget(
                        prometheusQuery(
                                "sum(rate(prometheus_target_sync_length_seconds_sum{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval])) by (instance, scrape_job)",
                                "{{instance}}/{{scrape_job}}"
                        )
                );
    }

    static PanelBuilder targetsTimeSeries() {
        return defaultTimeSeries().
                title("Targets").
                description("Discovered targets by prometheus service discovery.").
                datasource(datasourceRef()).
                unit("short").
                withTarget(
                        prometheusQuery(
                                "sum by (instance) (prometheus_sd_discovered_targets{job=~\"$job\", instance=~\"$instance\"})",
                                "{{instance}}"
                        )
                );
    }
}
