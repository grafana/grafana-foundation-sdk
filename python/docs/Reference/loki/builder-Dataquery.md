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
def build() -> loki.Dataquery
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
def editor_mode(editor_mode: loki.QueryEditorMode) -> typing.Self
```

### <span class="badge object-method"></span> expr

The LogQL query.

```python
def expr(expr: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> instant

@deprecated, now use queryType.

```python
def instant(instant: bool) -> typing.Self
```

### <span class="badge object-method"></span> legend_format

Used to override the name of the series.

```python
def legend_format(legend_format: str) -> typing.Self
```

### <span class="badge object-method"></span> max_lines

Used to limit the number of log rows returned.

```python
def max_lines(max_lines: int) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> range_val

@deprecated, now use queryType.

```python
def range_val(range_val: bool) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> resolution

@deprecated, now use step.

```python
def resolution(resolution: int) -> typing.Self
```

### <span class="badge object-method"></span> step

Used to set step value for range queries.

```python
def step(step: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
