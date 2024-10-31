---
title: <span class="badge builder"></span> AnnotationTargetBuilder
---
# <span class="badge builder"></span> AnnotationTargetBuilder

## Constructor

```go
func NewAnnotationTargetBuilder() *AnnotationTargetBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationTargetBuilder) Build() (AnnotationTarget, error)
```

### <span class="badge object-method"></span> Limit

Only required/valid for the grafana datasource...

but code+tests is already depending on it so hard to change

```go
func (builder *AnnotationTargetBuilder) Limit(limit int64) *AnnotationTargetBuilder
```

### <span class="badge object-method"></span> MatchAny

Only required/valid for the grafana datasource...

but code+tests is already depending on it so hard to change

```go
func (builder *AnnotationTargetBuilder) MatchAny(matchAny bool) *AnnotationTargetBuilder
```

### <span class="badge object-method"></span> Tags

Only required/valid for the grafana datasource...

but code+tests is already depending on it so hard to change

```go
func (builder *AnnotationTargetBuilder) Tags(tags []string) *AnnotationTargetBuilder
```

### <span class="badge object-method"></span> Type

Only required/valid for the grafana datasource...

but code+tests is already depending on it so hard to change

```go
func (builder *AnnotationTargetBuilder) Type(typeArg string) *AnnotationTargetBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationTarget](./object-AnnotationTarget.md)
