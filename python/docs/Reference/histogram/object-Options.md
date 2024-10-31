---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    # Size of each bucket
    bucket_size: typing.Optional[int]
    # Offset buckets by this amount
    bucket_offset: typing.Optional[int]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # Combines multiple series into a single histogram
    combine: typing.Optional[bool]
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

