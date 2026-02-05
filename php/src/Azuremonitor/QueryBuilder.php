<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind>
 */
class QueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
    $this->internal->kind = "DataQuery";
    $this->internal->group = "grafana-azure-monitor-datasource";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function version(string $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->hide = $hide;
    
        return $this;
    }

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    /**
     * Azure subscription containing the resource(s) to be queried.
     * Also used for template variable queries
     */
    public function subscription(string $subscription): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->subscription = $subscription;
    
        return $this;
    }

    /**
     * Subscriptions to be queried via Azure Resource Graph.
     * @param array<string> $subscriptions
     */
    public function subscriptions(array $subscriptions): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->subscriptions = $subscriptions;
    
        return $this;
    }

    /**
     * Azure Monitor Metrics sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricQuery> $azureMonitor
     */
    public function azureMonitor(\Grafana\Foundation\Cog\Builder $azureMonitor): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $azureMonitorResource = $azureMonitor->build();
        $this->internal->spec->azureMonitor = $azureMonitorResource;
    
        return $this;
    }

    /**
     * Azure Monitor Logs sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureLogsQuery> $azureLogAnalytics
     */
    public function azureLogAnalytics(\Grafana\Foundation\Cog\Builder $azureLogAnalytics): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $azureLogAnalyticsResource = $azureLogAnalytics->build();
        $this->internal->spec->azureLogAnalytics = $azureLogAnalyticsResource;
    
        return $this;
    }

    /**
     * Azure Resource Graph sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery> $azureResourceGraph
     */
    public function azureResourceGraph(\Grafana\Foundation\Cog\Builder $azureResourceGraph): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $azureResourceGraphResource = $azureResourceGraph->build();
        $this->internal->spec->azureResourceGraph = $azureResourceGraphResource;
    
        return $this;
    }

    /**
     * Application Insights Traces sub-query properties.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureTracesQuery> $azureTraces
     */
    public function azureTraces(\Grafana\Foundation\Cog\Builder $azureTraces): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $azureTracesResource = $azureTraces->build();
        $this->internal->spec->azureTraces = $azureTracesResource;
    
        return $this;
    }

    /**
     * @deprecated Legacy template variable support.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\SubscriptionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceGroupsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\ResourceNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamespaceQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\MetricNamesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\WorkspacesQuery>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\UnknownQuery> $grafanaTemplateVariableFn
     */
    public function grafanaTemplateVariableFn( $grafanaTemplateVariableFn): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $grafanaTemplateVariableFnResource = $grafanaTemplateVariableFn->build();
        $this->internal->spec->grafanaTemplateVariableFn = $grafanaTemplateVariableFnResource;
    
        return $this;
    }

    /**
     * Resource group used in template variable queries
     */
    public function resourceGroup(string $resourceGroup): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->resourceGroup = $resourceGroup;
    
        return $this;
    }

    /**
     * Namespace used in template variable queries
     */
    public function namespace(string $namespace): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->namespace = $namespace;
    
        return $this;
    }

    /**
     * Resource used in template variable queries
     */
    public function resource(string $resource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->resource = $resource;
    
        return $this;
    }

    /**
     * Region used in template variable queries
     */
    public function region(string $region): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->region = $region;
    
        return $this;
    }

    /**
     * Custom namespace used in template variable queries
     */
    public function customNamespace(string $customNamespace): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->customNamespace = $customNamespace;
    
        return $this;
    }

    /**
     * Used only for exemplar queries from Prometheus
     */
    public function query(string $query): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Azuremonitor\AzureMonitorQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery);
        $this->internal->spec->query = $query;
    
        return $this;
    }

}
