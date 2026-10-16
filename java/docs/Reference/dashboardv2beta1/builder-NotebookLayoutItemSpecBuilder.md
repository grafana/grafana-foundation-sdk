---
title: <span class="badge builder"></span> NotebookLayoutItemSpecBuilder
---
# <span class="badge builder"></span> NotebookLayoutItemSpecBuilder

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Constructor

```java
new NotebookLayoutItemSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public NotebookLayoutItemSpec build()
```

### <span class="badge object-method"></span> collapsed

```java
public NotebookLayoutItemSpecBuilder collapsed(Boolean collapsed)
```

### <span class="badge object-method"></span> element

```java
public NotebookLayoutItemSpecBuilder element(com.grafana.foundation.cog.Builder<ElementReference> element)
```

### <span class="badge object-method"></span> source

```java
public NotebookLayoutItemSpecBuilder source(NotebookLayoutItemSpecSource source)
```

## See also

 * <span class="badge object-type-class"></span> [NotebookLayoutItemSpec](./object-NotebookLayoutItemSpec.md)
