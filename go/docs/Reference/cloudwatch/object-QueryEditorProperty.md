---
title: <span class="badge object-type-struct"></span> QueryEditorProperty
---
# <span class="badge object-type-struct"></span> QueryEditorProperty

## Definition

```go
type QueryEditorProperty struct {
    Type cloudwatch.QueryEditorPropertyType `json:"type"`
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorProperty` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorProperty *QueryEditorProperty) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorProperty` objects.

```go
func (queryEditorProperty *QueryEditorProperty) Equals(other QueryEditorProperty) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorProperty` fields for violations and returns them.

```go
func (queryEditorProperty *QueryEditorProperty) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorPropertyBuilder](./builder-QueryEditorPropertyBuilder.md)
