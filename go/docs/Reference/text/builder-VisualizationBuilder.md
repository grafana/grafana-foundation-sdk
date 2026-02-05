---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```go
func NewVisualizationBuilder() *VisualizationBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VisualizationBuilder) Build() (dashboardv2beta1.VizConfigKind, error)
```

### <span class="badge object-method"></span> Actions

Define interactive HTTP requests that can be triggered from data visualizations.

```go
func (builder *VisualizationBuilder) Actions(actions []cog.Builder[dashboardv2beta1.Action]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Code

```go
func (builder *VisualizationBuilder) Code(code cog.Builder[text.CodeOptions]) *VisualizationBuilder
```

### <span class="badge object-method"></span> ColorScheme

Panel color configuration

```go
func (builder *VisualizationBuilder) ColorScheme(color cog.Builder[dashboardv2beta1.FieldColor]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Content

```go
func (builder *VisualizationBuilder) Content(content string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DataLinks

The behavior when clicking on a result

```go
func (builder *VisualizationBuilder) DataLinks(links []any) *VisualizationBuilder
```

### <span class="badge object-method"></span> Decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```go
func (builder *VisualizationBuilder) Decimals(decimals float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Description

Human readable field metadata

```go
func (builder *VisualizationBuilder) Description(description string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DisplayName

The display value for this field.  This supports template variables blank is auto

```go
func (builder *VisualizationBuilder) DisplayName(displayName string) *VisualizationBuilder
```

### <span class="badge object-method"></span> DisplayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```go
func (builder *VisualizationBuilder) DisplayNameFromDS(displayNameFromDS string) *VisualizationBuilder
```

### <span class="badge object-method"></span> Mappings

Convert input values into a display string

```go
func (builder *VisualizationBuilder) Mappings(mappings []dashboardv2beta1.ValueMapping) *VisualizationBuilder
```

### <span class="badge object-method"></span> Max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationBuilder) Max(max float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VisualizationBuilder) Min(min float64) *VisualizationBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *VisualizationBuilder) Mode(mode text.TextMode) *VisualizationBuilder
```

### <span class="badge object-method"></span> NoValue

Alternative to empty string

```go
func (builder *VisualizationBuilder) NoValue(noValue string) *VisualizationBuilder
```

### <span class="badge object-method"></span> Override

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationBuilder) Override(matcher dashboardv2beta1.MatcherConfig, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByFieldType

Adds override rules for all the fields of the given type.

```go
func (builder *VisualizationBuilder) OverrideByFieldType(fieldType string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByName

Adds override rules for a specific field, referred to by its name.

```go
func (builder *VisualizationBuilder) OverrideByName(name string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByQuery

```go
func (builder *VisualizationBuilder) OverrideByQuery(queryRefId string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> OverrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```go
func (builder *VisualizationBuilder) OverrideByRegexp(regexp string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder
```

### <span class="badge object-method"></span> Overrides

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VisualizationBuilder) Overrides(overrides []cog.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```go
func (builder *VisualizationBuilder) Path(path string) *VisualizationBuilder
```

### <span class="badge object-method"></span> Thresholds

Map numeric values to states

```go
func (builder *VisualizationBuilder) Thresholds(thresholds cog.Builder[dashboardv2beta1.ThresholdsConfig]) *VisualizationBuilder
```

### <span class="badge object-method"></span> Unit

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

```go
func (builder *VisualizationBuilder) Unit(unit string) *VisualizationBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
