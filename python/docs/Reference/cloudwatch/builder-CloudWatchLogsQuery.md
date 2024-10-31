---
title: <span class="badge builder"></span> CloudWatchLogsQuery
---
# <span class="badge builder"></span> CloudWatchLogsQuery

## Constructor

```python
CloudWatchLogsQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> cloudwatch.CloudWatchLogsQuery
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> expression

The CloudWatch Logs Insights query to execute

```python
def expression(expression: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> id_val

```python
def id_val(id_val: str) -> typing.Self
```

### <span class="badge object-method"></span> log_group_names

@deprecated use logGroups

```python
def log_group_names(log_group_names: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> log_groups

Log groups to query

```python
def log_groups(log_groups: list[cogbuilder.Builder[cloudwatch.LogGroup]]) -> typing.Self
```

### <span class="badge object-method"></span> query_mode

Whether a query is a Metrics, Logs, or Annotations query

```python
def query_mode(query_mode: cloudwatch.CloudWatchQueryMode) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> region

AWS region to query for the logs

```python
def region(region: str) -> typing.Self
```

### <span class="badge object-method"></span> stats_groups

Fields to group the results by, this field is automatically populated whenever the query is updated

```python
def stats_groups(stats_groups: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CloudWatchLogsQuery](./object-CloudWatchLogsQuery.md)
