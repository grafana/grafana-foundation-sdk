---
title: <span class="badge builder"></span> Expression
---
# <span class="badge builder"></span> Expression

## Constructor

```python
Expression()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.Expression
```

### <span class="badge object-method"></span> datasource_uid

The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource

```python
def datasource_uid(datasource_uid: alertingv1beta1.DatasourceUID) -> typing.Self
```

### <span class="badge object-method"></span> model

```python
def model(model: object) -> typing.Self
```

### <span class="badge object-method"></span> query_type

The type of query if this is a query expression

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> relative_time_range

```python
def relative_time_range(relative_time_range: cogbuilder.Builder[alertingv1beta1.RelativeTimeRange]) -> typing.Self
```

### <span class="badge object-method"></span> source

Used to mark the expression to be used as the final source for the rule evaluation

Only one expression in a rule can be marked as the source

For AlertRules, this is the expression that will be evaluated against the alerting condition

For RecordingRules, this is the expression that will be recorded

```python
def source(source: bool) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Expression](./object-Expression.md)
