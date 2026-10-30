---
title: <span class="badge object-type-class"></span> AlertRuleInstanceTotals
---
# <span class="badge object-type-class"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Definition

```python
class AlertRuleInstanceTotals:
    """
    Count of alert instances per state. error also counts instances whose evaluation
    errored but were mapped to another state via execErrState, so it can overlap.
    """

    healthy: typing.Optional[int]
    firing: typing.Optional[int]
    pending: typing.Optional[int]
    recovering: typing.Optional[int]
    nodata: typing.Optional[int]
    error: typing.Optional[int]
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

 * <span class="badge builder"></span> [AlertRuleInstanceTotals](./builder-AlertRuleInstanceTotals.md)
