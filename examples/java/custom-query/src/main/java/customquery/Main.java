package customquery;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.common.DataSourceRef;
import com.grafana.foundation.dashboard.Dashboard;
import com.grafana.foundation.timeseries.PanelBuilder;
import customquery.CustomQuery.CustomQueryBuilder;

public class Main {
    public static void main(String[] args) {
        Registry.registerDataquery("custom-query", CustomQuery.class);

        Dashboard dashboard = new Dashboard.Builder("[Example] Custom query").
                uid("example-custom-query").
                withPanel(new PanelBuilder().
                    title("Sample panel").
                    withTarget(new CustomQueryBuilder("query-here").
                        datasource(
                            new DataSourceRef.Builder().
                                type("custom-query").
                                uid("grafana")
                        )
                    )
                ).
                build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
    }
}