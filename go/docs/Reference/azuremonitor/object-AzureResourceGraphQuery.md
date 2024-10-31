---
title: <span class="badge object-type-struct"></span> AzureResourceGraphQuery
---
# <span class="badge object-type-struct"></span> AzureResourceGraphQuery

## Definition

```go
type AzureResourceGraphQuery struct {
    // Azure Resource Graph KQL query to be executed.
    Query *string `json:"query,omitempty"`
    // Specifies the format results should be returned as. Defaults to table.
    ResultFormat *string `json:"resultFormat,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureResourceGraphQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureResourceGraphQuery *AzureResourceGraphQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureResourceGraphQuery` objects.

```go
func (azureResourceGraphQuery *AzureResourceGraphQuery) Equals(other AzureResourceGraphQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureResourceGraphQuery` fields for violations and returns them.

```go
func (azureResourceGraphQuery *AzureResourceGraphQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureResourceGraphQueryBuilder](./builder-AzureResourceGraphQueryBuilder.md)
