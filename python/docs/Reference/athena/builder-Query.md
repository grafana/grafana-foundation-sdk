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

### <span class="badge object-method"></span> column

```python
def column(column: str) -> typing.Self
```

### <span class="badge object-method"></span> connection_args

```python
def connection_args(connection_args: cogbuilder.Builder[athena.ConnectionArgs]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> format_val

```python
def format_val(format_val: athena.FormatOptions) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> old_datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def old_datasource(datasource: common.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> query_id

```python
def query_id(query_id: str) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
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

### <span class="badge object-method"></span> table

```python
def table(table: str) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
