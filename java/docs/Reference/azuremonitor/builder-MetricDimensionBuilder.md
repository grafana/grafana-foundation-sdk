---
title: <span class="badge builder"></span> MetricDimensionBuilder
---
# <span class="badge builder"></span> MetricDimensionBuilder

## Constructor

```java
new MetricDimensionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MetricDimension build()
```

### <span class="badge object-method"></span> dimension

Name of Dimension to be filtered on.

```java
public MetricDimensionBuilder dimension(String dimension)
```

### <span class="badge object-method"></span> filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```java
public MetricDimensionBuilder filter(String filter)
```

### <span class="badge object-method"></span> filters

Values to match with the filter.

```java
public MetricDimensionBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```java
public MetricDimensionBuilder operator(String operator)
```

## See also

 * <span class="badge object-type-class"></span> [MetricDimension](./object-MetricDimension.md)
