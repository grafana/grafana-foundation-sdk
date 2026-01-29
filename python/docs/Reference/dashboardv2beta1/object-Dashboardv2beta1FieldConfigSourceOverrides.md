---
title: <span class="badge object-type-class"></span> Dashboardv2beta1FieldConfigSourceOverrides
---
# <span class="badge object-type-class"></span> Dashboardv2beta1FieldConfigSourceOverrides

## Definition

```python
class Dashboardv2beta1FieldConfigSourceOverrides:
    # Describes config override rules created when interacting with Grafana.
    system_ref: typing.Optional[str]
    matcher: dashboardv2beta1.MatcherConfig
    properties: list[dashboardv2beta1.DynamicConfigValue]
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

 * <span class="badge builder"></span> [Dashboardv2beta1FieldConfigSourceOverrides](./builder-Dashboardv2beta1FieldConfigSourceOverrides.md)
