<?php

namespace Grafana\Foundation\Azuremonitor;

/**
 * Azure Monitor Logs sub-query properties
 */
class AzureLogsQuery implements \JsonSerializable
{
    /**
     * KQL query to be executed.
     */
    public ?string $query;

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
     * If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
     */
    public ?bool $intersectTime;

    /**
     * Workspace ID. This was removed in Grafana 8, but remains for backwards compat
     */
    public ?string $workspace;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resource;

    /**
     * @param string|null $query
     * @param \Grafana\Foundation\Azuremonitor\ResultFormat|null $resultFormat
     * @param array<string>|null $resources
     * @param bool|null $intersectTime
     * @param string|null $workspace
     * @param string|null $resource
     */
    public function __construct(?string $query = null, ?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat = null, ?array $resources = null, ?bool $intersectTime = null, ?string $workspace = null, ?string $resource = null)
    {
        $this->query = $query;
        $this->resultFormat = $resultFormat;
        $this->resources = $resources;
        $this->intersectTime = $intersectTime;
        $this->workspace = $workspace;
        $this->resource = $resource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: string, resultFormat?: string, resources?: array<string>, intersectTime?: bool, workspace?: string, resource?: string} $inputData */
        $data = $inputData;
        return new self(
            query: $data["query"] ?? null,
            resultFormat: isset($data["resultFormat"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\ResultFormat::fromValue($input); })($data["resultFormat"]) : null,
            resources: $data["resources"] ?? null,
            intersectTime: $data["intersectTime"] ?? null,
            workspace: $data["workspace"] ?? null,
            resource: $data["resource"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->query)) {
            $data->query = $this->query;
        }
        if (isset($this->resultFormat)) {
            $data->resultFormat = $this->resultFormat;
        }
        if (isset($this->resources)) {
            $data->resources = $this->resources;
        }
        if (isset($this->intersectTime)) {
            $data->intersectTime = $this->intersectTime;
        }
        if (isset($this->workspace)) {
            $data->workspace = $this->workspace;
        }
        if (isset($this->resource)) {
            $data->resource = $this->resource;
        }
        return $data;
    }
}
