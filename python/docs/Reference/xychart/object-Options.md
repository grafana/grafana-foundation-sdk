---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    series_mapping: typing.Optional[xychart.SeriesMapping]
    # Table Mode (auto)
    dims: xychart.XYDimensionConfig
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # Manual Mode
    series: list[xychart.ScatterSeriesConfig]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

