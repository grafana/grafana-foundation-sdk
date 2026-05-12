---
title: <span class="badge builder"></span> Dashboardv2RangeMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2RangeMapOptionsBuilder

## Constructor

```java
new Dashboardv2RangeMapOptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dashboardv2RangeMapOptions build()
```

### <span class="badge object-method"></span> from

Min value of the range. It can be null which means -Infinity

```java
public Dashboardv2RangeMapOptionsBuilder from(Double from)
```

### <span class="badge object-method"></span> result

Config to apply when the value is within the range

```java
public Dashboardv2RangeMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result)
```

### <span class="badge object-method"></span> to

Max value of the range. It can be null which means +Infinity

```java
public Dashboardv2RangeMapOptionsBuilder to(Double to)
```

## See also

 * <span class="badge object-type-class"></span> [Dashboardv2RangeMapOptions](./object-Dashboardv2RangeMapOptions.md)
