package red;

import com.grafana.foundation.testdata.DataqueryBuilder;
import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.units.Constants;
import com.grafana.foundation.dashboard.*;

import java.util.List;

import static com.grafana.foundation.common.Constants.TimeZoneBrowser;
import static red.Common.defaultTimeSeries;

public class Red {
    static DashboardBuilder red(String dashboardTitle, List<String> serviceIDs) {
        DashboardBuilder builder = new DashboardBuilder("[Example] " + dashboardTitle).
                uid("example-red-method").
                tags(List.of("generated", "red")).
                editable().
                tooltip(DashboardCursorSync.CROSSHAIR).
                refresh("30s").
                time(new DashboardDashboardTimeBuilder().
                        from("now-30m").
                        to("now")
                ).
                timezone(TimeZoneBrowser).
                timepicker(new TimePickerBuilder().
                        refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"))
                ).
                link(new DashboardLinkBuilder("RED method").
                        type(DashboardLinkType.LINK).icon("question").
                        targetBlank(true).
                        url("https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method")
                );

        for (String serviceId : serviceIDs) {
            builder.withRow(new RowBuilder(serviceId)).
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
                unit(Constants.RequestsPerSecond).
                withTarget(new DataqueryBuilder().
                        queryType("randomWalk").
                        datasource(new DataSourceRef("grafana", "grafana"))
                );
    }

    private static PanelBuilder errorRateTimeSeries() {
        return defaultTimeSeries().
                title("Error rate").
                description("Number of failed requests, per second.").
                unit(Constants.RequestsPerSecond).
                withTarget(new DataqueryBuilder().
                        queryType("randomWalk").
                        datasource(new DataSourceRef("grafana", "grafana"))
                );
    }

    private static PanelBuilder durationRateTimeSeries() {
        return defaultTimeSeries().
                title("Duration").
                description("Time taken to process the requests.").
                unit(Constants.Seconds).
                withTarget(new DataqueryBuilder().
                        queryType("randomWalk").
                        datasource(new DataSourceRef("grafana", "grafana"))
                );
    }
}
