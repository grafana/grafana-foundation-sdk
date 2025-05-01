---
title: <span class="badge builder"></span> AzureMetricDimensionBuilder
---
# <span class="badge builder"></span> AzureMetricDimensionBuilder

## Constructor

```java
new AzureMetricDimensionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureMetricDimension build()
```

### <span class="badge object-method"></span> dimension

Name of Dimension to be filtered on.

```java
public AzureMetricDimensionBuilder dimension(String dimension)
```

### <span class="badge object-method"></span> filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```java
public AzureMetricDimensionBuilder filter(String filter)
```

### <span class="badge object-method"></span> filters

Values to match with the filter.

```java
public AzureMetricDimensionBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```java
public AzureMetricDimensionBuilder operator(String operator)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)
