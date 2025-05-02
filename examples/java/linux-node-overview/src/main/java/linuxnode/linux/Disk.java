package linuxnode.linux;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.common.FieldTextAlignment;
import com.grafana.foundation.common.TableCellHeight;
import com.grafana.foundation.common.TableFooterOptionsBuilder;
import com.grafana.foundation.dashboard.*;
import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.units.Constants;

import java.util.List;
import java.util.Map;

public class Disk {
    public static PanelBuilder diskIOTimeseries() {
        return Common.defaultTimeSeries().
                title("Disk I/O").
                fillOpacity(0.0).
                unit(Constants.BytesPerSecondSI).
                withTarget(Common.basicPrometheusQuery("rate(node_disk_read_bytes_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", device!=\"\"}[$__rate_interval])", "{{device}} read")).
                withTarget(Common.basicPrometheusQuery("rate(node_disk_written_bytes_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", device!=\"\"}[$__rate_interval])", "{{device}} written")).
                withTarget(Common.basicPrometheusQuery("rate(node_disk_io_time_seconds_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", device!=\"\"}[$__rate_interval])", "{{device}} IO time")).
                overrideByRegexp("/ io time/", List.of(
                        new DynamicConfigValue("unit", Constants.PercentUnit)
                ));
    }

    public static com.grafana.foundation.table.PanelBuilder diskSpaceUsageTable() {
        return new com.grafana.foundation.table.PanelBuilder().
                title("Disk Space Usage").
                align(FieldTextAlignment.AUTO).
                unit(Constants.BytesSI).
                cellHeight(TableCellHeight.SM).
                footer(new TableFooterOptionsBuilder().
                        countRows(false).
                        reducer(List.of("sum"))
                ).
                withTarget(Common.tablePrometheusQuery("max by (mountpoint) (node_filesystem_size_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\", fstype!=\"\"})", "A")).
                withTarget(Common.tablePrometheusQuery("max by (mountpoint) (node_filesystem_avail_bytes{job=\"integrations/raspberrypi-node\", instance=\"$instance\", fstype!=\"\"})", "B")).
                withTransformation(transformer1()).
                withTransformation(transformer2()).
                withTransformation(transformer3()).
                withTransformation(transformer4()).
                withTransformation(transformer5()).
                withTransformation(transformer6()).
                withOverride(defaultOverrides("Mounted on", 260)).
                withOverride(defaultOverrides("Size", 93)).
                withOverride(defaultOverrides("Used", 72)).
                withOverride(defaultOverrides("Available", 88)).
                withOverride(complexOverrides());
    }

    private static DataTransformerConfig transformer1() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "groupBy";
        dataTransformerConfig.options = Map.of(
                "fields", Map.of(
                        "Value #A", Map.of(
                                "aggregations", List.of("lastNotNull"),
                                "operation", "aggregate"
                        ),
                        "Value #B", Map.of(
                                "aggregations", List.of("lastNotNull"),
                                "operation", "aggregate"
                        ),
                        "mountpoint", Map.of(
                                "aggregations", List.of(),
                                "operation", "groupby"
                        )
                )
        );

        return dataTransformerConfig;
    }

    private static DataTransformerConfig transformer2() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "merge";

        return dataTransformerConfig;
    }

    private static DataTransformerConfig transformer3() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "calculateField";
        dataTransformerConfig.options = Map.of(
                "alias", "Used",
                "binary", Map.of(
                        "left", "Value #A (lastNotNull)",
                        "operator", "-",
                        "reducer", "sum",
                        "right", "Value #B (lastNotNull)"
                ), "mode", "binary",
                "reduce", Map.of("reducer", "sum"));

        return dataTransformerConfig;
    }

    private static DataTransformerConfig transformer4() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "calculateField";
        dataTransformerConfig.options = Map.of(
                "alias", "Used, %",
                "binary", Map.of(
                        "left", "Used",
                        "operator", "/",
                        "reducer", "sum",
                        "right", "Value #A (lastNotNull)"
                ),
                "mode", "binary",
                "reduce", Map.of("reducer", "sum"));

        return dataTransformerConfig;
    }

    private static DataTransformerConfig transformer5() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "organize";
        dataTransformerConfig.options = Map.of(
                "excludeByName", List.of(),
                "indexByName", List.of(),
                "renameByName", Map.of(
                        "Value #A (lastNotNull)", "Size",
                        "Value #B (lastNotNull)", "Available",
                        "mountpoint", "Mounted on"
                )
        );

        return dataTransformerConfig;
    }

    private static DataTransformerConfig transformer6() {
        DataTransformerConfig dataTransformerConfig = new DataTransformerConfig();
        dataTransformerConfig.id = "sortBy";
        dataTransformerConfig.options = Map.of(
                "fields", List.of(),
                "sort", Map.of("field", "Mounted on")
        );

        return dataTransformerConfig;
    }

    private static Builder<DashboardFieldConfigSourceOverrides> defaultOverrides(String options, Integer value) {
        MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = options;

        DynamicConfigValue dynamicConfigValue = new DynamicConfigValue();
        dynamicConfigValue.id = "custom.width";
        dynamicConfigValue.value = value;
        return new DashboardFieldConfigSourceOverridesBuilder().matcher(matcherConfig).properties(List.of(dynamicConfigValue));
    }

    private static Builder<DashboardFieldConfigSourceOverrides> complexOverrides() {
        MatcherConfig matcherConfig = new MatcherConfig();
        matcherConfig.id = "byName";
        matcherConfig.options = "Used, %";

        DynamicConfigValue dynamicConfigValue1 = new DynamicConfigValue();
        dynamicConfigValue1.id = "unit";
        dynamicConfigValue1.value = "percentunit";

        DynamicConfigValue dynamicConfigValue2 = new DynamicConfigValue();
        dynamicConfigValue2.id = "custom.cellOptions";
        dynamicConfigValue2.value = Map.of("mode", "gradient", "type", "gauge");

        DynamicConfigValue dynamicConfigValue3 = new DynamicConfigValue();
        dynamicConfigValue3.id = "min";
        dynamicConfigValue3.value = 0;

        DynamicConfigValue dynamicConfigValue4 = new DynamicConfigValue();
        dynamicConfigValue4.id = "max";
        dynamicConfigValue4.value = 1;

        return new DashboardFieldConfigSourceOverridesBuilder().
                matcher(matcherConfig).
                properties(List.of(dynamicConfigValue1, dynamicConfigValue2, dynamicConfigValue3, dynamicConfigValue4));
    }
}
