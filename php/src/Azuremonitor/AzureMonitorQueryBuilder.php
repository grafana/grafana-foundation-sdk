<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMonitorQuery>
 */
class AzureMonitorQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureMonitorQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureMonitorQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * Azure subscription containing the resource(s) to be queried.
     */
    public function subscription(string $subscription): static
    {
        $this->internal->subscription = $subscription;
    
        return $this;
    }
    /**
     * Subscriptions to be queried via Azure Resource Graph.
     * @param array<string> $subscriptions
     */
    public function subscriptions(array $subscriptions): static
    {
        $this->internal->subscriptions = $subscriptions;
    
        return $this;
    }
    /**
     * Azure Monitor Metrics sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricQuery> $azureMonitor
     */
    public function azureMonitor(\Grafana\Foundation\Cog\Builder $azureMonitor): static
    {
        $azureMonitorResource = $azureMonitor->build();
        $this->internal->azureMonitor = $azureMonitorResource;
    
        return $this;
    }
    /**
     * Azure Monitor Logs sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureLogsQuery> $azureLogAnalytics
     */
    public function azureLogAnalytics(\Grafana\Foundation\Cog\Builder $azureLogAnalytics): static
    {
        $azureLogAnalyticsResource = $azureLogAnalytics->build();
        $this->internal->azureLogAnalytics = $azureLogAnalyticsResource;
    
        return $this;
    }
    /**
     * Azure Resource Graph sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery> $azureResourceGraph
     */
    public function azureResourceGraph(\Grafana\Foundation\Cog\Builder $azureResourceGraph): static
    {
        $azureResourceGraphResource = $azureResourceGraph->build();
        $this->internal->azureResourceGraph = $azureResourceGraphResource;
    
        return $this;
    }
    /**
     * Application Insights Traces sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesQuery> $azureTraces
     */
    public function azureTraces(\Grafana\Foundation\Cog\Builder $azureTraces): static
    {
        $azureTracesResource = $azureTraces->build();
        $this->internal->azureTraces = $azureTracesResource;
    
        return $this;
    }
    /**
     * @deprecated Legacy template variable support.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\SubscriptionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\WorkspacesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\UnknownQuery> $grafanaTemplateVariableFn
     */
    public function grafanaTemplateVariableFn( $grafanaTemplateVariableFn): static
    {
        $grafanaTemplateVariableFnResource = $grafanaTemplateVariableFn->build();
        $this->internal->grafanaTemplateVariableFn = $grafanaTemplateVariableFnResource;
    
        return $this;
    }
    /**
     * Template variables params. These exist for backwards compatiblity with legacy template variables.
     */
    public function resourceGroup(string $resourceGroup): static
    {
        $this->internal->resourceGroup = $resourceGroup;
    
        return $this;
    }
    public function namespace(string $namespace): static
    {
        $this->internal->namespace = $namespace;
    
        return $this;
    }
    public function resource(string $resource): static
    {
        $this->internal->resource = $resource;
    
        return $this;
    }
    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * Azure Monitor query type.
     * queryType: #AzureQueryType
     */
    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }

}
