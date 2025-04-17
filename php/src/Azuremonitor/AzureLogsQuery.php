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
     * If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
     */
    public ?bool $dashboardTime;

    /**
     * If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
     */
    public ?string $timeColumn;

    /**
     * Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
     */
    public ?string $workspace;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resource;

    /**
     * @deprecated Use dashboardTime instead
     */
    public ?bool $intersectTime;

    /**
     * @param string|null $query
     * @param \Grafana\Foundation\Azuremonitor\ResultFormat|null $resultFormat
     * @param array<string>|null $resources
     * @param bool|null $dashboardTime
     * @param string|null $timeColumn
     * @param string|null $workspace
     * @param string|null $resource
     * @param bool|null $intersectTime
     */
    public function __construct(?string $query = null, ?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat = null, ?array $resources = null, ?bool $dashboardTime = null, ?string $timeColumn = null, ?string $workspace = null, ?string $resource = null, ?bool $intersectTime = null)
    {
        $this->query = $query;
        $this->resultFormat = $resultFormat;
        $this->resources = $resources;
        $this->dashboardTime = $dashboardTime;
        $this->timeColumn = $timeColumn;
        $this->workspace = $workspace;
        $this->resource = $resource;
        $this->intersectTime = $intersectTime;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: string, resultFormat?: string, resources?: array<string>, dashboardTime?: bool, timeColumn?: string, workspace?: string, resource?: string, intersectTime?: bool} $inputData */
        $data = $inputData;
        return new self(
            query: $data["query"] ?? null,
            resultFormat: isset($data["resultFormat"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\ResultFormat::fromValue($input); })($data["resultFormat"]) : null,
            resources: $data["resources"] ?? null,
            dashboardTime: $data["dashboardTime"] ?? null,
            timeColumn: $data["timeColumn"] ?? null,
            workspace: $data["workspace"] ?? null,
            resource: $data["resource"] ?? null,
            intersectTime: $data["intersectTime"] ?? null,
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
        if (isset($this->dashboardTime)) {
            $data->dashboardTime = $this->dashboardTime;
        }
        if (isset($this->timeColumn)) {
            $data->timeColumn = $this->timeColumn;
        }
        if (isset($this->workspace)) {
            $data->workspace = $this->workspace;
        }
        if (isset($this->resource)) {
            $data->resource = $this->resource;
        }
        if (isset($this->intersectTime)) {
            $data->intersectTime = $this->intersectTime;
        }
        return $data;
    }
}
