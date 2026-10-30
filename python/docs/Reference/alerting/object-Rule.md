---
title: <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Rule
---
# <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Rule

<span class="badge deprecated"></span>Prefer using the alertingv1beta1 package instead

## Definition

```python
class Rule:
    warnings.warn("Prefer using the alertingv1beta1 package instead", DeprecationWarning)
    annotations: typing.Optional[dict[str, str]]
    condition: str
    data: list[alerting.Query]
    exec_err_state: typing.Literal["Alerting", "Error", "OK", "KeepLast"]
    folder_uid: str
    # The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    # Before this time has elapsed, the rule is only considered to be Pending.
    for_val: str
    id_val: typing.Optional[int]
    is_paused: typing.Optional[bool]
    keep_firing_for: typing.Optional[str]
    labels: typing.Optional[dict[str, str]]
    missing_series_evals_to_resolve: typing.Optional[int]
    no_data_state: typing.Literal["OK", "Alerting", "NoData", "KeepLast"]
    notification_settings: typing.Optional[alerting.NotificationSettings]
    org_id: int
    provenance: typing.Optional[alerting.Provenance]
    record: typing.Optional[alerting.RecordRule]
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

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [Rule](./builder-Rule.md)
