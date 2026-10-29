---
title: <span class="badge object-type-class"></span> Expression
---
# <span class="badge object-type-class"></span> Expression

## Definition

```python
class Expression:
    # The type of query if this is a query expression
    query_type: typing.Optional[str]
    relative_time_range: typing.Optional[alertingv0alpha1.RelativeTimeRange]
    # The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
    datasource_uid: typing.Optional[alertingv0alpha1.DatasourceUID]
    model: object
    # Used to mark the expression to be used as the final source for the rule evaluation
    # Only one expression in a rule can be marked as the source
    # For AlertRules, this is the expression that will be evaluated against the alerting condition
    # For RecordingRules, this is the expression that will be recorded
    source: typing.Optional[bool]
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

 * <span class="badge builder"></span> [Expression](./builder-Expression.md)
