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
def build() -> grafanapyroscope.Dataquery
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> group_by

Allows to group the results.

```python
def group_by(group_by: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> label_selector

Specifies the query label selectors.

```python
def label_selector(label_selector: str) -> typing.Self
```

### <span class="badge object-method"></span> max_nodes

Sets the maximum number of nodes in the flamegraph.

```python
def max_nodes(max_nodes: int) -> typing.Self
```

### <span class="badge object-method"></span> profile_type_id

Specifies the type of profile to query.

```python
def profile_type_id(profile_type_id: str) -> typing.Self
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

### <span class="badge object-method"></span> span_selector

Specifies the query span selectors.

```python
def span_selector(span_selector: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
