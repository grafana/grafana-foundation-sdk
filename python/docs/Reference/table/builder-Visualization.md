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

### <span class="badge object-method"></span> align

```python
def align(align: common.FieldTextAlignment) -> typing.Self
```

### <span class="badge object-method"></span> cell_height

Controls the height of the rows

```python
def cell_height(cell_height: common.TableCellHeight) -> typing.Self
```

### <span class="badge object-method"></span> cell_options

```python
def cell_options(cell_options: common.TableCellOptions) -> typing.Self
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

### <span class="badge object-method"></span> display_mode

This field is deprecated in favor of using cellOptions

```python
def display_mode(display_mode: common.TableCellDisplayMode) -> typing.Self
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

### <span class="badge object-method"></span> field_min_max

Calculate min max per field

```python
def field_min_max(field_min_max: bool) -> typing.Self
```

### <span class="badge object-method"></span> filterable

```python
def filterable(filterable: bool) -> typing.Self
```

### <span class="badge object-method"></span> footer

Controls footer options

```python
def footer(footer: cogbuilder.Builder[common.TableFooterOptions]) -> typing.Self
```

### <span class="badge object-method"></span> frame_index

Represents the index of the selected frame

```python
def frame_index(frame_index: float) -> typing.Self
```

### <span class="badge object-method"></span> hidden

?? default is missing or false ??

```python
def hidden(hidden: bool) -> typing.Self
```

### <span class="badge object-method"></span> hide_header

Hides any header for a column, useful for columns that show some static content or buttons.

```python
def hide_header(hide_header: bool) -> typing.Self
```

### <span class="badge object-method"></span> inspect

```python
def inspect(inspect: bool) -> typing.Self
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

### <span class="badge object-method"></span> min_width

```python
def min_width(min_width: float) -> typing.Self
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

### <span class="badge object-method"></span> show_header

Controls whether the panel should show the header

```python
def show_header(show_header: bool) -> typing.Self
```

### <span class="badge object-method"></span> show_type_icons

Controls whether the header should show icons for the column types

```python
def show_type_icons(show_type_icons: bool) -> typing.Self
```

### <span class="badge object-method"></span> sort_by

Used to control row sorting

```python
def sort_by(sort_by: list[cogbuilder.Builder[common.TableSortByFieldState]]) -> typing.Self
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```python
def thresholds(thresholds: cogbuilder.Builder[dashboardv2beta1.ThresholdsConfig]) -> typing.Self
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

### <span class="badge object-method"></span> width

```python
def width(width: float) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
