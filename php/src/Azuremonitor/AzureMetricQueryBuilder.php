<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricQuery>
 */
class AzureMetricQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Azuremonitor\AzureMetricQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Azuremonitor\AzureMetricQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Azuremonitor\AzureMetricQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Array of resource URIs to be queried.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>> $resources
     */
    public function resources(array $resources): static
    {
            $resourcesResources = [];
            foreach ($resources as $r1) {
                    $resourcesResources[] = $r1->build();
            }
        $this->internal->resources = $resourcesResources;
    
        return $this;
    }
    /**
     * metricNamespace is used as the resource type (or resource namespace).
     * It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
     * Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
     */
    public function metricNamespace(string $metricNamespace): static
    {
        $this->internal->metricNamespace = $metricNamespace;
    
        return $this;
    }
    /**
     * Used as the value for the metricNamespace property when it's different from the resource namespace.
     */
    public function customNamespace(string $customNamespace): static
    {
        $this->internal->customNamespace = $customNamespace;
    
        return $this;
    }
    /**
     * The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
     */
    public function metricName(string $metricName): static
    {
        $this->internal->metricName = $metricName;
    
        return $this;
    }
    /**
     * The Azure region containing the resource(s).
     */
    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }
    /**
     * The granularity of data points to be queried. Defaults to auto.
     */
    public function timeGrain(string $timeGrain): static
    {
        $this->internal->timeGrain = $timeGrain;
    
        return $this;
    }
    /**
     * The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
     */
    public function aggregation(string $aggregation): static
    {
        $this->internal->aggregation = $aggregation;
    
        return $this;
    }
    /**
     * Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>> $dimensionFilters
     */
    public function dimensionFilters(array $dimensionFilters): static
    {
            $dimensionFiltersResources = [];
            foreach ($dimensionFilters as $r1) {
                    $dimensionFiltersResources[] = $r1->build();
            }
        $this->internal->dimensionFilters = $dimensionFiltersResources;
    
        return $this;
    }
    /**
     * Maximum number of records to return. Defaults to 10.
     */
    public function top(string $top): static
    {
        $this->internal->top = $top;
    
        return $this;
    }
    /**
     * Time grains that are supported by the metric.
     * @param array<int> $allowedTimeGrainsMs
     */
    public function allowedTimeGrainsMs(array $allowedTimeGrainsMs): static
    {
        $this->internal->allowedTimeGrainsMs = $allowedTimeGrainsMs;
    
        return $this;
    }
    /**
     * Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
     */
    public function alias(string $alias): static
    {
        $this->internal->alias = $alias;
    
        return $this;
    }
    /**
     * @deprecated
     */
    public function timeGrainUnit(string $timeGrainUnit): static
    {
        $this->internal->timeGrainUnit = $timeGrainUnit;
    
        return $this;
    }
    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public function dimension(string $dimension): static
    {
        $this->internal->dimension = $dimension;
    
        return $this;
    }
    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public function dimensionFilter(string $dimensionFilter): static
    {
        $this->internal->dimensionFilter = $dimensionFilter;
    
        return $this;
    }
    /**
     * @deprecated Use metricNamespace instead
     */
    public function metricDefinition(string $metricDefinition): static
    {
        $this->internal->metricDefinition = $metricDefinition;
    
        return $this;
    }
    /**
     * @deprecated Use resourceGroup, resourceName and metricNamespace instead
     */
    public function resourceUri(string $resourceUri): static
    {
        $this->internal->resourceUri = $resourceUri;
    
        return $this;
    }
    /**
     * @deprecated Use resources instead
     */
    public function resourceGroup(string $resourceGroup): static
    {
        $this->internal->resourceGroup = $resourceGroup;
    
        return $this;
    }
    /**
     * @deprecated Use resources instead
     */
    public function resourceName(string $resourceName): static
    {
        $this->internal->resourceName = $resourceName;
    
        return $this;
    }

}
