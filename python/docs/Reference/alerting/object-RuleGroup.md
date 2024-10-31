---
title: <span class="badge object-type-class"></span> RuleGroup
---
# <span class="badge object-type-class"></span> RuleGroup

## Definition

```python
class RuleGroup:
    folder_uid: typing.Optional[str]
    # The interval, in seconds, at which all rules in the group are evaluated.
    # If a group contains many rules, the rules are evaluated sequentially.
    interval: typing.Optional[alerting.Duration]
    rules: typing.Optional[list[alerting.Rule]]
    title: typing.Optional[str]
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

 * <span class="badge builder"></span> [RuleGroup](./builder-RuleGroup.md)
