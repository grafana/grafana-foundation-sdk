---
title: <span class="badge builder"></span> PanelModelBuilder
---
# <span class="badge builder"></span> PanelModelBuilder

## Constructor

```typescript
new PanelModelBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasource

The datasource used in all targets.

```typescript
datasource(datasource: common.DataSourceRef)
```

### <span class="badge object-method"></span> description

Panel description.

```typescript
description(description: string)
```

### <span class="badge object-method"></span> fieldConfig

Field options allow you to change how the data is displayed in your visualizations.

```typescript
fieldConfig(fieldConfig: dashboard.FieldConfigSource)
```

### <span class="badge object-method"></span> hideTimeOverride

Controls if the timeFrom or timeShift overrides are shown in the panel header

```typescript
hideTimeOverride(hideTimeOverride: boolean)
```

### <span class="badge object-method"></span> interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```typescript
interval(interval: string)
```

### <span class="badge object-method"></span> links

Panel links.

```typescript
links(links: cog.Builder<dashboard.DashboardLink>[])
```

### <span class="badge object-method"></span> maxDataPoints

The maximum number of data points that the panel queries are retrieving.

```typescript
maxDataPoints(maxDataPoints: number)
```

### <span class="badge object-method"></span> maxPerRow

Option for repeated panels that controls max items per row

Only relevant for horizontally repeated panels

```typescript
maxPerRow(maxPerRow: number)
```

### <span class="badge object-method"></span> options

It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.

```typescript
options(options: any)
```

### <span class="badge object-method"></span> pluginVersion

The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.

```typescript
pluginVersion(pluginVersion: string)
```

### <span class="badge object-method"></span> repeat

Name of template variable to repeat for.

```typescript
repeat(repeat: string)
```

### <span class="badge object-method"></span> repeatDirection

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```typescript
repeatDirection(repeatDirection: "h" | "v")
```

### <span class="badge object-method"></span> tags

Tags for the panel.

```typescript
tags(tags: string[])
```

### <span class="badge object-method"></span> targets

Depends on the panel plugin. See the plugin documentation for details.

```typescript
targets(targets: cog.Builder<cog.Dataquery>[])
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

```typescript
timeFrom(timeFrom: string)
```

### <span class="badge object-method"></span> timeShift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```typescript
timeShift(timeShift: string)
```

### <span class="badge object-method"></span> title

Panel title.

```typescript
title(title: string)
```

### <span class="badge object-method"></span> transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```typescript
transformations(transformations: dashboard.DataTransformerConfig[])
```

### <span class="badge object-method"></span> transparent

Whether to display the panel without a background.

```typescript
transparent(transparent: boolean)
```

### <span class="badge object-method"></span> type

The panel plugin type id. This is used to find the plugin to display the panel.

```typescript
type(type: string)
```

## See also

 * <span class="badge object-type-interface"></span> [PanelModel](./object-PanelModel.md)
