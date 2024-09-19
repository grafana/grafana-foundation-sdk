package red;

import com.grafana.foundation.testdata.Dataquery;
import com.grafana.foundation.testdata.Datasource;
import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.dashboard.*;

import java.util.List;

import static com.grafana.foundation.common.Constants.TimeZoneBrowser;
import static red.Common.defaultTimeSeries;

public class Red {
    static Dashboard.Builder red(String dashboardTitle, List<String> serviceIDs) {
        Dashboard.Builder builder = new Dashboard.Builder("[Example] " + dashboardTitle).
                uid("example-red-method").
                tags(List.of("generated", "red")).
                editable().
                tooltip(DashboardCursorSync.CROSSHAIR).
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
                link(new DashboardLink.Builder("RED method").
                        type(DashboardLinkType.LINK).icon("question").
                        targetBlank(true).
                        url("https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method")
                );

        for (String serviceId : serviceIDs) {
            builder.withRow(new RowPanel.Builder(serviceId)).
                    withPanel(requestRateTimeSeries().span(8).height(8)).
                    withPanel(errorRateTimeSeries().span(8).height(8)).
                    withPanel(durationRateTimeSeries().span(8).height(8));
        }

        return builder;
    }

    private static PanelBuilder requestRateTimeSeries() {
        return defaultTimeSeries().
                title("Request rate").
                description("Number of requests handled by the service, per second.").
                unit("reqps").
                withTarget(new Dataquery.Builder().
                        queryType("randomWalk").
                        datasource(new Datasource.Builder().
                                type("grafana").
                                uid("grafana"))
                );
    }

    private static PanelBuilder errorRateTimeSeries() {
        return defaultTimeSeries().
                title("Error rate").
                description("Number of failed requests, per second.").
                unit("reqps").
                withTarget(new Dataquery.Builder().
                        queryType("randomWalk").
                        datasource(new Datasource.Builder().
                                type("grafana").
                                uid("grafana"))
                );
    }

    private static PanelBuilder durationRateTimeSeries() {
        return defaultTimeSeries().
                title("Duration").
                description("Time taken to process the requests.").
                unit("s").
                withTarget(new Dataquery.Builder().
                        queryType("randomWalk").
                        datasource(new Datasource.Builder().
                                type("grafana").
                                uid("grafana"))
                );
    }
}
