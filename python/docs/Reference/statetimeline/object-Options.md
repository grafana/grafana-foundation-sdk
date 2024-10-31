---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Show timeline values on chart
    show_value: common.VisibilityMode
    # Controls the row height
    row_height: float
    # Merge equal consecutive values
    merge_values: typing.Optional[bool]
    # Controls value alignment on the timelines
    align_value: typing.Optional[common.TimelineValueAlignment]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    timezone: typing.Optional[list[common.TimeZone]]
    # Enables pagination when > 0
    per_page: typing.Optional[float]
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

