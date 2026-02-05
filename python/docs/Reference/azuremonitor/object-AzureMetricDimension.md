---
title: <span class="badge object-type-class"></span> AzureMetricDimension
---
# <span class="badge object-type-class"></span> AzureMetricDimension

## Definition

```python
class AzureMetricDimension:
    # Name of Dimension to be filtered on.
    dimension: typing.Optional[str]
    # String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    operator: typing.Optional[str]
    # Values to match with the filter.
    filters: typing.Optional[list[str]]
    # @deprecated filter is deprecated in favour of filters to support multiselect.
    filter_val: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureMetricDimension](./builder-AzureMetricDimension.md)
