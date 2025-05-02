package customquery;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.dashboard.Dashboard;
import com.grafana.foundation.dashboard.DashboardBuilder;
import com.grafana.foundation.timeseries.PanelBuilder;

public class Main {
    public static void main(String[] args) {
        Registry.registerDataquery("custom-query", CustomQuery.class);

        Dashboard dashboard = new DashboardBuilder("[Example] Custom query").
                uid("example-custom-query").
                withPanel(new PanelBuilder().
                    title("Sample panel").
                    withTarget(new customquery.CustomQueryBuilder("query-here"))
                ).
                build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
    }
}
