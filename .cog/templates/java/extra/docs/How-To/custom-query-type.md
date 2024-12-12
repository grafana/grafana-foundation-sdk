# Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type for the custom query:

```java
public class CustomQuery implements Dataquery {
    @JsonProperty("refId")
    private String refId;
    @JsonProperty("hide")
    private Boolean hide;
    @JsonProperty("query")
    private String query;

    public static class Builder implements com.grafana.foundation.cog.Builder<Dataquery> {
        private final CustomQuery internal;

        public Builder(String query) {
            internal = new CustomQuery();
            this.internal.query = query;
        }

        public Builder refId(String refId) {
            this.internal.refId = refId;
            return this;
        }

        public Builder hide(Boolean hide) {
            this.internal.hide = hide;
            return this;
        }

        @Override
        public CustomQuery build() {
            return this.internal;
        }
    }
}
```

Register the type with cog, and use it as usual to build a dashboard:

```java
public class Main {
    public static void main(String[] args) {
        Registry.registerDataquery("custom", CustomQuery.class);

        Dashboard dashboard = new Dashboard.Builder("Custom query type").
                uid("test-custom-query-type").
                refresh("1m").
                time(new DashboardDashboardTime.Builder().from("now-30m").to("now")).
                withRow(new RowPanel.Builder("Overview")).
                withPanel(new PanelBuilder().
                        title("Sample panel").
                        withTarget(
                                new CustomQuery.Builder("query here")
                        )
                ).
                build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
```
