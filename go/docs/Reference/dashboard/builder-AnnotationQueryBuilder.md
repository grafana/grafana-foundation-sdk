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
func (builder *AnnotationQueryBuilder) Build() (AnnotationQuery, error)
```

### <span class="badge object-method"></span> BuiltIn

Set to 1 for the standard annotation query all dashboards have by default.

```go
func (builder *AnnotationQueryBuilder) BuiltIn(builtIn float64) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Datasource

Datasource where the annotations data is

```go
func (builder *AnnotationQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Enable

When enabled the annotation query is issued with every dashboard refresh

```go
func (builder *AnnotationQueryBuilder) Enable(enable bool) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Expr

```go
func (builder *AnnotationQueryBuilder) Expr(expr string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Filter

Filters to apply when fetching annotations

```go
func (builder *AnnotationQueryBuilder) Filter(filter cog.Builder[dashboard.AnnotationPanelFilter]) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Hide

Annotation queries can be toggled on or off at the top of the dashboard.

When hide is true, the toggle is not shown in the dashboard.

```go
func (builder *AnnotationQueryBuilder) Hide(hide bool) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> IconColor

Color to use for the annotation event markers

```go
func (builder *AnnotationQueryBuilder) IconColor(iconColor string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Name

Name of annotation.

```go
func (builder *AnnotationQueryBuilder) Name(name string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Target

TODO.. this should just be a normal query target

```go
func (builder *AnnotationQueryBuilder) Target(target cog.Builder[dashboard.AnnotationTarget]) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Type

TODO -- this should not exist here, it is based on the --grafana-- datasource

```go
func (builder *AnnotationQueryBuilder) Type(typeArg string) *AnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationQuery](./object-AnnotationQuery.md)
