---
title: <span class="badge object-type-class"></span> Rule
---
# <span class="badge object-type-class"></span> Rule

## Definition

```python
class Rule:
    annotations: typing.Optional[dict[str, str]]
    condition: str
    data: list[alerting.Query]
    exec_err_state: typing.Literal["OK", "Alerting", "Error"]
    folder_uid: str
    # The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    # Before this time has elapsed, the rule is only considered to be Pending.
    for_val: str
    id_val: typing.Optional[int]
    is_paused: typing.Optional[bool]
    labels: typing.Optional[dict[str, str]]
    no_data_state: typing.Literal["Alerting", "NoData", "OK"]
    org_id: int
    provenance: typing.Optional[alerting.Provenance]
    rule_group: str
    title: str
    uid: typing.Optional[str]
    updated: typing.Optional[str]
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

 * <span class="badge builder"></span> [Rule](./builder-Rule.md)
