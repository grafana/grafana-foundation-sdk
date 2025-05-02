package linuxnode.linux;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.dashboard.*;

import java.util.List;

public class NodeExporter {
    public static Dashboard Build() {
        return new DashboardBuilder("[TEST] Node Exporter / Raspberry").
                uid("test-dashboard-raspberry").
                tags(List.of("generated", "raspberrypi-node-integration")).
                refresh("30s").
                time(new DashboardDashboardTimeBuilder().from("now-30m").to("now")).
                timezone("browser").
                timepicker(new TimePickerBuilder().
                        refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"))
                ).
                tooltip(DashboardCursorSync.CROSSHAIR).
                withVariable(datasourceVariable()).
                withVariable(queryVariable()).
                // CPU
                withRow(new RowBuilder("CPU")).
                withPanel(CPU.cpuUsageTimeseries()).
                withPanel(CPU.cpuTemperatureGauge()).
                withPanel(CPU.loadAverageTimeseries()).
                // Memory
                withRow(new RowBuilder("Memory")).
                withPanel(Memory.memoryUsageTimeseries()).
                withPanel(Memory.memoryUsageGauge()).
                // Disk
                withRow(new RowBuilder("Disk")).
                withPanel(Disk.diskIOTimeseries()).
                withPanel(Disk.diskSpaceUsageTable()).
                // Network
                withRow(new RowBuilder("Network")).
                withPanel(Network.networkReceivedTimeseries()).
                withPanel(Network.networkTransmittedTimeseries()).
                // Logs
                withRow(new RowBuilder("Logs")).
                withPanel(Logs.errorsInSystemLogs()).
                withPanel(Logs.authLogs()).
                withPanel(Logs.kernelLogs()).
                withPanel(Logs.allSystemLogs()).
                build();
    }

    private static Builder<VariableModel> datasourceVariable() {
        return new DatasourceVariableBuilder("datasource").
                label("Data Source").
                hide(VariableHide.DONT_HIDE).
                type("prometheus").
                current(new VariableOption(
                        true,
                        StringOrArrayOfString.createString("grafanacloud-potatopi-prom"),
                        StringOrArrayOfString.createString("grafanacloud-prom")
                ));
    }

    private static Builder<VariableModel> queryVariable() {
        return new QueryVariableBuilder("instance").
                label("Instance").
                hide(VariableHide.DONT_HIDE).
                refresh(VariableRefresh.ON_TIME_RANGE_CHANGED).
                query(StringOrMap.createString("label_values(node_uname_info{job=\"integrations/raspberrypi-node\", sysname!=\"Darwin\"}, instance)")).
                datasource(new DataSourceRef("prometheus", "$datasource")).
                current(new VariableOption(
                        false,
                        StringOrArrayOfString.createString("potato"),
                        StringOrArrayOfString.createString("potato")
                )).
                sort(VariableSort.DISABLED);
    }
}
