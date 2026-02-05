---
title: <span class="badge builder"></span> VizConfigKindBuilder
---
# <span class="badge builder"></span> VizConfigKindBuilder

## Constructor

```java
new VizConfigKindBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public VizConfigKind build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```java
public VizConfigKindBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

```java
public VizConfigKindBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color)
```

### <span class="badge object-method"></span> custom

custom is specified by the FieldConfig field

in panel plugin schemas.

```java
public VizConfigKindBuilder custom(Object custom)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

```java
public VizConfigKindBuilder dataLinks(List<Object> links)
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```java
public VizConfigKindBuilder decimals(Double decimals)
```

### <span class="badge object-method"></span> defaults

Defaults are the options applied to all fields.

```java
public VizConfigKindBuilder defaults(FieldConfig defaults)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```java
public VizConfigKindBuilder description(String description)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```java
public VizConfigKindBuilder displayName(String displayName)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```java
public VizConfigKindBuilder displayNameFromDS(String displayNameFromDS)
```

### <span class="badge object-method"></span> fieldConfig

```java
public VizConfigKindBuilder fieldConfig(FieldConfigSource fieldConfig)
```

### <span class="badge object-method"></span> filterable

True if data source field supports ad-hoc filters

```java
public VizConfigKindBuilder filterable(Boolean filterable)
```

### <span class="badge object-method"></span> group

The group is the plugin ID

```java
public VizConfigKindBuilder group(String group)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```java
public VizConfigKindBuilder mappings(List<ValueMapping> mappings)
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VizConfigKindBuilder max(Double max)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VizConfigKindBuilder min(Double min)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```java
public VizConfigKindBuilder noValue(String noValue)
```

### <span class="badge object-method"></span> options

```java
public VizConfigKindBuilder options(Object options)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```java
public VizConfigKindBuilder override(MatcherConfig matcher, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

```java
public VizConfigKindBuilder overrideByFieldType(String fieldType, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

```java
public VizConfigKindBuilder overrideByName(String name, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByQuery

```java
public VizConfigKindBuilder overrideByQuery(String queryRefId, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```java
public VizConfigKindBuilder overrideByRegexp(String regexp, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```java
public VizConfigKindBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides)
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```java
public VizConfigKindBuilder path(String path)
```

### <span class="badge object-method"></span> spec

```java
public VizConfigKindBuilder spec(VizConfigSpec spec)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```java
public VizConfigKindBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds)
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

```java
public VizConfigKindBuilder unit(String unit)
```

### <span class="badge object-method"></span> version

```java
public VizConfigKindBuilder version(String version)
```

### <span class="badge object-method"></span> writeable

True if data source can write a value to the path. Auth/authz are supported separately

```java
public VizConfigKindBuilder writeable(Boolean writeable)
```

## See also

 * <span class="badge object-type-class"></span> [VizConfigKind](./object-VizConfigKind.md)
