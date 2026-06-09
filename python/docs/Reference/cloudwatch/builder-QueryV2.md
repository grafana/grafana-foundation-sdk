---
title: <span class="badge builder"></span> QueryV2
---
# <span class="badge builder"></span> QueryV2

## Constructor

```python
QueryV2()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2.DataQueryKind
```

### <span class="badge object-method"></span> annotation_query

```python
def annotation_query(annotation_query: cogbuilder.Builder[cloudwatch.AnnotationQuery]) -> typing.Self
```

### <span class="badge object-method"></span> logs_query

```python
def logs_query(logs_query: cogbuilder.Builder[cloudwatch.LogsQuery]) -> typing.Self
```

### <span class="badge object-method"></span> metrics_query

```python
def metrics_query(metrics_query: cogbuilder.Builder[cloudwatch.MetricsQuery]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
