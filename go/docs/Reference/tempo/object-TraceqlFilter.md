---
title: <span class="badge object-type-struct"></span> TraceqlFilter
---
# <span class="badge object-type-struct"></span> TraceqlFilter

## Definition

```go
type TraceqlFilter struct {
    // Uniquely identify the filter, will not be used in the query generation
    Id string `json:"id"`
    // The tag for the search filter, for example: .http.status_code, .service.name, status
    Tag *string `json:"tag,omitempty"`
    // The operator that connects the tag to the value, for example: =, >, !=, =~
    Operator *string `json:"operator,omitempty"`
    // The value for the search filter
    Value *tempo.StringOrArrayOfString `json:"value,omitempty"`
    // The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
    ValueType *string `json:"valueType,omitempty"`
    // The scope of the filter, can either be unscoped/all scopes, resource or span
    Scope *tempo.TraceqlSearchScope `json:"scope,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TraceqlFilter` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (traceqlFilter *TraceqlFilter) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TraceqlFilter` objects.

```go
func (traceqlFilter *TraceqlFilter) Equals(other TraceqlFilter) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TraceqlFilter` fields for violations and returns them.

```go
func (traceqlFilter *TraceqlFilter) Validate() error
```

## See also

 * <span class="badge builder"></span> [TraceqlFilterBuilder](./builder-TraceqlFilterBuilder.md)
