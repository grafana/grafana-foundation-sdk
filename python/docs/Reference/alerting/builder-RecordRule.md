---
title: <span class="badge builder"></span> RecordRule
---
# <span class="badge builder"></span> RecordRule

## Constructor

```python
RecordRule()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alerting.RecordRule
```

### <span class="badge object-method"></span> from_val

Which expression node should be used as the input for the recorded metric.

```python
def from_val(from_val: str) -> typing.Self
```

### <span class="badge object-method"></span> metric

Name of the recorded metric.

```python
def metric(metric: str) -> typing.Self
```

### <span class="badge object-method"></span> target_datasource_uid

Which data source should be used to write the output of the recording rule, specified by UID.

```python
def target_datasource_uid(target_datasource_uid: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RecordRule](./object-RecordRule.md)
