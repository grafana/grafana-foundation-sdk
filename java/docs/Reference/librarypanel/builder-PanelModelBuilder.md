---
title: <span class="badge builder"></span> PanelModelBuilder
---
# <span class="badge builder"></span> PanelModelBuilder

## Constructor

```java
new PanelModelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public PanelModel build()
```

### <span class="badge object-method"></span> datasource

The datasource used in all targets.

```java
public PanelModelBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> description

Panel description.

```java
public PanelModelBuilder description(String description)
```

### <span class="badge object-method"></span> fieldConfig

Field options allow you to change how the data is displayed in your visualizations.

```java
public PanelModelBuilder fieldConfig(FieldConfigSource fieldConfig)
```

### <span class="badge object-method"></span> interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```java
public PanelModelBuilder interval(String interval)
```

### <span class="badge object-method"></span> links

Panel links.

```java
public PanelModelBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links)
```

### <span class="badge object-method"></span> maxDataPoints

The maximum number of data points that the panel queries are retrieving.

```java
public PanelModelBuilder maxDataPoints(Double maxDataPoints)
```

### <span class="badge object-method"></span> options

It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.

```java
public PanelModelBuilder options(Object options)
```

### <span class="badge object-method"></span> pluginVersion

The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.

```java
public PanelModelBuilder pluginVersion(String pluginVersion)
```

### <span class="badge object-method"></span> repeat

Name of template variable to repeat for.

```java
public PanelModelBuilder repeat(String repeat)
```

### <span class="badge object-method"></span> repeatDirection

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```java
public PanelModelBuilder repeatDirection(PanelModelRepeatDirection repeatDirection)
```

### <span class="badge object-method"></span> repeatPanelId

Id of the repeating panel.

```java
public PanelModelBuilder repeatPanelId(Long repeatPanelId)
```

### <span class="badge object-method"></span> tags

Tags for the panel.

```java
public PanelModelBuilder tags(List<String> tags)
```

### <span class="badge object-method"></span> targets

Depends on the panel plugin. See the plugin documentation for details.

```java
public PanelModelBuilder targets(List<com.grafana.foundation.cog.Builder<Dataquery>> targets)
```

### <span class="badge object-method"></span> timeFrom

Overrides the relative time range for individual panels,

which causes them to be different than what is selected in

the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different

time periods or days on the same dashboard.

The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),

`now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```java
public PanelModelBuilder timeFrom(String timeFrom)
```

### <span class="badge object-method"></span> timeShift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```java
public PanelModelBuilder timeShift(String timeShift)
```

### <span class="badge object-method"></span> title

Panel title.

```java
public PanelModelBuilder title(String title)
```

### <span class="badge object-method"></span> transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```java
public PanelModelBuilder transformations(List<DataTransformerConfig> transformations)
```

### <span class="badge object-method"></span> transparent

Whether to display the panel without a background.

```java
public PanelModelBuilder transparent(Boolean transparent)
```

### <span class="badge object-method"></span> type

The panel plugin type id. This is used to find the plugin to display the panel.

```java
public PanelModelBuilder type(String type)
```

## See also

 * <span class="badge object-type-class"></span> [PanelModel](./object-PanelModel.md)
