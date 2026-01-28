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

### <span class="badge object-method"></span> cellHeight

Controls the height of the rows

```java
public OptionsBuilder cellHeight(TableCellHeight cellHeight)
```

### <span class="badge object-method"></span> footer

Controls footer options

```java
public OptionsBuilder footer(com.grafana.foundation.cog.Builder<TableFooterOptions> footer)
```

### <span class="badge object-method"></span> frameIndex

Represents the index of the selected frame

```java
public OptionsBuilder frameIndex(Double frameIndex)
```

### <span class="badge object-method"></span> showHeader

Controls whether the panel should show the header

```java
public OptionsBuilder showHeader(Boolean showHeader)
```

### <span class="badge object-method"></span> showTypeIcons

Controls whether the header should show icons for the column types

```java
public OptionsBuilder showTypeIcons(Boolean showTypeIcons)
```

### <span class="badge object-method"></span> sortBy

Used to control row sorting

```java
public OptionsBuilder sortBy(List<com.grafana.foundation.cog.Builder<TableSortByFieldState>> sortBy)
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
