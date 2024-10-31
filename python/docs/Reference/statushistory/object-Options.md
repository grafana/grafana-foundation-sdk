---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Set the height of the rows
    row_height: float
    # Show values on the columns
    show_value: common.VisibilityMode
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    timezone: typing.Optional[list[common.TimeZone]]
    # Controls the column width
    col_width: typing.Optional[float]
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

