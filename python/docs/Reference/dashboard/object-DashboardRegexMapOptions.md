---
title: <span class="badge object-type-class"></span> DashboardRegexMapOptions
---
# <span class="badge object-type-class"></span> DashboardRegexMapOptions

## Definition

```python
class DashboardRegexMapOptions:
    # Regular expression to match against
    pattern: str
    # Config to apply when the value matches the regex
    result: dashboard.ValueMappingResult
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

 * <span class="badge builder"></span> [DashboardRegexMapOptions](./builder-DashboardRegexMapOptions.md)
