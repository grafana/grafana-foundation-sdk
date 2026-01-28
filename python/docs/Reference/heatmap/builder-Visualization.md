---
title: <span class="badge builder"></span> Visualization
---
# <span class="badge builder"></span> Visualization

## Constructor

```python
Visualization()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.VizConfigKind
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```python
def actions(actions: list[cogbuilder.Builder[dashboardv2beta1.Action]]) -> typing.Self
```

### <span class="badge object-method"></span> calculate

Controls if the heatmap should be calculated from data

```python
def calculate(calculate: bool) -> typing.Self
```

### <span class="badge object-method"></span> calculation

Calculation options for the heatmap

```python
def calculation(calculation: cogbuilder.Builder[common.HeatmapCalculationOptions]) -> typing.Self
```

### <span class="badge object-method"></span> cell_gap

Controls gap between cells

```python
def cell_gap(cell_gap: int) -> typing.Self
```

### <span class="badge object-method"></span> cell_radius

Controls cell radius

```python
def cell_radius(cell_radius: float) -> typing.Self
```

### <span class="badge object-method"></span> cell_values

Controls cell value unit

```python
def cell_values(cell_values: cogbuilder.Builder[heatmap.CellValues]) -> typing.Self
```

### <span class="badge object-method"></span> color

Controls the color options

```python
def color(color: cogbuilder.Builder[heatmap.HeatmapColorOptions]) -> typing.Self
```

### <span class="badge object-method"></span> color_scheme

Panel color configuration

```python
def color_scheme(color: cogbuilder.Builder[dashboardv2beta1.FieldColor]) -> typing.Self
```

### <span class="badge object-method"></span> data_links

The behavior when clicking on a result

```python
def data_links(links: list[object]) -> typing.Self
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```python
def decimals(decimals: float) -> typing.Self
```

### <span class="badge object-method"></span> description

Human readable field metadata

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> display_name

The display value for this field.  This supports template variables blank is auto

```python
def display_name(display_name: str) -> typing.Self
```

### <span class="badge object-method"></span> display_name_from_ds

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```python
def display_name_from_ds(display_name_from_ds: str) -> typing.Self
```

### <span class="badge object-method"></span> exemplars

Controls exemplar options

```python
def exemplars(exemplars: cogbuilder.Builder[heatmap.ExemplarConfig]) -> typing.Self
```

### <span class="badge object-method"></span> field_min_max

Calculate min max per field

```python
def field_min_max(field_min_max: bool) -> typing.Self
```

### <span class="badge object-method"></span> filter_values

Filters values between a given range

```python
def filter_values(filter_values: cogbuilder.Builder[heatmap.FilterValueRange]) -> typing.Self
```

### <span class="badge object-method"></span> hide_from

```python
def hide_from(hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self
```

### <span class="badge object-method"></span> legend

| *{

	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls legend options

```python
def legend(legend: cogbuilder.Builder[heatmap.HeatmapLegend]) -> typing.Self
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```python
def mappings(mappings: list[dashboardv2beta1.ValueMapping]) -> typing.Self
```

### <span class="badge object-method"></span> max_val

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```python
def max_val(max_val: float) -> typing.Self
```

### <span class="badge object-method"></span> min_val

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```python
def min_val(min_val: float) -> typing.Self
```

### <span class="badge object-method"></span> no_value

Alternative to empty string

```python
def no_value(no_value: str) -> typing.Self
```

### <span class="badge object-method"></span> null_value_mode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```python
def null_value_mode(null_value_mode: dashboardv2beta1.NullValueMode) -> typing.Self
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```python
def override(override: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) -> typing.Self
```

### <span class="badge object-method"></span> override_by_field_type

Adds override rules for all the fields of the given type.

```python
def override_by_field_type(field_type: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self
```

### <span class="badge object-method"></span> override_by_name

Adds override rules for a specific field, referred to by its name.

```python
def override_by_name(name: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self
```

### <span class="badge object-method"></span> override_by_query

```python
def override_by_query(query_ref_id: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self
```

### <span class="badge object-method"></span> override_by_regexp

Adds override rules for the fields whose name match the given regexp.

```python
def override_by_regexp(regexp: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```python
def overrides(overrides: list[cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]]) -> typing.Self
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```python
def path(path: str) -> typing.Self
```

### <span class="badge object-method"></span> rows_frame

Controls tick alignment and value name when not calculating from data

```python
def rows_frame(rows_frame: cogbuilder.Builder[heatmap.RowsHeatmapOptions]) -> typing.Self
```

### <span class="badge object-method"></span> scale_distribution

```python
def scale_distribution(scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self
```

### <span class="badge object-method"></span> selection_mode

Controls which axis to allow selection on

```python
def selection_mode(selection_mode: heatmap.HeatmapSelectionMode) -> typing.Self
```

### <span class="badge object-method"></span> show_value

| *{

	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed

}

Controls the display of the value in the cell

```python
def show_value(show_value: common.VisibilityMode) -> typing.Self
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```python
def thresholds(thresholds: cogbuilder.Builder[dashboardv2beta1.ThresholdsConfig]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

Controls tooltip options

```python
def tooltip(tooltip: cogbuilder.Builder[heatmap.HeatmapTooltip]) -> typing.Self
```

### <span class="badge object-method"></span> unit

Unit a field should use. The unit you select is applied to all fields except time.

You can use the units ID availables in Grafana or a custom unit.

Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts

As custom unit, you can use the following formats:

`suffix:<suffix>` for custom unit that should go after value.

`prefix:<prefix>` for custom unit that should go before value.

`time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.

`si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.

`count:<unit>` for a custom count unit.

`currency:<unit>` for custom a currency unit.

```python
def unit(unit: str) -> typing.Self
```

### <span class="badge object-method"></span> y_axis

Controls yAxis placement

```python
def y_axis(y_axis: cogbuilder.Builder[heatmap.YAxisConfig]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
