---
title: <span class="badge object-type-class"></span> LogGroup
---
# <span class="badge object-type-class"></span> LogGroup

## Definition

```python
class LogGroup:
    # ARN of the log group
    arn: str
    # Name of the log group
    name: str
    # AccountId of the log group
    account_id: typing.Optional[str]
    # Label of the log group
    account_label: typing.Optional[str]
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

 * <span class="badge builder"></span> [LogGroup](./builder-LogGroup.md)
