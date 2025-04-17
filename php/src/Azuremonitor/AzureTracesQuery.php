<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * Application Insights Traces sub-query properties
 */
class AzureTracesQuery implements \JsonSerializable
{
    /**
     * Specifies the format results should be returned as.
     */
    public ?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat;

    /**
     * Array of resource URIs to be queried.
     * @var array<string>|null
     */
    public ?array $resources;

    /**
     * Operation ID. Used only for Traces queries.
     */
    public ?string $operationId;

    /**
     * Types of events to filter by.
     * @var array<string>|null
     */
    public ?array $traceTypes;

    /**
     * Filters for property values.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>|null
     */
    public ?array $filters;

    /**
     * KQL query to be executed.
     */
    public ?string $query;

    /**
     * @param \Grafana\Foundation\Azuremonitor\ResultFormat|null $resultFormat
     * @param array<string>|null $resources
     * @param string|null $operationId
     * @param array<string>|null $traceTypes
     * @param array<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>|null $filters
     * @param string|null $query
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat = null, ?array $resources = null, ?string $operationId = null, ?array $traceTypes = null, ?array $filters = null, ?string $query = null)
    {
        $this->resultFormat = $resultFormat;
        $this->resources = $resources;
        $this->operationId = $operationId;
        $this->traceTypes = $traceTypes;
        $this->filters = $filters;
        $this->query = $query;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{resultFormat?: string, resources?: array<string>, operationId?: string, traceTypes?: array<string>, filters?: array<mixed>, query?: string} $inputData */
        $data = $inputData;
        return new self(
            resultFormat: isset($data["resultFormat"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\ResultFormat::fromValue($input); })($data["resultFormat"]) : null,
            resources: $data["resources"] ?? null,
            operationId: $data["operationId"] ?? null,
            traceTypes: $data["traceTypes"] ?? null,
            filters: array_filter(array_map((function($input) {
    	/** @var array{property?: string, operation?: string, filters?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\AzureTracesFilter::fromArray($val);
    }), $data["filters"] ?? [])),
            query: $data["query"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->resultFormat)) {
            $data->resultFormat = $this->resultFormat;
        }
        if (isset($this->resources)) {
            $data->resources = $this->resources;
        }
        if (isset($this->operationId)) {
            $data->operationId = $this->operationId;
        }
        if (isset($this->traceTypes)) {
            $data->traceTypes = $this->traceTypes;
        }
        if (isset($this->filters)) {
            $data->filters = $this->filters;
        }
        if (isset($this->query)) {
            $data->query = $this->query;
        }
        return $data;
    }
}
