---
title: <span class="badge builder"></span> V2PanelSpec
---
# <span class="badge builder"></span> V2PanelSpec

## Constructor

```python
V2PanelSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.V2PanelSpec
```

### <span class="badge object-method"></span> data

```python
def data(data: cogbuilder.Builder[dashboardv2beta1.V2QueryGroupKind]) -> typing.Self
```

### <span class="badge object-method"></span> description

Shown in a info icon tooltip next to panel title

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

### <span class="badge object-method"></span> subtitle

Shown in a sub header below the title.

```python
def subtitle(subtitle: str) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> transparent

```python
def transparent(transparent: bool) -> typing.Self
```

### <span class="badge object-method"></span> viz_config

```python
def viz_config(viz_config: cogbuilder.Builder[dashboardv2beta1.VizConfigKind]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [V2PanelSpec](./object-V2PanelSpec.md)
