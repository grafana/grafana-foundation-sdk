---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```java
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Options build()
```

### <span class="badge object-method"></span> bucketCount

Bucket count (approx)

```java
public OptionsBuilder bucketCount(Integer bucketCount)
```

### <span class="badge object-method"></span> bucketOffset

Offset buckets by this amount

```java
public OptionsBuilder bucketOffset(Float bucketOffset)
```

### <span class="badge object-method"></span> bucketSize

Size of each bucket

```java
public OptionsBuilder bucketSize(Integer bucketSize)
```

### <span class="badge object-method"></span> combine

Combines multiple series into a single histogram

```java
public OptionsBuilder combine(Boolean combine)
```

### <span class="badge object-method"></span> legend

```java
public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend)
```

### <span class="badge object-method"></span> tooltip

```java
public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
