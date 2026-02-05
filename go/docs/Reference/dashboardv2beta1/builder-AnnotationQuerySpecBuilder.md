---
title: <span class="badge builder"></span> AnnotationQuerySpecBuilder
---
# <span class="badge builder"></span> AnnotationQuerySpecBuilder

## Constructor

```go
func NewAnnotationQuerySpecBuilder() *AnnotationQuerySpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationQuerySpecBuilder) Build() (AnnotationQuerySpec, error)
```

### <span class="badge object-method"></span> BuiltIn

```go
func (builder *AnnotationQuerySpecBuilder) BuiltIn(builtIn bool) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Enable

```go
func (builder *AnnotationQuerySpecBuilder) Enable(enable bool) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Filter

```go
func (builder *AnnotationQuerySpecBuilder) Filter(filter cog.Builder[dashboardv2beta1.AnnotationPanelFilter]) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *AnnotationQuerySpecBuilder) Hide(hide bool) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> IconColor

```go
func (builder *AnnotationQuerySpecBuilder) IconColor(iconColor string) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> LegacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```go
func (builder *AnnotationQuerySpecBuilder) LegacyOptions(legacyOptions map[string]any) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *AnnotationQuerySpecBuilder) Name(name string) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```go
func (builder *AnnotationQuerySpecBuilder) Placement(placement string) *AnnotationQuerySpecBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *AnnotationQuerySpecBuilder) Query(query cog.Builder[dashboardv2beta1.DataQueryKind]) *AnnotationQuerySpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationQuerySpec](./object-AnnotationQuerySpec.md)
