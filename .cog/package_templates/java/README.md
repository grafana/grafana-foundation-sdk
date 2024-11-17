# Grafana Foundation SDK – Java

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in Java.

> [!NOTE]
> This branch contains **types and builders generated for Grafana {{ .Extra.GrafanaVersion }}.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Installing

### Gradle
```kotlin
implementation("com.grafana:grafana-foundation-sdk:{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}")
```

### Maven
```xml
<dependency>
    <groupId>com.grafana</groupId>
    <artifactId>grafana-foundation-sdk</artifactId>
    <version>{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}</version>
</dependency>
```

## Example usage

### Building a dashboard

```java
public class Main {
    public static void main(String[] args) {
        Dashboard dashboard = new Dashboard.Builder("Sample Dashboard").
                Uid("generated-from-java").
                Tags(List.of("generated", "from", "java")).
                Refresh("1m").Time(new DashboardDashboardTime.Builder().
                        From("now-30m").
                        To("now")
                ).
                Timezone(Constants.TimeZoneBrowser).
                WithRow(new RowPanel.Builder("Overview")).
                WithPanel(new PanelBuilder().
                        Title("Network Received").
                        Unit("bps").
                        Min(0.0).
                        WithTarget(new Dataquery.Builder().
                                Expr("rate(node_network_receive_bytes_total{job=\"integrations/raspberrypi-node\", device!=\"lo\"}[$__rate_interval]) * 8").
                                LegendFormat({{ `"{{ device }}"` }})
                        )
                ).build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
    }
}
```

### Unmarshaling a dashboard

```java
public class Main {
    public static void main(String[] args) {
        ObjectMapper mapper = new ObjectMapper();
        try {
            InputStream json = Main.class.getResourceAsStream("/dashboard.json");
            Dashboard dashboard = mapper.readValue(json, Dashboard.class);
            System.out.println(dashboard.toJSON());
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

### Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type and a builder for the custom query:

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

### Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type and a builder for the custom panel's options:

```java
public class CustomOptions {
    @JsonProperty("makeItBeautiful")
    private Boolean makeItBeautiful;

    public void setMakeItBeautiful(Boolean makeItBeautiful) {
        this.makeItBeautiful = makeItBeautiful;
    }
}
```

Optionally you can add FieldConfig:

```java
public class CustomFieldConfig {
    @JsonProperty("sayHello")
    private String sayHello;

    public void setSayHello(String sayHello) {
        this.sayHello = sayHello;
    }
}
```

After that you need to create the panel's builder for your new panel:

```java
public class PanelBuilder implements Builder<Panel> {
    private final Panel internal;

    public PanelBuilder() {
        this.internal = new Panel();
        this.internal.type = "custom-panel";
    }

    public PanelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }

    public PanelBuilder withTarget(Builder<Dataquery> target) {
        this.internal.targets.add(target.build());
        return this;
    }

    public PanelBuilder makeItBeautiful() {
        if (this.internal.options == null) {
            this.internal.options = new CustomOptions();
        }

        CustomOptions options = (CustomOptions) this.internal.options;
        options.setMakeItBeautiful(true);

        this.internal.options = options;
        return this;
    }

    public PanelBuilder sayHello() {
        if (this.internal.fieldConfig == null) {
            this.internal.fieldConfig = new FieldConfigSource();
        }

        if (this.internal.fieldConfig.defaults == null) {
            this.internal.fieldConfig.defaults = new FieldConfig();
        }

        if (this.internal.fieldConfig.defaults.custom == null) {
            this.internal.fieldConfig.defaults.custom = new CustomFieldConfig();
        }

        CustomFieldConfig customFieldConfig = (CustomFieldConfig) this.internal.fieldConfig.defaults.custom;
        customFieldConfig.setSayHello("hello!");

        this.internal.fieldConfig.defaults.custom = customFieldConfig;
        return this;
    }

    @Override
    public Panel build() {
        return internal;
    }
}
```

Register the type with cog, and use it as usual to build a dashboard:

```java
public class Main {
    public static void main(String[] args) {
        Registry.registerPanel("custom-panel", CustomOptions.class, CustomFieldConfig.class);

        Dashboard dashboard = new Dashboard.Builder("Custom panel type").
                uid("test-custom-panel").
                refresh("1m").
                time(new DashboardDashboardTime.Builder().from("now-30m").to("now")).
                withRow(new RowPanel.Builder("Overview")).
                withPanel(new PanelBuilder().
                        title("Sample panel").
                        makeItBeautiful().
                        sayHello()
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

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
