---
title: <span class="badge builder"></span> Dataquery
---
# <span class="badge builder"></span> Dataquery

## Constructor

```python
Dataquery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> bigquery.Dataquery
```

### <span class="badge object-method"></span> convert_to_utc

```python
def convert_to_utc(convert_to_utc: bool) -> typing.Self
```

### <span class="badge object-method"></span> dataset

```python
def dataset(dataset: str) -> typing.Self
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> editor_mode

```python
def editor_mode(editor_mode: bigquery.EditorMode) -> typing.Self
```

### <span class="badge object-method"></span> format_val

```python
def format_val(format_val: bigquery.QueryFormat) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> location

```python
def location(location: str) -> typing.Self
```

### <span class="badge object-method"></span> partitioned

```python
def partitioned(partitioned: bool) -> typing.Self
```

### <span class="badge object-method"></span> partitioned_field

```python
def partitioned_field(partitioned_field: str) -> typing.Self
```

### <span class="badge object-method"></span> project

```python
def project(project: str) -> typing.Self
```

### <span class="badge object-method"></span> query_priority

```python
def query_priority(query_priority: bigquery.QueryPriority) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> raw_query

```python
def raw_query(raw_query: bool) -> typing.Self
```

### <span class="badge object-method"></span> raw_sql

```python
def raw_sql(raw_sql: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> sharded

```python
def sharded(sharded: bool) -> typing.Self
```

### <span class="badge object-method"></span> sql

```python
def sql(sql: cogbuilder.Builder[bigquery.SQLExpression]) -> typing.Self
```

### <span class="badge object-method"></span> table

```python
def table(table: str) -> typing.Self
```

### <span class="badge object-method"></span> time_shift

```python
def time_shift(time_shift: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
