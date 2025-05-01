---
title: <span class="badge builder"></span> ThresholdsConfigBuilder
---
# <span class="badge builder"></span> ThresholdsConfigBuilder

## Constructor

```java
new ThresholdsConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ThresholdsConfig build()
```

### <span class="badge object-method"></span> mode

Thresholds mode.

```java
public ThresholdsConfigBuilder mode(ThresholdsMode mode)
```

### <span class="badge object-method"></span> steps

Must be sorted by 'value', first value is always -Infinity

```java
public ThresholdsConfigBuilder steps(List<Threshold> steps)
```

## See also

 * <span class="badge object-type-class"></span> [ThresholdsConfig](./object-ThresholdsConfig.md)
