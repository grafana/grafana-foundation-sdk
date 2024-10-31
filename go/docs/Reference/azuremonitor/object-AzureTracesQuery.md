---
title: <span class="badge object-type-struct"></span> AzureTracesQuery
---
# <span class="badge object-type-struct"></span> AzureTracesQuery

Application Insights Traces sub-query properties

## Definition

```go
type AzureTracesQuery struct {
    // Specifies the format results should be returned as.
    ResultFormat *azuremonitor.ResultFormat `json:"resultFormat,omitempty"`
    // Array of resource URIs to be queried.
    Resources []string `json:"resources,omitempty"`
    // Operation ID. Used only for Traces queries.
    OperationId *string `json:"operationId,omitempty"`
    // Types of events to filter by.
    TraceTypes []string `json:"traceTypes,omitempty"`
    // Filters for property values.
    Filters []azuremonitor.AzureTracesFilter `json:"filters,omitempty"`
    // KQL query to be executed.
    Query *string `json:"query,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureTracesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureTracesQuery *AzureTracesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureTracesQuery` objects.

```go
func (azureTracesQuery *AzureTracesQuery) Equals(other AzureTracesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureTracesQuery` fields for violations and returns them.

```go
func (azureTracesQuery *AzureTracesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureTracesQueryBuilder](./builder-AzureTracesQueryBuilder.md)
