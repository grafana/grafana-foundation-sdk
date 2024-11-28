---
title: <span class="badge object-type-class"></span> AzureTracesFilter
---
# <span class="badge object-type-class"></span> AzureTracesFilter

## Definition

```python
class AzureTracesFilter:
    # Property name, auto-populated based on available traces.
    property_val: str
    # Comparison operator to use. Either equals or not equals.
    operation: str
    # Values to filter by.
    filters: list[str]
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

 * <span class="badge builder"></span> [AzureTracesFilter](./builder-AzureTracesFilter.md)