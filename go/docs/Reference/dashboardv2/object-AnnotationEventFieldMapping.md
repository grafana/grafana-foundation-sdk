---
title: <span class="badge object-type-struct"></span> AnnotationEventFieldMapping
---
# <span class="badge object-type-struct"></span> AnnotationEventFieldMapping

Annotation event field mapping. Defines how to map a data frame field to an annotation event field.

## Definition

```go
type AnnotationEventFieldMapping struct {
    // Source type for the field value
    Source *string `json:"source,omitempty"`
    // Constant value to use when source is "text"
    Value *string `json:"value,omitempty"`
    // Regular expression to apply to the field value
    Regex *string `json:"regex,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AnnotationEventFieldMapping` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (annotationEventFieldMapping *AnnotationEventFieldMapping) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AnnotationEventFieldMapping` objects.

```go
func (annotationEventFieldMapping *AnnotationEventFieldMapping) Equals(other AnnotationEventFieldMapping) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AnnotationEventFieldMapping` fields for violations and returns them.

```go
func (annotationEventFieldMapping *AnnotationEventFieldMapping) Validate() error
```

## See also

 * <span class="badge builder"></span> [AnnotationEventFieldMappingBuilder](./builder-AnnotationEventFieldMappingBuilder.md)
