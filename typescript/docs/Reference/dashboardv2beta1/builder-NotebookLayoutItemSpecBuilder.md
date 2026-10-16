---
title: <span class="badge builder"></span> NotebookLayoutItemSpecBuilder
---
# <span class="badge builder"></span> NotebookLayoutItemSpecBuilder

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Constructor

```typescript
new NotebookLayoutItemSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> collapsed

```typescript
collapsed(collapsed: boolean)
```

### <span class="badge object-method"></span> element

```typescript
element(element: cog.Builder<dashboardv2beta1.ElementReference>)
```

### <span class="badge object-method"></span> source

```typescript
source(source: "assistant" | "user")
```

## See also

 * <span class="badge object-type-interface"></span> [NotebookLayoutItemSpec](./object-NotebookLayoutItemSpec.md)
