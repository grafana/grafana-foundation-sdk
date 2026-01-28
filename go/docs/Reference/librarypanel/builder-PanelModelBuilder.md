---
title: <span class="badge builder"></span> PanelModelBuilder
---
# <span class="badge builder"></span> PanelModelBuilder

## Constructor

```go
func NewPanelModelBuilder() *PanelModelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PanelModelBuilder) Build() (PanelModel, error)
```

### <span class="badge object-method"></span> CacheTimeout

Sets panel queries cache timeout.

```go
func (builder *PanelModelBuilder) CacheTimeout(cacheTimeout string) *PanelModelBuilder
```

### <span class="badge object-method"></span> Datasource

The datasource used in all targets.

```go
func (builder *PanelModelBuilder) Datasource(datasource common.DataSourceRef) *PanelModelBuilder
```

### <span class="badge object-method"></span> Description

Panel description.

```go
func (builder *PanelModelBuilder) Description(description string) *PanelModelBuilder
```

### <span class="badge object-method"></span> FieldConfig

Field options allow you to change how the data is displayed in your visualizations.

```go
func (builder *PanelModelBuilder) FieldConfig(fieldConfig dashboard.FieldConfigSource) *PanelModelBuilder
```

### <span class="badge object-method"></span> HideTimeOverride

Controls if the timeFrom or timeShift overrides are shown in the panel header

```go
func (builder *PanelModelBuilder) HideTimeOverride(hideTimeOverride bool) *PanelModelBuilder
```

### <span class="badge object-method"></span> Interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelModelBuilder) Interval(interval string) *PanelModelBuilder
```

### <span class="badge object-method"></span> Links

Panel links.

```go
func (builder *PanelModelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *PanelModelBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

The maximum number of data points that the panel queries are retrieving.

```go
func (builder *PanelModelBuilder) MaxDataPoints(maxDataPoints float64) *PanelModelBuilder
```

### <span class="badge object-method"></span> MaxPerRow

Option for repeated panels that controls max items per row

Only relevant for horizontally repeated panels

```go
func (builder *PanelModelBuilder) MaxPerRow(maxPerRow float64) *PanelModelBuilder
```

### <span class="badge object-method"></span> Options

It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.

```go
func (builder *PanelModelBuilder) Options(options any) *PanelModelBuilder
```

### <span class="badge object-method"></span> PluginVersion

The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.

```go
func (builder *PanelModelBuilder) PluginVersion(pluginVersion string) *PanelModelBuilder
```

### <span class="badge object-method"></span> QueryCachingTTL

Overrides the data source configured time-to-live for a query cache item in milliseconds

```go
func (builder *PanelModelBuilder) QueryCachingTTL(queryCachingTTL float64) *PanelModelBuilder
```

### <span class="badge object-method"></span> Repeat

Name of template variable to repeat for.

```go
func (builder *PanelModelBuilder) Repeat(repeat string) *PanelModelBuilder
```

### <span class="badge object-method"></span> RepeatDirection

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```go
func (builder *PanelModelBuilder) RepeatDirection(repeatDirection librarypanel.PanelModelRepeatDirection) *PanelModelBuilder
```

### <span class="badge object-method"></span> Targets

Depends on the panel plugin. See the plugin documentation for details.

```go
func (builder *PanelModelBuilder) Targets(targets []cog.Builder[cog/variants.Dataquery]) *PanelModelBuilder
```

### <span class="badge object-method"></span> TimeFrom

Overrides the relative time range for individual panels,

which causes them to be different than what is selected in

the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different

time periods or days on the same dashboard.

The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),

`now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelModelBuilder) TimeFrom(timeFrom string) *PanelModelBuilder
```

### <span class="badge object-method"></span> TimeShift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *PanelModelBuilder) TimeShift(timeShift string) *PanelModelBuilder
```

### <span class="badge object-method"></span> Title

Panel title.

```go
func (builder *PanelModelBuilder) Title(title string) *PanelModelBuilder
```

### <span class="badge object-method"></span> Transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```go
func (builder *PanelModelBuilder) Transformations(transformations []dashboard.DataTransformerConfig) *PanelModelBuilder
```

### <span class="badge object-method"></span> Transparent

Whether to display the panel without a background.

```go
func (builder *PanelModelBuilder) Transparent(transparent bool) *PanelModelBuilder
```

### <span class="badge object-method"></span> Type

The panel plugin type id. This is used to find the plugin to display the panel.

```go
func (builder *PanelModelBuilder) Type(typeArg string) *PanelModelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PanelModel](./object-PanelModel.md)
