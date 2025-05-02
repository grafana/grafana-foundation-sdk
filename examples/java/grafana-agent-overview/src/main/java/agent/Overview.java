package agent;

import com.grafana.foundation.common.TableFooterOptionsBuilder;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.DynamicConfigValue;
import com.grafana.foundation.table.PanelBuilder;
import com.grafana.foundation.units.Constants;

import java.util.List;
import java.util.Map;

import static agent.Common.datasourceRef;
import static agent.Common.tablePrometheusQuery;

public class Overview {
    static PanelBuilder runningInstancesTable() {
        return new PanelBuilder().
                title("Running instances").
                description("General statistics of running grafana agent instances.").
                height(7).
                span(24).
                footer(new TableFooterOptionsBuilder().
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
                overrideByName("Value #B", List.of(
                        new DynamicConfigValue("unit", Constants.Seconds)
                ));
    }

    private static DataTransformerConfig transformationConfig(String id, Map<String, Object> map) {
        DataTransformerConfig cfg = new DataTransformerConfig();
        cfg.id = id;
        cfg.options = map;
        return cfg;
    }
}
