---
title: <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQuery
---
# <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQuery

## Constructor

```python
LegacyCloudMonitoringAnnotationQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> googlecloudmonitoring.LegacyCloudMonitoringAnnotationQuery
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```python
def filters(filters: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> metric_kind

```python
def metric_kind(metric_kind: googlecloudmonitoring.MetricKind) -> typing.Self
```

### <span class="badge object-method"></span> metric_type

```python
def metric_type(metric_type: str) -> typing.Self
```

### <span class="badge object-method"></span> project_name

GCP project to execute the query against.

```python
def project_name(project_name: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

Query refId.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> text

Annotation text.

```python
def text(text: str) -> typing.Self
```

### <span class="badge object-method"></span> title

Annotation title.

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> value_type

```python
def value_type(value_type: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [LegacyCloudMonitoringAnnotationQuery](./object-LegacyCloudMonitoringAnnotationQuery.md)
