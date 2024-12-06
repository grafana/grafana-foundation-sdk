---
title: <span class="badge object-type-class"></span> NotificationPolicy
---
# <span class="badge object-type-class"></span> NotificationPolicy

A Route is a node that contains definitions of how to handle alerts. This is modified

from the upstream alertmanager in that it adds the ObjectMatchers property.

## Definition

```python
class NotificationPolicy:
    """
    A Route is a node that contains definitions of how to handle alerts. This is modified
    from the upstream alertmanager in that it adds the ObjectMatchers property.
    """

    active_time_intervals: typing.Optional[list[str]]
    continue_val: typing.Optional[bool]
    group_by: typing.Optional[list[str]]
    group_interval: typing.Optional[str]
    group_wait: typing.Optional[str]
    # Deprecated. Remove before v1.0 release.
    match: typing.Optional[dict[str, str]]
    match_re: typing.Optional[alerting.MatchRegexps]
    # Matchers is a slice of Matchers that is sortable, implements Stringer, and
    # provides a Matches method to match a LabelSet against all Matchers in the
    # slice. Note that some users of Matchers might require it to be sorted.
    matchers: typing.Optional[alerting.Matchers]
    mute_time_intervals: typing.Optional[list[str]]
    object_matchers: typing.Optional[alerting.ObjectMatchers]
    provenance: typing.Optional[alerting.Provenance]
    receiver: typing.Optional[str]
    repeat_interval: typing.Optional[str]
    routes: typing.Optional[list[alerting.NotificationPolicy]]
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

 * <span class="badge builder"></span> [NotificationPolicy](./builder-NotificationPolicy.md)
