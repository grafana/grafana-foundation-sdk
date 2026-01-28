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
def build() -> piechart.Options
```

### <span class="badge object-method"></span> display_labels

```python
def display_labels(display_labels: list[piechart.PieChartLabels]) -> typing.Self
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[piechart.PieChartLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> orientation

```python
def orientation(orientation: common.VizOrientation) -> typing.Self
```

### <span class="badge object-method"></span> pie_type

```python
def pie_type(pie_type: piechart.PieChartType) -> typing.Self
```

### <span class="badge object-method"></span> reduce_options

```python
def reduce_options(reduce_options: cogbuilder.Builder[common.ReduceDataOptions]) -> typing.Self
```

### <span class="badge object-method"></span> text

```python
def text(text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
