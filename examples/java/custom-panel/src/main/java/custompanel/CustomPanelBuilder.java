package custompanel;

import com.grafana.foundation.dashboard.Panel;

public class CustomPanelBuilder extends Panel.Builder<CustomPanelBuilder> {

    public CustomPanelBuilder() {
        super();
        this.internal.type = "custompanel";
    }

    public CustomPanelBuilder makeItBeautiful() {
        if (this.internal.options == null) {
            this.internal.options = new CustomPanelOptions();
        }

        CustomPanelOptions options = (CustomPanelOptions) this.internal.options;
        options.setMakeItBeautiful(true);

        this.internal.options = options;
        return this;
    }
}
