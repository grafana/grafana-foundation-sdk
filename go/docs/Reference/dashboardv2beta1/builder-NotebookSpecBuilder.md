---
title: <span class="badge builder"></span> NotebookSpecBuilder
---
# <span class="badge builder"></span> NotebookSpecBuilder

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Constructor

```go
func NewNotebookSpecBuilder() *NotebookSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NotebookSpecBuilder) Build() (NotebookSpec, error)
```

### <span class="badge object-method"></span> Description

```go
func (builder *NotebookSpecBuilder) Description(description string) *NotebookSpecBuilder
```

### <span class="badge object-method"></span> Elements

```go
func (builder *NotebookSpecBuilder) Elements(elements map[string]dashboardv2beta1.NotebookElement) *NotebookSpecBuilder
```

### <span class="badge object-method"></span> Layout

```go
func (builder *NotebookSpecBuilder) Layout(layout cog.Builder[dashboardv2beta1.NotebookLayoutKind]) *NotebookSpecBuilder
```

### <span class="badge object-method"></span> Tags

```go
func (builder *NotebookSpecBuilder) Tags(tags []string) *NotebookSpecBuilder
```

### <span class="badge object-method"></span> TimeSettings

```go
func (builder *NotebookSpecBuilder) TimeSettings(timeSettings cog.Builder[dashboardv2beta1.TimeSettingsSpec]) *NotebookSpecBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *NotebookSpecBuilder) Title(title string) *NotebookSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NotebookSpec](./object-NotebookSpec.md)
