package custompanel;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.dashboard.Panel;

public class CustomPanelBuilder implements Builder<Panel> {
    private final Panel internal;

    public CustomPanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "custom-panel";
    }

    public CustomPanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }

    public CustomPanelBuilder withTarget(Builder<Dataquery> target) {
        this.internal.targets.add(target.build());
        return this;
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

    @Override
    public Panel build() {
        return internal;
    }
}
