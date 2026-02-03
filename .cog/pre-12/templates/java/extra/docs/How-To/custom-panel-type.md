# Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type for the custom panel's options:

```java
package custompanel;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CustomPanelOptions {
    @JsonProperty("makeItBeautiful")
    private Boolean makeItBeautiful;

    public void setMakeItBeautiful(Boolean makeItBeautiful) {
        this.makeItBeautiful = makeItBeautiful;
    }
}
```

Then, create the builder for your new panel:

```java
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
```

Register the type with the SDK, and use it as usual to build a dashboard:

```java
package custompanel;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.dashboard.Dashboard;
import com.grafana.foundation.dashboard.DashboardBuilder;

public class Main {
    public static void main(String[] args) throws JsonProcessingException {
        Registry.registerPanel("custom-panel", CustomPanelOptions.class, null);

        Dashboard dashboard = new DashboardBuilder("[Example] Custom Panel").
                uid("example-custom-panel").
                withPanel(new CustomPanelBuilder().
                        title("Sample custom panel").
                        makeItBeautiful()
                ).
                build();

        System.out.println(dashboard.toJSON());
    }
}
```
