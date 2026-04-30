---
title: <span class="badge builder"></span> VisualizationV2Builder
---
# <span class="badge builder"></span> VisualizationV2Builder

## Constructor

```go
func NewVisualizationV2Builder() *VisualizationV2Builder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VisualizationV2Builder) Build() (dashboardv2.VizConfigKind, error)
```

### <span class="badge object-method"></span> Actions

Define interactive HTTP requests that can be triggered from data visualizations.

```go
func (builder *VisualizationV2Builder) Actions(actions []cog.Builder[dashboardv2.Action]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Align

```go
func (builder *VisualizationV2Builder) Align(align common.FieldTextAlignment) *VisualizationV2Builder
```

### <span class="badge object-method"></span> CellHeight

Controls the height of the rows

```go
func (builder *VisualizationV2Builder) CellHeight(cellHeight common.TableCellHeight) *VisualizationV2Builder
```

### <span class="badge object-method"></span> CellOptions

```go
func (builder *VisualizationV2Builder) CellOptions(cellOptions common.TableCellOptions) *VisualizationV2Builder
```

### <span class="badge object-method"></span> ColorScheme

Panel color configuration

```go
func (builder *VisualizationV2Builder) ColorScheme(color cog.Builder[dashboardv2.FieldColor]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> DataLinks

The behavior when clicking on a result

```go
func (builder *VisualizationV2Builder) DataLinks(links []any) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```go
func (builder *VisualizationV2Builder) Decimals(decimals float64) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Description

Human readable field metadata

```go
func (builder *VisualizationV2Builder) Description(description string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> DisplayMode

This field is deprecated in favor of using cellOptions

```go
func (builder *VisualizationV2Builder) DisplayMode(displayMode common.TableCellDisplayMode) *VisualizationV2Builder
```

### <span class="badge object-method"></span> DisplayName

The display value for this field.  This supports template variables blank is auto

```go
func (builder *VisualizationV2Builder) DisplayName(displayName string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> DisplayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```go
func (builder *VisualizationV2Builder) DisplayNameFromDS(displayNameFromDS string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> FieldMinMax

Calculate min max per field

```go
func (builder *VisualizationV2Builder) FieldMinMax(fieldMinMax bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Filterable

```go
func (builder *VisualizationV2Builder) Filterable(filterable bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Footer

Controls footer options

```go
func (builder *VisualizationV2Builder) Footer(footer cog.Builder[common.TableFooterOptions]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> FrameIndex

Represents the index of the selected frame

```go
func (builder *VisualizationV2Builder) FrameIndex(frameIndex float64) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Hidden

?? default is missing or false ??

```go
func (builder *VisualizationV2Builder) Hidden(hidden bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> HideHeader

Hides any header for a column, useful for columns that show some static content or buttons.

```go
func (builder *VisualizationV2Builder) HideHeader(hideHeader bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Inspect

```go
func (builder *VisualizationV2Builder) Inspect(inspect bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Mappings

Convert input values into a display string

```go
func (builder *VisualizationV2Builder) Mappings(mappings []dashboardv2.ValueMapping) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationV2Builder) Max(max float64) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationV2Builder) Min(min float64) *VisualizationV2Builder
```

### <span class="badge object-method"></span> MinWidth

```go
func (builder *VisualizationV2Builder) MinWidth(minWidth float64) *VisualizationV2Builder
```

### <span class="badge object-method"></span> NoValue

Alternative to empty string

```go
func (builder *VisualizationV2Builder) NoValue(noValue string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> NullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```go
func (builder *VisualizationV2Builder) NullValueMode(nullValueMode dashboardv2.NullValueMode) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Override

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationV2Builder) Override(matcher dashboardv2.MatcherConfig, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder
```

### <span class="badge object-method"></span> OverrideByFieldType

Adds override rules for all the fields of the given type.

```go
func (builder *VisualizationV2Builder) OverrideByFieldType(fieldType string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder
```

### <span class="badge object-method"></span> OverrideByName

Adds override rules for a specific field, referred to by its name.

```go
func (builder *VisualizationV2Builder) OverrideByName(name string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder
```

### <span class="badge object-method"></span> OverrideByQuery

```go
func (builder *VisualizationV2Builder) OverrideByQuery(queryRefId string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder
```

### <span class="badge object-method"></span> OverrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```go
func (builder *VisualizationV2Builder) OverrideByRegexp(regexp string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Overrides

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationV2Builder) Overrides(overrides []cog.Builder[dashboardv2.Dashboardv2FieldConfigSourceOverrides]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```go
func (builder *VisualizationV2Builder) Path(path string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> ShowHeader

Controls whether the panel should show the header

```go
func (builder *VisualizationV2Builder) ShowHeader(showHeader bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> ShowTypeIcons

Controls whether the header should show icons for the column types

```go
func (builder *VisualizationV2Builder) ShowTypeIcons(showTypeIcons bool) *VisualizationV2Builder
```

### <span class="badge object-method"></span> SortBy

Used to control row sorting

```go
func (builder *VisualizationV2Builder) SortBy(sortBy []cog.Builder[common.TableSortByFieldState]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Thresholds

Map numeric values to states

```go
func (builder *VisualizationV2Builder) Thresholds(thresholds cog.Builder[dashboardv2.ThresholdsConfig]) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Unit

Unit a field should use. The unit you select is applied to all fields except time.

You can use the units ID available in Grafana or a custom unit.

Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts

As custom unit, you can use the following formats:

`suffix:<suffix>` for custom unit that should go after value.

`prefix:<prefix>` for custom unit that should go before value.

`time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.

`si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.

`count:<unit>` for a custom count unit.

`currency:<unit>` for custom a currency unit.

```go
func (builder *VisualizationV2Builder) Unit(unit string) *VisualizationV2Builder
```

### <span class="badge object-method"></span> Width

```go
func (builder *VisualizationV2Builder) Width(width float64) *VisualizationV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.VizConfigKind](../dashboardv2/object-VizConfigKind.md)
