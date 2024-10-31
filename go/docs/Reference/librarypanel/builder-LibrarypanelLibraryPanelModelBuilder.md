---
title: <span class="badge builder"></span> LibrarypanelLibraryPanelModelBuilder
---
# <span class="badge builder"></span> LibrarypanelLibraryPanelModelBuilder

## Constructor

```go
func NewLibrarypanelLibraryPanelModelBuilder() *LibrarypanelLibraryPanelModelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Build() (LibrarypanelLibraryPanelModel, error)
```

### <span class="badge object-method"></span> Datasource

The datasource used in all targets.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Datasource(datasource dashboard.DataSourceRef) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Description

Panel description.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Description(description string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> FieldConfig

Field options allow you to change how the data is displayed in your visualizations.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) FieldConfig(fieldConfig dashboard.FieldConfigSource) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Interval(interval string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Links

Panel links.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

The maximum number of data points that the panel queries are retrieving.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) MaxDataPoints(maxDataPoints float64) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Options

It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Options(options any) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> PluginVersion

The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) PluginVersion(pluginVersion string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Repeat

Name of template variable to repeat for.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Repeat(repeat string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> RepeatDirection

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) RepeatDirection(repeatDirection librarypanel.LibraryPanelRepeatDirection) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> RepeatPanelId

Id of the repeating panel.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) RepeatPanelId(repeatPanelId int64) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Tags

Tags for the panel.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Tags(tags []string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Targets

Depends on the panel plugin. See the plugin documentation for details.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Targets(targets []cog.Builder[cog/variants.Dataquery]) *LibrarypanelLibraryPanelModelBuilder
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
func (builder *LibrarypanelLibraryPanelModelBuilder) TimeFrom(timeFrom string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> TimeShift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) TimeShift(timeShift string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Title

Panel title.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Title(title string) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Transformations(transformations []dashboard.DataTransformerConfig) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Transparent

Whether to display the panel without a background.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Transparent(transparent bool) *LibrarypanelLibraryPanelModelBuilder
```

### <span class="badge object-method"></span> Type

The panel plugin type id. This is used to find the plugin to display the panel.

```go
func (builder *LibrarypanelLibraryPanelModelBuilder) Type(typeArg string) *LibrarypanelLibraryPanelModelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibrarypanelLibraryPanelModel](./object-LibrarypanelLibraryPanelModel.md)
