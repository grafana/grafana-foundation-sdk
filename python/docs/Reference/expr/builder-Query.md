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

### <span class="badge object-method"></span> type_classic_conditions

```python
def type_classic_conditions(type_classic_conditions: cogbuilder.Builder[expr.TypeClassicConditions]) -> typing.Self
```

### <span class="badge object-method"></span> type_math

```python
def type_math(type_math: cogbuilder.Builder[expr.TypeMath]) -> typing.Self
```

### <span class="badge object-method"></span> type_reduce

```python
def type_reduce(type_reduce: cogbuilder.Builder[expr.TypeReduce]) -> typing.Self
```

### <span class="badge object-method"></span> type_resample

```python
def type_resample(type_resample: cogbuilder.Builder[expr.TypeResample]) -> typing.Self
```

### <span class="badge object-method"></span> type_sql

```python
def type_sql(type_sql: cogbuilder.Builder[expr.TypeSql]) -> typing.Self
```

### <span class="badge object-method"></span> type_threshold

```python
def type_threshold(type_threshold: cogbuilder.Builder[expr.TypeThreshold]) -> typing.Self
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
