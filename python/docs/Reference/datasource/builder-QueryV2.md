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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> panel_id

Panel ID from wich the queries will be reused.

```python
def panel_id(panel_id: int) -> typing.Self
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

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

### <span class="badge object-method"></span> with_transforms

```python
def with_transforms(with_transforms: bool) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
