---
title: <span class="badge builder"></span> FieldColorBuilder
---
# <span class="badge builder"></span> FieldColorBuilder

## Constructor

```java
new FieldColorBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public FieldColor build()
```

### <span class="badge object-method"></span> fixedColor

The fixed color value for fixed or shades color modes.

```java
public FieldColorBuilder fixedColor(String fixedColor)
```

### <span class="badge object-method"></span> mode

The main color scheme mode.

```java
public FieldColorBuilder mode(FieldColorModeId mode)
```

### <span class="badge object-method"></span> seriesBy

Some visualizations need to know how to assign a series color from by value color schemes.

```java
public FieldColorBuilder seriesBy(FieldColorSeriesByMode seriesBy)
```

## See also

 * <span class="badge object-type-class"></span> [FieldColor](./object-FieldColor.md)
