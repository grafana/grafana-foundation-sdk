package agent;

import com.grafana.foundation.common.TableFooterOptions;
import com.grafana.foundation.dashboard.DashboardFieldConfigSourceOverrides;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.DynamicConfigValue;
import com.grafana.foundation.dashboard.MatcherConfig;
import com.grafana.foundation.table.PanelBuilder;

import java.util.List;
import java.util.Map;

import static agent.Common.datasourceRef;
import static agent.Common.tablePrometheusQuery;

public class Overview {
    static PanelBuilder runningInstancesTable() {
        MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = "Value #B";

        DynamicConfigValue dynamicConfigValue = new DynamicConfigValue();
        dynamicConfigValue.id = "unit";
        dynamicConfigValue.value = "s";

        return new PanelBuilder().
                title("Running instances").
                description("General statistics of running grafana agent instances.").
                height(7).
                span(24).
                footer(new TableFooterOptions.Builder().
                        countRows(false).
                        reducer(List.of("sum"))
                ).
                datasource(datasourceRef()).
                withTarget(
                        tablePrometheusQuery("count by (instance, version) (agent_build_info{job=~\"$job\", instance=~\"$instance\"})", "A")
                ).
                withTarget(
                        tablePrometheusQuery("max by (instance) (time() - process_start_time_seconds{job=~\"$job\", instance=~\"$instance\"})", "B")
                ).
                withTransformation(transformationConfig("merge", Map.of())).
                withTransformation(transformationConfig("organize", Map.of(
                        "excludeByName", Map.of("Time", true, "Value #A", true),
                        "renameByName", Map.of("Value #B", "Uptime")))
                ).
                withOverride(new DashboardFieldConfigSourceOverrides.Builder().matcher(matcherConfig).properties(List.of(dynamicConfigValue)));
    }

    private static DataTransformerConfig transformationConfig(String id, Map<String, Object> map) {
        DataTransformerConfig cfg = new DataTransformerConfig();
        cfg.id = id;
        cfg.options = map;
        return cfg;
    }
}
