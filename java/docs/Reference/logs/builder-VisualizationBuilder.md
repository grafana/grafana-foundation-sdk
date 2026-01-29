---
title: <span class="badge builder"></span> VisualizationBuilder
---
# <span class="badge builder"></span> VisualizationBuilder

## Constructor

```java
new VisualizationBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Visualization build()
```

### <span class="badge object-method"></span> actions

Define interactive HTTP requests that can be triggered from data visualizations.

```java
public VisualizationBuilder actions(List<com.grafana.foundation.cog.Builder<Action>> actions)
```

### <span class="badge object-method"></span> colorScheme

Panel color configuration

```java
public VisualizationBuilder colorScheme(com.grafana.foundation.cog.Builder<FieldColor> color)
```

### <span class="badge object-method"></span> controlsStorageKey

```java
public VisualizationBuilder controlsStorageKey(String controlsStorageKey)
```

### <span class="badge object-method"></span> dataLinks

The behavior when clicking on a result

```java
public VisualizationBuilder dataLinks(List<Object> links)
```

### <span class="badge object-method"></span> decimals

Specify the number of decimals Grafana includes in the rendered value.

If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.

For example 1.1234 will display as 1.12 and 100.456 will display as 100.

To display all decimals, set the unit to `String`.

```java
public VisualizationBuilder decimals(Double decimals)
```

### <span class="badge object-method"></span> dedupStrategy

```java
public VisualizationBuilder dedupStrategy(LogsDedupStrategy dedupStrategy)
```

### <span class="badge object-method"></span> description

Human readable field metadata

```java
public VisualizationBuilder description(String description)
```

### <span class="badge object-method"></span> displayName

The display value for this field.  This supports template variables blank is auto

```java
public VisualizationBuilder displayName(String displayName)
```

### <span class="badge object-method"></span> displayNameFromDS

This can be used by data sources that return and explicit naming structure for values and labels

When this property is configured, this value is used rather than the default naming strategy.

```java
public VisualizationBuilder displayNameFromDS(String displayNameFromDS)
```

### <span class="badge object-method"></span> displayedFields

```java
public VisualizationBuilder displayedFields(List<String> displayedFields)
```

### <span class="badge object-method"></span> enableInfiniteScrolling

```java
public VisualizationBuilder enableInfiniteScrolling(Boolean enableInfiniteScrolling)
```

### <span class="badge object-method"></span> enableLogDetails

```java
public VisualizationBuilder enableLogDetails(Boolean enableLogDetails)
```

### <span class="badge object-method"></span> fieldMinMax

Calculate min max per field

```java
public VisualizationBuilder fieldMinMax(Boolean fieldMinMax)
```

### <span class="badge object-method"></span> isFilterLabelActive

```java
public VisualizationBuilder isFilterLabelActive(Object isFilterLabelActive)
```

### <span class="badge object-method"></span> logRowMenuIconsAfter

```java
public VisualizationBuilder logRowMenuIconsAfter(Object logRowMenuIconsAfter)
```

### <span class="badge object-method"></span> logRowMenuIconsBefore

```java
public VisualizationBuilder logRowMenuIconsBefore(Object logRowMenuIconsBefore)
```

### <span class="badge object-method"></span> mappings

Convert input values into a display string

```java
public VisualizationBuilder mappings(List<ValueMapping> mappings)
```

### <span class="badge object-method"></span> max

The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationBuilder max(Double max)
```

### <span class="badge object-method"></span> min

The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.

```java
public VisualizationBuilder min(Double min)
```

### <span class="badge object-method"></span> noValue

Alternative to empty string

```java
public VisualizationBuilder noValue(String noValue)
```

### <span class="badge object-method"></span> nullValueMode

How null values should be handled when calculating field stats

"null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero

```java
public VisualizationBuilder nullValueMode(NullValueMode nullValueMode)
```

### <span class="badge object-method"></span> onClickFilterLabel

TODO: figure out how to define callbacks

```java
public VisualizationBuilder onClickFilterLabel(Object onClickFilterLabel)
```

### <span class="badge object-method"></span> onClickFilterOutLabel

```java
public VisualizationBuilder onClickFilterOutLabel(Object onClickFilterOutLabel)
```

### <span class="badge object-method"></span> onClickFilterOutString

```java
public VisualizationBuilder onClickFilterOutString(Object onClickFilterOutString)
```

### <span class="badge object-method"></span> onClickFilterString

```java
public VisualizationBuilder onClickFilterString(Object onClickFilterString)
```

### <span class="badge object-method"></span> onClickHideField

```java
public VisualizationBuilder onClickHideField(Object onClickHideField)
```

### <span class="badge object-method"></span> onClickShowField

```java
public VisualizationBuilder onClickShowField(Object onClickShowField)
```

### <span class="badge object-method"></span> onLogOptionsChange

```java
public VisualizationBuilder onLogOptionsChange(Object onLogOptionsChange)
```

### <span class="badge object-method"></span> onNewLogsReceived

```java
public VisualizationBuilder onNewLogsReceived(Object onNewLogsReceived)
```

### <span class="badge object-method"></span> override

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationBuilder override(String systemRef, MatcherConfig matcher, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByFieldType

Adds override rules for all the fields of the given type.

```java
public VisualizationBuilder overrideByFieldType(String fieldType, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByName

Adds override rules for a specific field, referred to by its name.

```java
public VisualizationBuilder overrideByName(String name, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByQuery

```java
public VisualizationBuilder overrideByQuery(String queryRefId, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrideByRegexp

Adds override rules for the fields whose name match the given regexp.

```java
public VisualizationBuilder overrideByRegexp(String regexp, List<DynamicConfigValue> properties)
```

### <span class="badge object-method"></span> overrides

Overrides are the options applied to specific fields overriding the defaults.

```java
public VisualizationBuilder overrides(List<com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides>> overrides)
```

### <span class="badge object-method"></span> path

An explicit path to the field in the datasource.  When the frame meta includes a path,

This will default to `${frame.meta.path}/${field.name}



When defined, this value can be used as an identifier within the datasource scope, and

may be used to update the results

```java
public VisualizationBuilder path(String path)
```

### <span class="badge object-method"></span> prettifyLogMessage

```java
public VisualizationBuilder prettifyLogMessage(Boolean prettifyLogMessage)
```

### <span class="badge object-method"></span> showCommonLabels

```java
public VisualizationBuilder showCommonLabels(Boolean showCommonLabels)
```

### <span class="badge object-method"></span> showControls

```java
public VisualizationBuilder showControls(Boolean showControls)
```

### <span class="badge object-method"></span> showLabels

```java
public VisualizationBuilder showLabels(Boolean showLabels)
```

### <span class="badge object-method"></span> showLogContextToggle

```java
public VisualizationBuilder showLogContextToggle(Boolean showLogContextToggle)
```

### <span class="badge object-method"></span> showTime

```java
public VisualizationBuilder showTime(Boolean showTime)
```

### <span class="badge object-method"></span> sortOrder

```java
public VisualizationBuilder sortOrder(LogsSortOrder sortOrder)
```

### <span class="badge object-method"></span> thresholds

Map numeric values to states

```java
public VisualizationBuilder thresholds(com.grafana.foundation.cog.Builder<ThresholdsConfig> thresholds)
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
public VisualizationBuilder unit(String unit)
```

### <span class="badge object-method"></span> wrapLogMessage

```java
public VisualizationBuilder wrapLogMessage(Boolean wrapLogMessage)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.VizConfigKind](../dashboardv2beta1/object-VizConfigKind.md)
