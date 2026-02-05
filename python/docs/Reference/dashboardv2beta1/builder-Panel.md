---
title: <span class="badge builder"></span> Panel
---
# <span class="badge builder"></span> Panel

## Constructor

```python
Panel()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.PanelKind
```

### <span class="badge object-method"></span> data

```python
def data(data: cogbuilder.Builder[dashboardv2beta1.QueryGroupKind]) -> typing.Self
```

### <span class="badge object-method"></span> description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> id_val

```python
def id_val(id_val: float) -> typing.Self
```

### <span class="badge object-method"></span> links

```python
def links(links: list[cogbuilder.Builder[dashboardv2beta1.DataLink]]) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> transparent

```python
def transparent(transparent: bool) -> typing.Self
```

### <span class="badge object-method"></span> visualization

```python
def visualization(viz_config: cogbuilder.Builder[dashboardv2beta1.VizConfigKind]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [PanelKind](./object-PanelKind.md)
