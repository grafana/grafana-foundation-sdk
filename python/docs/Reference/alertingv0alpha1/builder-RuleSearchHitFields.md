---
title: <span class="badge builder"></span> RuleSearchHitFields
---
# <span class="badge builder"></span> RuleSearchHitFields

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Constructor

```python
RuleSearchHitFields()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.RuleSearchHitFields
```

### <span class="badge object-method"></span> annotations

Alert-rule fields.

```python
def annotations(annotations: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> dashboard_uid

```python
def dashboard_uid(dashboard_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> datasource_ui_ds

```python
def datasource_ui_ds(datasource_ui_ds: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> folder

```python
def folder(folder: str) -> typing.Self
```

### <span class="badge object-method"></span> for_val

```python
def for_val(for_val: str) -> typing.Self
```

### <span class="badge object-method"></span> interval

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> keep_firing_for

```python
def keep_firing_for(keep_firing_for: str) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> metric

Recording-rule fields.

```python
def metric(metric: str) -> typing.Self
```

### <span class="badge object-method"></span> notification_type

```python
def notification_type(notification_type: str) -> typing.Self
```

### <span class="badge object-method"></span> panel_id

```python
def panel_id(panel_id: int) -> typing.Self
```

### <span class="badge object-method"></span> paused

```python
def paused(paused: bool) -> typing.Self
```

### <span class="badge object-method"></span> receiver

```python
def receiver(receiver: str) -> typing.Self
```

### <span class="badge object-method"></span> routing_tree

```python
def routing_tree(routing_tree: str) -> typing.Self
```

### <span class="badge object-method"></span> target_datasource_uid

```python
def target_datasource_uid(target_datasource_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

```python
def type_val(type_val: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
