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

### <span class="badge object-method"></span> classic_conditions

```python
def classic_conditions(type_classic_conditions: cogbuilder.Builder[expr.TypeClassicConditions]) -> typing.Self
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

### <span class="badge object-method"></span> math

```python
def math(type_math: cogbuilder.Builder[expr.TypeMath]) -> typing.Self
```

### <span class="badge object-method"></span> reduce

```python
def reduce(type_reduce: cogbuilder.Builder[expr.TypeReduce]) -> typing.Self
```

### <span class="badge object-method"></span> resample

```python
def resample(type_resample: cogbuilder.Builder[expr.TypeResample]) -> typing.Self
```

### <span class="badge object-method"></span> sql

```python
def sql(type_sql: cogbuilder.Builder[expr.TypeSql]) -> typing.Self
```

### <span class="badge object-method"></span> threshold

```python
def threshold(type_threshold: cogbuilder.Builder[expr.TypeThreshold]) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
