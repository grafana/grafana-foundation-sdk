# Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type for the custom panel's options:

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
