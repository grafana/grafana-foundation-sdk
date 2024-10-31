---
title: <span class="badge object-type-class"></span> AzureMonitorResource
---
# <span class="badge object-type-class"></span> AzureMonitorResource

## Definition

```python
class AzureMonitorResource:
    subscription: typing.Optional[str]
    resource_group: typing.Optional[str]
    resource_name: typing.Optional[str]
    metric_namespace: typing.Optional[str]
    region: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureMonitorResource](./builder-AzureMonitorResource.md)
