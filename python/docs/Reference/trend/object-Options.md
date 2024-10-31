---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

Identical to timeseries... except it does not have timezone settings

## Definition

```python
class Options:
    """
    Identical to timeseries... except it does not have timezone settings
    """

    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # Name of the x field to use (defaults to first number)
    x_field: typing.Optional[str]
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

