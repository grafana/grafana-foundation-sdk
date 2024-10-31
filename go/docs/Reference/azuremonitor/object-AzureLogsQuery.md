---
title: <span class="badge object-type-struct"></span> AzureLogsQuery
---
# <span class="badge object-type-struct"></span> AzureLogsQuery

Azure Monitor Logs sub-query properties

## Definition

```go
type AzureLogsQuery struct {
    // KQL query to be executed.
    Query *string `json:"query,omitempty"`
    // Specifies the format results should be returned as.
    ResultFormat *azuremonitor.ResultFormat `json:"resultFormat,omitempty"`
    // Array of resource URIs to be queried.
    Resources []string `json:"resources,omitempty"`
    // If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
    DashboardTime *bool `json:"dashboardTime,omitempty"`
    // If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
    TimeColumn *string `json:"timeColumn,omitempty"`
    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
    Workspace *string `json:"workspace,omitempty"`
    // @deprecated Use resources instead
    Resource *string `json:"resource,omitempty"`
    // @deprecated Use dashboardTime instead
    IntersectTime *bool `json:"intersectTime,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureLogsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureLogsQuery *AzureLogsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureLogsQuery` objects.

```go
func (azureLogsQuery *AzureLogsQuery) Equals(other AzureLogsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureLogsQuery` fields for violations and returns them.

```go
func (azureLogsQuery *AzureLogsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)
