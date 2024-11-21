---
title: <span class="badge object-type-class"></span> PanelModel
---
# <span class="badge object-type-class"></span> PanelModel

Dashboard panels are the basic visualization building blocks.

## Definition

```python
class PanelModel:
    """
    Dashboard panels are the basic visualization building blocks.
    """

    # The panel plugin type id. This is used to find the plugin to display the panel.
    type_val: str
    # The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    plugin_version: typing.Optional[str]
    # Tags for the panel.
    tags: typing.Optional[list[str]]
    # Depends on the panel plugin. See the plugin documentation for details.
    targets: typing.Optional[list[cogvariants.Dataquery]]
    # Panel title.
    title: typing.Optional[str]
    # Panel description.
    description: typing.Optional[str]
    # Whether to display the panel without a background.
    transparent: typing.Optional[bool]
    # The datasource used in all targets.
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Panel links.
    links: typing.Optional[list[dashboard.DashboardLink]]
    # Name of template variable to repeat for.
    repeat: typing.Optional[str]
    # Direction to repeat in if 'repeat' is set.
    # `h` for horizontal, `v` for vertical.
    repeat_direction: typing.Optional[typing.Literal["h", "v"]]
    # Option for repeated panels that controls max items per row
    # Only relevant for horizontally repeated panels
    max_per_row: typing.Optional[float]
    # The maximum number of data points that the panel queries are retrieving.
    max_data_points: typing.Optional[float]
    # List of transformations that are applied to the panel data before rendering.
    # When there are multiple transformations, Grafana applies them in the order they are listed.
    # Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    transformations: typing.Optional[list[dashboard.DataTransformerConfig]]
    # The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    # This value must be formatted as a number followed by a valid time
    # identifier like: "40s", "3d", etc.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    interval: typing.Optional[str]
    # Overrides the relative time range for individual panels,
    # which causes them to be different than what is selected in
    # the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    # time periods or days on the same dashboard.
    # The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    # `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    # Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    time_from: typing.Optional[str]
    # Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    # For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    # Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    time_shift: typing.Optional[str]
    # Controls if the timeFrom or timeShift overrides are shown in the panel header
    hide_time_override: typing.Optional[bool]
    # It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    options: typing.Optional[object]
    # Field options allow you to change how the data is displayed in your visualizations.
    field_config: typing.Optional[dashboard.FieldConfigSource]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [PanelModel](./builder-PanelModel.md)
