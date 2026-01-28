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
def build() -> statetimeline.Options
```

### <span class="badge object-method"></span> align_value

Controls value alignment on the timelines

```python
def align_value(align_value: common.TimelineValueAlignment) -> typing.Self
```

### <span class="badge object-method"></span> legend

```python
def legend(legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self
```

### <span class="badge object-method"></span> merge_values

Merge equal consecutive values

```python
def merge_values(merge_values: bool) -> typing.Self
```

### <span class="badge object-method"></span> row_height

Controls the row height

```python
def row_height(row_height: float) -> typing.Self
```

### <span class="badge object-method"></span> show_value

Show timeline values on chart

```python
def show_value(show_value: common.VisibilityMode) -> typing.Self
```

### <span class="badge object-method"></span> timezone

```python
def timezone(timezone: list[common.TimeZone]) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

```python
def tooltip(tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Options](./object-Options.md)
