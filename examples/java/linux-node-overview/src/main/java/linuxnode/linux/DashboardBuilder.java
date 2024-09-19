package linuxnode.linux;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.dashboard.*;

import java.util.List;

public class DashboardBuilder {
    public static Dashboard Build() {
        return new Dashboard.Builder("[TEST] Node Exporter / Raspberry").
                uid("test-dashboard-raspberry").
                tags(List.of("generated", "raspberrypi-node-integration")).
                refresh("30s").
                time(new DashboardDashboardTime.Builder().from("now-30m").to("now")).
                timezone("browser").
                timepicker(new TimePickerConfig.Builder().
                        refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d")).
                        timeOptions(List.of("5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"))
                ).
                tooltip(DashboardCursorSync.CROSSHAIR).
                withVariable(datasourceVariable()).
                withVariable(queryVariable()).
                // CPU
                withRow(new RowPanel.Builder("CPU")).
                withPanel(CPU.cpuUsageTimeseries()).
                withPanel(CPU.cpuTemperatureGauge()).
                withPanel(CPU.loadAverageTimeseries()).
                // Memory
                withRow(new RowPanel.Builder("Memory")).
                withPanel(Memory.memoryUsageTimeseries()).
                withPanel(Memory.memoryUsageGauge()).
                // Disk
                withRow(new RowPanel.Builder("Disk")).
                withPanel(Disk.diskIOTimeseries()).
                withPanel(Disk.diskSpaceUsageTable()).
                // Network
                withRow(new RowPanel.Builder("Network")).
                withPanel(Network.networkReceivedTimeseries()).
                withPanel(Network.networkTransmittedTimeseries()).
                // Logs
                withRow(new RowPanel.Builder("Logs")).
                withPanel(Logs.errorsInSystemLogs()).
                withPanel(Logs.authLogs()).
                withPanel(Logs.kernelLogs()).
                withPanel(Logs.allSystemLogs()).
                build();
    }

    private static Builder<VariableModel> datasourceVariable() {
        VariableOption current = new VariableOption();
        current.selected = true;
        current.text = StringOrArrayOfString.createString("grafanacloud-potatopi-prom");
        current.value = StringOrArrayOfString.createString("grafanacloud-prom");

        return new VariableModel.DatasourceVariableBuilder("datasource").
                label("Data Source").
                hide(VariableHide.DONT_HIDE).
                type("prometheus").
                current(current);
    }

    private static Builder<VariableModel> queryVariable() {
        VariableOption current = new VariableOption();
        current.selected = false;
        current.text = StringOrArrayOfString.createString("potato");
        current.value = StringOrArrayOfString.createString("potato");

        DataSourceRef datasource = new DataSourceRef();
        datasource.uid = "$datasource";
        datasource.type = "prometheus";

        return new VariableModel.QueryVariableBuilder("instance").
                label("Instance").
                hide(VariableHide.DONT_HIDE).
                refresh(VariableRefresh.ON_TIME_RANGE_CHANGED).
                query(StringOrMap.createString("label_values(node_uname_info{job=\"integrations/raspberrypi-node\", sysname!=\"Darwin\"}, instance)")).
                datasource(datasource).
                current(current).
                sort(VariableSort.DISABLED);
    }

}
