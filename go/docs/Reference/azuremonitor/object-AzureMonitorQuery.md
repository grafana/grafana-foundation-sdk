---
title: <span class="badge object-type-struct"></span> AzureMonitorQuery
---
# <span class="badge object-type-struct"></span> AzureMonitorQuery

## Definition

```go
type AzureMonitorQuery struct {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    RefId string `json:"refId"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Specify the query flavor
    // TODO make this required and give it a default
    QueryType *string `json:"queryType,omitempty"`
    // Azure subscription containing the resource(s) to be queried.
    Subscription *string `json:"subscription,omitempty"`
    // Subscriptions to be queried via Azure Resource Graph.
    Subscriptions []string `json:"subscriptions,omitempty"`
    // Azure Monitor Metrics sub-query properties.
    AzureMonitor *azuremonitor.AzureMetricQuery `json:"azureMonitor,omitempty"`
    // Azure Monitor Logs sub-query properties.
    AzureLogAnalytics *azuremonitor.AzureLogsQuery `json:"azureLogAnalytics,omitempty"`
    // Azure Resource Graph sub-query properties.
    AzureResourceGraph *azuremonitor.AzureResourceGraphQuery `json:"azureResourceGraph,omitempty"`
    // Application Insights Traces sub-query properties.
    AzureTraces *azuremonitor.AzureTracesQuery `json:"azureTraces,omitempty"`
    // @deprecated Legacy template variable support.
    GrafanaTemplateVariableFn *azuremonitor.GrafanaTemplateVariableQuery `json:"grafanaTemplateVariableFn,omitempty"`
    // Template variables params. These exist for backwards compatiblity with legacy template variables.
    ResourceGroup *string `json:"resourceGroup,omitempty"`
    Namespace *string `json:"namespace,omitempty"`
    Resource *string `json:"resource,omitempty"`
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // Azure Monitor query type.
    // queryType: #AzureQueryType
    Region *string `json:"region,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMonitorQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureMonitorQuery *AzureMonitorQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureMonitorQuery` objects.

```go
func (azureMonitorQuery *AzureMonitorQuery) Equals(other AzureMonitorQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureMonitorQuery` fields for violations and returns them.

```go
func (azureMonitorQuery *AzureMonitorQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureMonitorQueryBuilder](./builder-AzureMonitorQueryBuilder.md)
