---
title: <span class="badge object-type-class"></span> Dashboardv2RangeMapOptions
---
# <span class="badge object-type-class"></span> Dashboardv2RangeMapOptions

## Definition

```python
class Dashboardv2RangeMapOptions:
    # Min value of the range. It can be null which means -Infinity
    from_val: typing.Optional[float]
    # Max value of the range. It can be null which means +Infinity
    to: typing.Optional[float]
    # Config to apply when the value is within the range
    result: dashboardv2.ValueMappingResult
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

## See also

 * <span class="badge builder"></span> [Dashboardv2RangeMapOptions](./builder-Dashboardv2RangeMapOptions.md)
