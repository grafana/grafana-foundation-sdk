---
title: <span class="badge object-type-struct"></span> MonitorQuery
---
# <span class="badge object-type-struct"></span> MonitorQuery

## Definition

```go
type MonitorQuery struct {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    RefId string `json:"refId"`
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    Hide *bool `json:"hide,omitempty"`
    // Specify the query flavor
    // TODO make this required and give it a default
    QueryType *string `json:"queryType,omitempty"`
    // Azure subscription containing the resource(s) to be queried.
    // Also used for template variable queries
    Subscription *string `json:"subscription,omitempty"`
    // Subscriptions to be queried via Azure Resource Graph.
    Subscriptions []string `json:"subscriptions,omitempty"`
    // Azure Monitor Metrics sub-query properties.
    AzureMonitor *azuremonitor.MetricQuery `json:"azureMonitor,omitempty"`
    // Azure Monitor Logs sub-query properties.
    AzureLogAnalytics *azuremonitor.LogsQuery `json:"azureLogAnalytics,omitempty"`
    // Azure Resource Graph sub-query properties.
    AzureResourceGraph *azuremonitor.ResourceGraphQuery `json:"azureResourceGraph,omitempty"`
    // Application Insights Traces sub-query properties.
    AzureTraces *azuremonitor.TracesQuery `json:"azureTraces,omitempty"`
    // @deprecated Legacy template variable support.
    GrafanaTemplateVariableFn *azuremonitor.GrafanaTemplateVariableQuery `json:"grafanaTemplateVariableFn,omitempty"`
    // Resource group used in template variable queries
    ResourceGroup *string `json:"resourceGroup,omitempty"`
    // Namespace used in template variable queries
    Namespace *string `json:"namespace,omitempty"`
    // Resource used in template variable queries
    Resource *string `json:"resource,omitempty"`
    // Region used in template variable queries
    Region *string `json:"region,omitempty"`
    // Custom namespace used in template variable queries
    CustomNamespace *string `json:"customNamespace,omitempty"`
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    Datasource any `json:"datasource,omitempty"`
    // Used only for exemplar queries from Prometheus
    Query *string `json:"query,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MonitorQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (monitorQuery *MonitorQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MonitorQuery` objects.

```go
func (monitorQuery *MonitorQuery) Equals(other MonitorQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MonitorQuery` fields for violations and returns them.

```go
func (monitorQuery *MonitorQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MonitorQueryBuilder](./builder-MonitorQueryBuilder.md)
