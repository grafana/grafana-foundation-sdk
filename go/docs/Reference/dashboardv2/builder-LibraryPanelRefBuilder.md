---
title: <span class="badge builder"></span> LibraryPanelRefBuilder
---
# <span class="badge builder"></span> LibraryPanelRefBuilder

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Constructor

```go
func NewLibraryPanelRefBuilder() *LibraryPanelRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibraryPanelRefBuilder) Build() (LibraryPanelRef, error)
```

### <span class="badge object-method"></span> Name

Library panel name

```go
func (builder *LibraryPanelRefBuilder) Name(name string) *LibraryPanelRefBuilder
```

### <span class="badge object-method"></span> Uid

Library panel uid

```go
func (builder *LibraryPanelRefBuilder) Uid(uid string) *LibraryPanelRefBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibraryPanelRef](./object-LibraryPanelRef.md)
