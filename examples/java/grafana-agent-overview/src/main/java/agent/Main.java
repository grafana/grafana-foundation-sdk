package agent;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.dashboard.*;

import javax.print.DocFlavor;
import java.util.List;

import static agent.Common.datasourceRef;
import static agent.Discovery.targetSyncTimeSeries;
import static agent.Discovery.targetsTimeSeries;
import static agent.Overview.runningInstancesTable;
import static agent.Retrieval.*;
import static com.grafana.foundation.common.Constants.TimeZoneBrowser;

public class Main {
    public static void main(String[] args) {
        Dashboard dashboard = new Dashboard.Builder("[Example] Grafana Agent Overview").
                uid("example-integration-agent").
                tags(List.of("generated", "grafana-foundation-sdk", "grafana-agent-integration")).
                editable().
                tooltip(DashboardCursorSync.OFF).
                refresh("30s").
                time(new DashboardDashboardTime.Builder().
                        from("now-30m").
                        to("now")
                ).
                timezone(TimeZoneBrowser).
                timepicker(new TimePickerConfig.Builder().
                        refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d")).
                        timeOptions(List.of("5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"))
                ).
                link(new DashboardLink.Builder("Grafana Agent Dashbaords").
                        type(DashboardLinkType.DASHBOARDS).
                        icon("external link").
                        tags(List.of("grafana-agent-integration")).
                        includeVars(true).
                        keepTime(true)
                ).
                withVariable(new VariableModel.DatasourceVariableBuilder("prometheus_datasource").
                        label("Data source").
                        type("prometheus").
                        regex("(?!grafanacloud-usage|grafanacloud-ml-metrics).+").
                        multi(false)).
                withVariable(jobQueryBuilder()).
                withVariable(instanceQueryBuilder()).
                withRow(new RowPanel.Builder("Overview")).
                withPanel(runningInstancesTable()).
                withRow(new RowPanel.Builder("Prometheus Discovery")).
                withPanel(targetSyncTimeSeries()).
                withPanel(targetsTimeSeries()).
                withRow(new RowPanel.Builder("Prometheus Retrieval")).
                withPanel(avgScrapeIntervalDurationTimeSeries()).
                withPanel(scrapeFailuresTimeSeries()).
                withPanel(appendedSamplesTimeSeries()).
                build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }

    }

    private static VariableModel.QueryVariableBuilder jobQueryBuilder() {
        VariableOption option = new VariableOption();
        option.selected = true;
        option.text = StringOrArrayOfString.createArrayOfString(List.of("All"));
        option.value = StringOrArrayOfString.createArrayOfString(List.of("$__all"));

        return new VariableModel.QueryVariableBuilder("job").
                label("Job").
                query(StringOrMap.createString("label_values(agent_build_info, job)")).
                datasource(datasourceRef()).
                current(option).
                refresh(VariableRefresh.ON_TIME_RANGE_CHANGED).
                sort(VariableSort.ALPHABETICAL_ASC).
                multi(true).
                includeAll(true);
    }

    private static VariableModel.QueryVariableBuilder instanceQueryBuilder() {
        VariableOption option = new VariableOption();
        option.selected = true;
        option.text = StringOrArrayOfString.createArrayOfString(List.of("All"));
        option.value = StringOrArrayOfString.createArrayOfString(List.of("$__all"));

        return new VariableModel.QueryVariableBuilder("instance").
                label("Instance").
                query(StringOrMap.createString("label_values(agent_build_info{job=~\\\"$job\\\"}, instance)")).
                datasource(datasourceRef()).
                current(option).
                refresh(VariableRefresh.ON_TIME_RANGE_CHANGED).
                sort(VariableSort.ALPHABETICAL_ASC).
                multi(true).
                includeAll(true);
    }
}
