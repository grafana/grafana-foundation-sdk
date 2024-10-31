---
title: <span class="badge object-type-struct"></span> BaseGrafanaTemplateVariableQuery
---
# <span class="badge object-type-struct"></span> BaseGrafanaTemplateVariableQuery

## Definition

```go
type BaseGrafanaTemplateVariableQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseGrafanaTemplateVariableQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (baseGrafanaTemplateVariableQuery *BaseGrafanaTemplateVariableQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BaseGrafanaTemplateVariableQuery` objects.

```go
func (baseGrafanaTemplateVariableQuery *BaseGrafanaTemplateVariableQuery) Equals(other BaseGrafanaTemplateVariableQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BaseGrafanaTemplateVariableQuery` fields for violations and returns them.

```go
func (baseGrafanaTemplateVariableQuery *BaseGrafanaTemplateVariableQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [BaseGrafanaTemplateVariableQueryBuilder](./builder-BaseGrafanaTemplateVariableQueryBuilder.md)
