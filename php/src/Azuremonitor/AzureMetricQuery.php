<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureMetricQuery implements \JsonSerializable
{
    /**
     * Array of resource URIs to be queried.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>|null
     */
    public ?array $resources;

    /**
     * metricNamespace is used as the resource type (or resource namespace).
     * It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
     * Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
     */
    public ?string $metricNamespace;

    /**
     * Used as the value for the metricNamespace property when it's different from the resource namespace.
     */
    public ?string $customNamespace;

    /**
     * The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
     */
    public ?string $metricName;

    /**
     * The Azure region containing the resource(s).
     */
    public ?string $region;

    /**
     * The granularity of data points to be queried. Defaults to auto.
     */
    public ?string $timeGrain;

    /**
     * The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
     */
    public ?string $aggregation;

    /**
     * Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>|null
     */
    public ?array $dimensionFilters;

    /**
     * Maximum number of records to return. Defaults to 10.
     */
    public ?string $top;

    /**
     * Time grains that are supported by the metric.
     * @var array<int>|null
     */
    public ?array $allowedTimeGrainsMs;

    /**
     * Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
     */
    public ?string $alias;

    /**
     * @deprecated
     */
    public ?string $timeGrainUnit;

    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public ?string $dimension;

    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public ?string $dimensionFilter;

    /**
     * @deprecated Use metricNamespace instead
     */
    public ?string $metricDefinition;

    /**
     * @deprecated Use resourceGroup, resourceName and metricNamespace instead
     */
    public ?string $resourceUri;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resourceGroup;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resourceName;

    /**
     * @param array<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>|null $resources
     * @param string|null $metricNamespace
     * @param string|null $customNamespace
     * @param string|null $metricName
     * @param string|null $region
     * @param string|null $timeGrain
     * @param string|null $aggregation
     * @param array<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>|null $dimensionFilters
     * @param string|null $top
     * @param array<int>|null $allowedTimeGrainsMs
     * @param string|null $alias
     * @param string|null $timeGrainUnit
     * @param string|null $dimension
     * @param string|null $dimensionFilter
     * @param string|null $metricDefinition
     * @param string|null $resourceUri
     * @param string|null $resourceGroup
     * @param string|null $resourceName
     */
    public function __construct(?array $resources = null, ?string $metricNamespace = null, ?string $customNamespace = null, ?string $metricName = null, ?string $region = null, ?string $timeGrain = null, ?string $aggregation = null, ?array $dimensionFilters = null, ?string $top = null, ?array $allowedTimeGrainsMs = null, ?string $alias = null, ?string $timeGrainUnit = null, ?string $dimension = null, ?string $dimensionFilter = null, ?string $metricDefinition = null, ?string $resourceUri = null, ?string $resourceGroup = null, ?string $resourceName = null)
    {
        $this->resources = $resources;
        $this->metricNamespace = $metricNamespace;
        $this->customNamespace = $customNamespace;
        $this->metricName = $metricName;
        $this->region = $region;
        $this->timeGrain = $timeGrain;
        $this->aggregation = $aggregation;
        $this->dimensionFilters = $dimensionFilters;
        $this->top = $top;
        $this->allowedTimeGrainsMs = $allowedTimeGrainsMs;
        $this->alias = $alias;
        $this->timeGrainUnit = $timeGrainUnit;
        $this->dimension = $dimension;
        $this->dimensionFilter = $dimensionFilter;
        $this->metricDefinition = $metricDefinition;
        $this->resourceUri = $resourceUri;
        $this->resourceGroup = $resourceGroup;
        $this->resourceName = $resourceName;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{resources?: array<mixed>, metricNamespace?: string, customNamespace?: string, metricName?: string, region?: string, timeGrain?: string, aggregation?: string, dimensionFilters?: array<mixed>, top?: string, allowedTimeGrainsMs?: array<int>, alias?: string, timeGrainUnit?: string, dimension?: string, dimensionFilter?: string, metricDefinition?: string, resourceUri?: string, resourceGroup?: string, resourceName?: string} $inputData */
        $data = $inputData;
        return new self(
            resources: array_filter(array_map((function($input) {
    	/** @var array{subscription?: string, resourceGroup?: string, resourceName?: string, metricNamespace?: string, region?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureMonitorResource::fromArray($val);
    }), $data["resources"] ?? [])),
            metricNamespace: $data["metricNamespace"] ?? null,
            customNamespace: $data["customNamespace"] ?? null,
            metricName: $data["metricName"] ?? null,
            region: $data["region"] ?? null,
            timeGrain: $data["timeGrain"] ?? null,
            aggregation: $data["aggregation"] ?? null,
            dimensionFilters: array_filter(array_map((function($input) {
    	/** @var array{dimension?: string, operator?: string, filters?: array<string>, filter?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureMetricDimension::fromArray($val);
    }), $data["dimensionFilters"] ?? [])),
            top: $data["top"] ?? null,
            allowedTimeGrainsMs: $data["allowedTimeGrainsMs"] ?? null,
            alias: $data["alias"] ?? null,
            timeGrainUnit: $data["timeGrainUnit"] ?? null,
            dimension: $data["dimension"] ?? null,
            dimensionFilter: $data["dimensionFilter"] ?? null,
            metricDefinition: $data["metricDefinition"] ?? null,
            resourceUri: $data["resourceUri"] ?? null,
            resourceGroup: $data["resourceGroup"] ?? null,
            resourceName: $data["resourceName"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->resources)) {
            $data["resources"] = $this->resources;
        }
        if (isset($this->metricNamespace)) {
            $data["metricNamespace"] = $this->metricNamespace;
        }
        if (isset($this->customNamespace)) {
            $data["customNamespace"] = $this->customNamespace;
        }
        if (isset($this->metricName)) {
            $data["metricName"] = $this->metricName;
        }
        if (isset($this->region)) {
            $data["region"] = $this->region;
        }
        if (isset($this->timeGrain)) {
            $data["timeGrain"] = $this->timeGrain;
        }
        if (isset($this->aggregation)) {
            $data["aggregation"] = $this->aggregation;
        }
        if (isset($this->dimensionFilters)) {
            $data["dimensionFilters"] = $this->dimensionFilters;
        }
        if (isset($this->top)) {
            $data["top"] = $this->top;
        }
        if (isset($this->allowedTimeGrainsMs)) {
            $data["allowedTimeGrainsMs"] = $this->allowedTimeGrainsMs;
        }
        if (isset($this->alias)) {
            $data["alias"] = $this->alias;
        }
        if (isset($this->timeGrainUnit)) {
            $data["timeGrainUnit"] = $this->timeGrainUnit;
        }
        if (isset($this->dimension)) {
            $data["dimension"] = $this->dimension;
        }
        if (isset($this->dimensionFilter)) {
            $data["dimensionFilter"] = $this->dimensionFilter;
        }
        if (isset($this->metricDefinition)) {
            $data["metricDefinition"] = $this->metricDefinition;
        }
        if (isset($this->resourceUri)) {
            $data["resourceUri"] = $this->resourceUri;
        }
        if (isset($this->resourceGroup)) {
            $data["resourceGroup"] = $this->resourceGroup;
        }
        if (isset($this->resourceName)) {
            $data["resourceName"] = $this->resourceName;
        }
        return $data;
    }
}
