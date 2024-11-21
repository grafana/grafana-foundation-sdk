---
title: <span class="badge builder"></span> PanelModel
---
# <span class="badge builder"></span> PanelModel

## Constructor

```python
PanelModel()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> librarypanel.PanelModel
```

### <span class="badge object-method"></span> datasource

The datasource used in all targets.

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> description

Panel description.

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> field_config

Field options allow you to change how the data is displayed in your visualizations.

```python
def field_config(field_config: dashboard.FieldConfigSource) -> typing.Self
```

### <span class="badge object-method"></span> hide_time_override

Controls if the timeFrom or timeShift overrides are shown in the panel header

```python
def hide_time_override(hide_time_override: bool) -> typing.Self
```

### <span class="badge object-method"></span> interval

The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.

This value must be formatted as a number followed by a valid time

identifier like: "40s", "3d", etc.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> links

Panel links.

```python
def links(links: list[cogbuilder.Builder[dashboard.DashboardLink]]) -> typing.Self
```

### <span class="badge object-method"></span> max_data_points

The maximum number of data points that the panel queries are retrieving.

```python
def max_data_points(max_data_points: float) -> typing.Self
```

### <span class="badge object-method"></span> max_per_row

Option for repeated panels that controls max items per row

Only relevant for horizontally repeated panels

```python
def max_per_row(max_per_row: float) -> typing.Self
```

### <span class="badge object-method"></span> options

It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.

```python
def options(options: object) -> typing.Self
```

### <span class="badge object-method"></span> plugin_version

The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.

```python
def plugin_version(plugin_version: str) -> typing.Self
```

### <span class="badge object-method"></span> repeat

Name of template variable to repeat for.

```python
def repeat(repeat: str) -> typing.Self
```

### <span class="badge object-method"></span> repeat_direction

Direction to repeat in if 'repeat' is set.

`h` for horizontal, `v` for vertical.

```python
def repeat_direction(repeat_direction: typing.Literal["h", "v"]) -> typing.Self
```

### <span class="badge object-method"></span> tags

Tags for the panel.

```python
def tags(tags: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> targets

Depends on the panel plugin. See the plugin documentation for details.

```python
def targets(targets: list[cogbuilder.Builder[cogvariants.Dataquery]]) -> typing.Self
```

### <span class="badge object-method"></span> time_from

Overrides the relative time range for individual panels,

which causes them to be different than what is selected in

the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different

time periods or days on the same dashboard.

The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),

`now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```python
def time_from(time_from: str) -> typing.Self
```

### <span class="badge object-method"></span> time_shift

Overrides the time range for individual panels by shifting its start and end relative to the time picker.

For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.

Note: Panel time overrides have no effect when the dashboard’s time range is absolute.

See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options

```python
def time_shift(time_shift: str) -> typing.Self
```

### <span class="badge object-method"></span> title

Panel title.

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> transformations

List of transformations that are applied to the panel data before rendering.

When there are multiple transformations, Grafana applies them in the order they are listed.

Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.

```python
def transformations(transformations: list[dashboard.DataTransformerConfig]) -> typing.Self
```

### <span class="badge object-method"></span> transparent

Whether to display the panel without a background.

```python
def transparent(transparent: bool) -> typing.Self
```

### <span class="badge object-method"></span> type_val

The panel plugin type id. This is used to find the plugin to display the panel.

```python
def type_val(type_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [PanelModel](./object-PanelModel.md)
