---
title: <span class="badge object-type-class"></span> AzureMonitorQuery
---
# <span class="badge object-type-class"></span> AzureMonitorQuery

## Definition

```php
class AzureMonitorQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * Azure subscription containing the resource(s) to be queried.
     */
    public ?string $subscription;

    /**
     * Subscriptions to be queried via Azure Resource Graph.
     * @var array<string>|null
     */
    public ?array $subscriptions;

    /**
     * Azure Monitor Metrics sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureMetricQuery $azureMonitor;

    /**
     * Azure Monitor Logs sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureLogsQuery $azureLogAnalytics;

    /**
     * Azure Resource Graph sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery $azureResourceGraph;

    /**
     * Application Insights Traces sub-query properties.
     */
    public ?\Grafana\Foundation\Azuremonitor\AzureTracesQuery $azureTraces;

    /**
     * @deprecated Legacy template variable support.
     * @var \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery|\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery|\Grafana\Foundation\Azuremonitor\SubscriptionsQuery|\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery|\Grafana\Foundation\Azuremonitor\ResourceNamesQuery|\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery|\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery|\Grafana\Foundation\Azuremonitor\MetricNamesQuery|\Grafana\Foundation\Azuremonitor\WorkspacesQuery|\Grafana\Foundation\Azuremonitor\UnknownQuery
     */
    public $grafanaTemplateVariableFn;

    /**
     * Template variables params. These exist for backwards compatiblity with legacy template variables.
     */
    public ?string $resourceGroup;

    public ?string $namespace;

    public ?string $resource;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Azure Monitor query type.
     * queryType: #AzureQueryType
     */
    public ?string $region;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

### <span class="badge object-method"></span> dataqueryType

Returns the type of this dataquery object.

```php
dataqueryType()
```

## See also

 * <span class="badge builder"></span> [AzureMonitorQueryBuilder](./builder-AzureMonitorQueryBuilder.md)
