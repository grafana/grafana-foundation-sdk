---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```go
func NewAnnotationQueryBuilder() *AnnotationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationQueryBuilder) Build() (AnnotationQueryKind, error)
```

### <span class="badge object-method"></span> BuiltIn

```go
func (builder *AnnotationQueryBuilder) BuiltIn(builtIn bool) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Enable

```go
func (builder *AnnotationQueryBuilder) Enable(enable bool) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Filter

```go
func (builder *AnnotationQueryBuilder) Filter(filter cog.Builder[dashboardv2beta1.AnnotationPanelFilter]) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *AnnotationQueryBuilder) Hide(hide bool) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> IconColor

```go
func (builder *AnnotationQueryBuilder) IconColor(iconColor string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> LegacyOptions

Catch-all field for datasource-specific properties. Should not be available in as code tooling.

```go
func (builder *AnnotationQueryBuilder) LegacyOptions(legacyOptions map[string]any) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *AnnotationQueryBuilder) Name(name string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Placement

Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.

```go
func (builder *AnnotationQueryBuilder) Placement(placement string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *AnnotationQueryBuilder) Query(query cog.Builder[dashboardv2beta1.DataQueryKind]) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *AnnotationQueryBuilder) Spec(spec cog.Builder[dashboardv2beta1.AnnotationQuerySpec]) *AnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationQueryKind](./object-AnnotationQueryKind.md)
