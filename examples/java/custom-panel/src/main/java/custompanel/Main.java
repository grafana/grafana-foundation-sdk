package custompanel;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.dashboard.Dashboard;
import com.grafana.foundation.dashboard.DashboardBuilder;

public class Main {
    public static void main(String[] args) {
        Registry.registerPanel("custom-panel", CustomPanelOptions.class, null);

        Dashboard dashboard = new DashboardBuilder("[Example] Custom Panel").
                uid("example-custom-panel").
                withPanel(new CustomPanelBuilder().
                        title("Sample custom panel").
                        makeItBeautiful()
                ).
                build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
