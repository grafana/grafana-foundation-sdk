---
title: <span class="badge builder"></span> NotebookLayoutItemSpec
---
# <span class="badge builder"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Constructor

```python
NotebookLayoutItemSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.NotebookLayoutItemSpec
```

### <span class="badge object-method"></span> collapsed

```python
def collapsed(collapsed: bool) -> typing.Self
```

### <span class="badge object-method"></span> element

```python
def element(element: cogbuilder.Builder[dashboardv2beta1.ElementReference]) -> typing.Self
```

### <span class="badge object-method"></span> source

```python
def source(source: typing.Literal["assistant", "user"]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [NotebookLayoutItemSpec](./object-NotebookLayoutItemSpec.md)
