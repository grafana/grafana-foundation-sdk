---
title: <span class="badge object-type-class"></span> Dashboardv2beta1RegexMapOptions
---
# <span class="badge object-type-class"></span> Dashboardv2beta1RegexMapOptions

## Definition

```python
class Dashboardv2beta1RegexMapOptions:
    # Regular expression to match against
    pattern: str
    # Config to apply when the value matches the regex
    result: dashboardv2beta1.ValueMappingResult
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

 * <span class="badge builder"></span> [Dashboardv2beta1RegexMapOptions](./builder-Dashboardv2beta1RegexMapOptions.md)
