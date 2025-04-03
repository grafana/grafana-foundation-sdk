---
title: <span class="badge builder"></span> AnnotationPanelFilterBuilder
---
# <span class="badge builder"></span> AnnotationPanelFilterBuilder

## Constructor

```go
func NewAnnotationPanelFilterBuilder() *AnnotationPanelFilterBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationPanelFilterBuilder) Build() (AnnotationPanelFilter, error)
```

### <span class="badge object-method"></span> Exclude

Should the specified panels be included or excluded

```go
func (builder *AnnotationPanelFilterBuilder) Exclude(exclude bool) *AnnotationPanelFilterBuilder
```

### <span class="badge object-method"></span> Ids

Panel IDs that should be included or excluded

```go
func (builder *AnnotationPanelFilterBuilder) Ids(ids []uint8) *AnnotationPanelFilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationPanelFilter](./object-AnnotationPanelFilter.md)
