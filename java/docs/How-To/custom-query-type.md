# Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type for the custom query:

```java
package customquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grafana.foundation.cog.variants.Dataquery;

public class CustomQuery implements Dataquery {
    @JsonProperty("refId")
    public String refId;
    @JsonProperty("hide")
    public Boolean hide;
    @JsonProperty("query")
    public String query;

    @Override
    public String dataqueryName() {
        return "custom-query";
    }
}
```

Then, create a builder for it:

```java
package customquery;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;

public class CustomQueryBuilder implements Builder<Dataquery> {
    private final CustomQuery internal;

    public CustomQueryBuilder(String query) {
        this.internal = new CustomQuery();
        this.internal.query = query;
    }

    public CustomQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }

    public CustomQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }

    @Override
    public CustomQuery build() {
        return this.internal;
    }
}

```

Finally, register the type with the SDK and use it as usual to build a dashboard:

```java
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

```
