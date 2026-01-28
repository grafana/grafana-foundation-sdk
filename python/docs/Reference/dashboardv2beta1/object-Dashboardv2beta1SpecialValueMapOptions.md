---
title: <span class="badge object-type-class"></span> Dashboardv2beta1SpecialValueMapOptions
---
# <span class="badge object-type-class"></span> Dashboardv2beta1SpecialValueMapOptions

## Definition

```python
class Dashboardv2beta1SpecialValueMapOptions:
    # Special value to match against
    match: dashboardv2beta1.SpecialValueMatch
    # Config to apply when the value matches the special value
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

 * <span class="badge builder"></span> [Dashboardv2beta1SpecialValueMapOptions](./builder-Dashboardv2beta1SpecialValueMapOptions.md)
