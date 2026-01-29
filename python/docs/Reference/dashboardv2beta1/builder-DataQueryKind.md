---
title: <span class="badge builder"></span> DataQueryKind
---
# <span class="badge builder"></span> DataQueryKind

## Constructor

```python
DataQueryKind()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.DataQueryKind
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> group

```python
def group(group: str) -> typing.Self
```

### <span class="badge object-method"></span> spec

```python
def spec(spec: object) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [DataQueryKind](./object-DataQueryKind.md)
