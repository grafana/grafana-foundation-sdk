---
title: <span class="badge builder"></span> AnnotationEventFieldMappingBuilder
---
# <span class="badge builder"></span> AnnotationEventFieldMappingBuilder

## Constructor

```go
func NewAnnotationEventFieldMappingBuilder() *AnnotationEventFieldMappingBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationEventFieldMappingBuilder) Build() (AnnotationEventFieldMapping, error)
```

### <span class="badge object-method"></span> Regex

Regular expression to apply to the field value

```go
func (builder *AnnotationEventFieldMappingBuilder) Regex(regex string) *AnnotationEventFieldMappingBuilder
```

### <span class="badge object-method"></span> Source

Source type for the field value

```go
func (builder *AnnotationEventFieldMappingBuilder) Source(source string) *AnnotationEventFieldMappingBuilder
```

### <span class="badge object-method"></span> Value

Constant value to use when source is "text"

```go
func (builder *AnnotationEventFieldMappingBuilder) Value(value string) *AnnotationEventFieldMappingBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationEventFieldMapping](./object-AnnotationEventFieldMapping.md)
