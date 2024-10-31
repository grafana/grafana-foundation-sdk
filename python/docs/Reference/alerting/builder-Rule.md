---
title: <span class="badge builder"></span> Rule
---
# <span class="badge builder"></span> Rule

## Constructor

```python
Rule(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.Rule
```

### <span class="badge object-method"></span> annotations

```python
def annotations(annotations: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> condition

```python
def condition(condition: str) -> typing.Self
```

### <span class="badge object-method"></span> exec_err_state

```python
def exec_err_state(exec_err_state: typing.Literal["OK", "Alerting", "Error"]) -> typing.Self
```

### <span class="badge object-method"></span> folder_uid

```python
def folder_uid(folder_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> for_val

The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.

Before this time has elapsed, the rule is only considered to be Pending.

```python
def for_val(for_val: str) -> typing.Self
```

### <span class="badge object-method"></span> id_val

```python
def id_val(id_val: int) -> typing.Self
```

### <span class="badge object-method"></span> is_paused

```python
def is_paused(is_paused: bool) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> no_data_state

```python
def no_data_state(no_data_state: typing.Literal["Alerting", "NoData", "OK"]) -> typing.Self
```

### <span class="badge object-method"></span> notification_settings

```python
def notification_settings(notification_settings: cogbuilder.Builder[alerting.NotificationSettings]) -> typing.Self
```

### <span class="badge object-method"></span> org_id

```python
def org_id(org_id: int) -> typing.Self
```

### <span class="badge object-method"></span> provenance

```python
def provenance(provenance: alerting.Provenance) -> typing.Self
```

### <span class="badge object-method"></span> queries

```python
def queries(data: list[cogbuilder.Builder[alerting.Query]]) -> typing.Self
```

### <span class="badge object-method"></span> record

```python
def record(record: cogbuilder.Builder[alerting.RecordRule]) -> typing.Self
```

### <span class="badge object-method"></span> rule_group

```python
def rule_group(rule_group: str) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> uid

```python
def uid(uid: str) -> typing.Self
```

### <span class="badge object-method"></span> updated

```python
def updated(updated: str) -> typing.Self
```

### <span class="badge object-method"></span> with_query

```python
def with_query(data: cogbuilder.Builder[alerting.Query]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Rule](./object-Rule.md)
