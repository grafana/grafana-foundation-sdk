---
title: <span class="badge builder"></span> VizConfigKindBuilder
---
# <span class="badge builder"></span> VizConfigKindBuilder

## Constructor

```go
func NewVizConfigKindBuilder() *VizConfigKindBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VizConfigKindBuilder) Build() (VizConfigKind, error)
```

### <span class="badge object-method"></span> Actions

Define interactive HTTP requests that can be triggered from data visualizations.

```go
func (builder *VizConfigKindBuilder) Actions(actions []cog.Builder[dashboardv2beta1.Action]) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> ColorScheme

Panel color configuration

```go
func (builder *VizConfigKindBuilder) ColorScheme(color cog.Builder[dashboardv2beta1.FieldColor]) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Custom

custom is specified by the FieldConfig field

in panel plugin schemas.

```go
func (builder *VizConfigKindBuilder) Custom(custom any) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> DataLinks

The behavior when clicking on a result

```go
func (builder *VizConfigKindBuilder) DataLinks(links []any) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```go
func (builder *VizConfigKindBuilder) Decimals(decimals float64) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Defaults

Defaults are the options applied to all fields.

```go
func (builder *VizConfigKindBuilder) Defaults(defaults dashboardv2beta1.FieldConfig) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Description

Human readable field metadata

```go
func (builder *VizConfigKindBuilder) Description(description string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> DisplayName

The display value for this field.  This supports template variables blank is auto

```go
func (builder *VizConfigKindBuilder) DisplayName(displayName string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> DisplayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```go
func (builder *VizConfigKindBuilder) DisplayNameFromDS(displayNameFromDS string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> FieldConfig

```go
func (builder *VizConfigKindBuilder) FieldConfig(fieldConfig dashboardv2beta1.FieldConfigSource) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> FieldMinMax

Calculate min max per field

```go
func (builder *VizConfigKindBuilder) FieldMinMax(fieldMinMax bool) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Filterable

True if data source field supports ad-hoc filters

```go
func (builder *VizConfigKindBuilder) Filterable(filterable bool) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Group

The group is the plugin ID

```go
func (builder *VizConfigKindBuilder) Group(group string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Mappings

Convert input values into a display string

```go
func (builder *VizConfigKindBuilder) Mappings(mappings []dashboardv2beta1.ValueMapping) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VizConfigKindBuilder) Max(max float64) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```go
func (builder *VizConfigKindBuilder) Min(min float64) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> NoValue

Alternative to empty string

```go
func (builder *VizConfigKindBuilder) NoValue(noValue string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> NullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```go
func (builder *VizConfigKindBuilder) NullValueMode(nullValueMode dashboardv2beta1.NullValueMode) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *VizConfigKindBuilder) Options(options any) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Override

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VizConfigKindBuilder) Override(systemRef string, matcher dashboardv2beta1.MatcherConfig, properties []dashboardv2beta1.DynamicConfigValue) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> OverrideByFieldType

Adds override rules for all the fields of the given type.

```go
func (builder *VizConfigKindBuilder) OverrideByFieldType(fieldType string, properties []dashboardv2beta1.DynamicConfigValue) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> OverrideByName

Adds override rules for a specific field, referred to by its name.

```go
func (builder *VizConfigKindBuilder) OverrideByName(name string, properties []dashboardv2beta1.DynamicConfigValue) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> OverrideByQuery

```go
func (builder *VizConfigKindBuilder) OverrideByQuery(queryRefId string, properties []dashboardv2beta1.DynamicConfigValue) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> OverrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```go
func (builder *VizConfigKindBuilder) OverrideByRegexp(regexp string, properties []dashboardv2beta1.DynamicConfigValue) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Overrides

Overrides are the options applied to specific fields overriding the defaults.

```go
func (builder *VizConfigKindBuilder) Overrides(overrides []cog.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```go
func (builder *VizConfigKindBuilder) Path(path string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *VizConfigKindBuilder) Spec(spec dashboardv2beta1.VizConfigSpec) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Thresholds

Map numeric values to states

```go
func (builder *VizConfigKindBuilder) Thresholds(thresholds cog.Builder[dashboardv2beta1.ThresholdsConfig]) *VizConfigKindBuilder
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
func (builder *VizConfigKindBuilder) Unit(unit string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *VizConfigKindBuilder) Version(version string) *VizConfigKindBuilder
```

### <span class="badge object-method"></span> Writeable

True if data source can write a value to the path. Auth/authz are supported separately

```go
func (builder *VizConfigKindBuilder) Writeable(writeable bool) *VizConfigKindBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VizConfigKind](./object-VizConfigKind.md)
