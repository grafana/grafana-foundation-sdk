---
title: <span class="badge object-type-class"></span> SimulationQuery
---
# <span class="badge object-type-class"></span> SimulationQuery

## Definition

```python
class SimulationQuery:
    config: typing.Optional[object]
    key: testdata.Key
    last: typing.Optional[bool]
    stream: typing.Optional[bool]
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

 * <span class="badge builder"></span> [SimulationQuery](./builder-SimulationQuery.md)
