---
title: <span class="badge builder"></span> Query
---
# <span class="badge builder"></span> Query

## Constructor

```python
Query()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.DataQueryKind
```

### <span class="badge object-method"></span> cloud_watch_annotation_query

```python
def cloud_watch_annotation_query(cloud_watch_annotation_query: cogbuilder.Builder[cloudwatch.CloudWatchAnnotationQuery]) -> typing.Self
```

### <span class="badge object-method"></span> cloud_watch_logs_query

```python
def cloud_watch_logs_query(cloud_watch_logs_query: cogbuilder.Builder[cloudwatch.CloudWatchLogsQuery]) -> typing.Self
```

### <span class="badge object-method"></span> cloud_watch_metrics_query

```python
def cloud_watch_metrics_query(cloud_watch_metrics_query: cogbuilder.Builder[cloudwatch.CloudWatchMetricsQuery]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
