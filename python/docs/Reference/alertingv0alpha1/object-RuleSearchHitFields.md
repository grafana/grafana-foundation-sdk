---
title: <span class="badge object-type-class"></span> RuleSearchHitFields
---
# <span class="badge object-type-class"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Definition

```python
class RuleSearchHitFields:
    """
    #RuleSearchHitFields is the per-kind field payload returned on each hit.
    It carries the union of alert- and recording-rule search fields; only the
    fields relevant to a hit's kind are populated. This maps to the kind's
    declared searchFields.
    """

    title: typing.Optional[str]
    folder: typing.Optional[str]
    type_val: typing.Optional[str]
    interval: typing.Optional[str]
    paused: typing.Optional[bool]
    labels: typing.Optional[dict[str, str]]
    datasource_ui_ds: typing.Optional[list[str]]
    # Alert-rule fields.
    annotations: typing.Optional[dict[str, str]]
    for_val: typing.Optional[str]
    keep_firing_for: typing.Optional[str]
    dashboard_uid: typing.Optional[str]
    panel_id: typing.Optional[int]
    receiver: typing.Optional[str]
    notification_type: typing.Optional[str]
    routing_tree: typing.Optional[str]
    # Recording-rule fields.
    metric: typing.Optional[str]
    target_datasource_uid: typing.Optional[str]
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

 * <span class="badge builder"></span> [RuleSearchHitFields](./builder-RuleSearchHitFields.md)
