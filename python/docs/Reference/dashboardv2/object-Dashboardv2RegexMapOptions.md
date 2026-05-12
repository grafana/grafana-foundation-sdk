---
title: <span class="badge object-type-class"></span> Dashboardv2RegexMapOptions
---
# <span class="badge object-type-class"></span> Dashboardv2RegexMapOptions

## Definition

```python
class Dashboardv2RegexMapOptions:
    # Regular expression to match against
    pattern: str
    # Config to apply when the value matches the regex
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

 * <span class="badge builder"></span> [Dashboardv2RegexMapOptions](./builder-Dashboardv2RegexMapOptions.md)
