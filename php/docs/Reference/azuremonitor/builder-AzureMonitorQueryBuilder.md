---
title: <span class="badge builder"></span> AzureMonitorQueryBuilder
---
# <span class="badge builder"></span> AzureMonitorQueryBuilder

## Constructor

```php
new AzureMonitorQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureLogsQuery> $azureLogAnalytics

```php
azureLogAnalytics(\Grafana\Foundation\Cog\Builder $azureLogAnalytics)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricQuery> $azureMonitor

```php
azureMonitor(\Grafana\Foundation\Cog\Builder $azureMonitor)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery> $azureResourceGraph

```php
azureResourceGraph(\Grafana\Foundation\Cog\Builder $azureResourceGraph)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesQuery> $azureTraces

```php
azureTraces(\Grafana\Foundation\Cog\Builder $azureTraces)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\SubscriptionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\WorkspacesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\UnknownQuery> $grafanaTemplateVariableFn

```php
grafanaTemplateVariableFn($grafanaTemplateVariableFn)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> namespace

```php
namespace(string $namespace)
```

### <span class="badge object-method"></span> query

Used only for exemplar queries from Prometheus

```php
query(string $query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> region

```php
region(string $region)
```

### <span class="badge object-method"></span> resource

```php
resource(string $resource)
```

### <span class="badge object-method"></span> resourceGroup

Template variables params. These exist for backwards compatiblity with legacy template variables.

```php
resourceGroup(string $resourceGroup)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

```php
subscription(string $subscription)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

@param array<string> $subscriptions

```php
subscriptions(array $subscriptions)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
