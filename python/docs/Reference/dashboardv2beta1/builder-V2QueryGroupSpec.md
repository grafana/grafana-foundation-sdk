---
title: <span class="badge builder"></span> V2QueryGroupSpec
---
# <span class="badge builder"></span> V2QueryGroupSpec

## Constructor

```python
V2QueryGroupSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.V2QueryGroupSpec
```

### <span class="badge object-method"></span> queries

```python
def queries(queries: list[cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]]) -> typing.Self
```

### <span class="badge object-method"></span> query_options

```python
def query_options(query_options: cogbuilder.Builder[dashboardv2beta1.QueryOptionsSpec]) -> typing.Self
```

### <span class="badge object-method"></span> transformations

```python
def transformations(transformations: list[cogbuilder.Builder[dashboardv2beta1.V2TransformationKind]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [V2QueryGroupSpec](./object-V2QueryGroupSpec.md)
