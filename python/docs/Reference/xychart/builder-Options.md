---
title: <span class="badge builder"></span> Options
---
# <span class="badge builder"></span> Options

## Constructor

```python
Options()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> xychart.Options
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> mapping

```python
def mapping(mapping: xychart.SeriesMapping) -> typing.Self
```

### <span class="badge object-method"></span> series

```python
def series(series: list[cogbuilder.Builder[xychart.XYSeriesConfig]]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
