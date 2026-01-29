---
title: <span class="badge builder"></span> QueryGroup
---
# <span class="badge builder"></span> QueryGroup

## Constructor

```python
QueryGroup()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.QueryGroupKind
```

### <span class="badge object-method"></span> query_options

```python
def query_options(query_options: cogbuilder.Builder[dashboardv2beta1.QueryOptionsSpec]) -> typing.Self
```

### <span class="badge object-method"></span> target

```python
def target(querie: cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]) -> typing.Self
```

### <span class="badge object-method"></span> targets

```python
def targets(queries: list[cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]]) -> typing.Self
```

### <span class="badge object-method"></span> transformation

```python
def transformation(transformation: cogbuilder.Builder[dashboardv2beta1.TransformationKind]) -> typing.Self
```

### <span class="badge object-method"></span> transformations

```python
def transformations(transformations: list[cogbuilder.Builder[dashboardv2beta1.TransformationKind]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [QueryGroupKind](./object-QueryGroupKind.md)
