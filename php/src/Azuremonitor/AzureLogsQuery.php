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
     * If set to true the query will be run as a basic logs query
     */
    public ?bool $basicLogsQuery;

    /**
     * Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
     */
    public ?string $workspace;

    /**
     * Denotes if logs query editor is in builder mode
     */
    public ?\Grafana\Foundation\Azuremonitor\LogsEditorMode $mode;

    /**
     * Builder query to be executed.
     */
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryExpression $builderQuery;

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
     * @param bool|null $basicLogsQuery
     * @param string|null $workspace
     * @param \Grafana\Foundation\Azuremonitor\LogsEditorMode|null $mode
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryExpression|null $builderQuery
     * @param string|null $resource
     * @param bool|null $intersectTime
     */
    public function __construct(?string $query = null, ?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat = null, ?array $resources = null, ?bool $dashboardTime = null, ?string $timeColumn = null, ?bool $basicLogsQuery = null, ?string $workspace = null, ?\Grafana\Foundation\Azuremonitor\LogsEditorMode $mode = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryExpression $builderQuery = null, ?string $resource = null, ?bool $intersectTime = null)
    {
        $this->query = $query;
        $this->resultFormat = $resultFormat;
        $this->resources = $resources;
        $this->dashboardTime = $dashboardTime;
        $this->timeColumn = $timeColumn;
        $this->basicLogsQuery = $basicLogsQuery;
        $this->workspace = $workspace;
        $this->mode = $mode;
        $this->builderQuery = $builderQuery;
        $this->resource = $resource;
        $this->intersectTime = $intersectTime;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: string, resultFormat?: string, resources?: array<string>, dashboardTime?: bool, timeColumn?: string, basicLogsQuery?: bool, workspace?: string, mode?: string, builderQuery?: mixed, resource?: string, intersectTime?: bool} $inputData */
        $data = $inputData;
        return new self(
            query: $data["query"] ?? null,
            resultFormat: isset($data["resultFormat"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\ResultFormat::fromValue($input); })($data["resultFormat"]) : null,
            resources: $data["resources"] ?? null,
            dashboardTime: $data["dashboardTime"] ?? null,
            timeColumn: $data["timeColumn"] ?? null,
            basicLogsQuery: $data["basicLogsQuery"] ?? null,
            workspace: $data["workspace"] ?? null,
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\LogsEditorMode::fromValue($input); })($data["mode"]) : null,
            builderQuery: isset($data["builderQuery"]) ? (function($input) {
    	/** @var array{from?: mixed, columns?: mixed, where?: mixed, reduce?: mixed, groupBy?: mixed, limit?: int, orderBy?: mixed, fuzzySearch?: mixed, timeFilter?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryExpression::fromArray($val);
    })($data["builderQuery"]) : null,
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
        if (isset($this->basicLogsQuery)) {
            $data->basicLogsQuery = $this->basicLogsQuery;
        }
        if (isset($this->workspace)) {
            $data->workspace = $this->workspace;
        }
        if (isset($this->mode)) {
            $data->mode = $this->mode;
        }
        if (isset($this->builderQuery)) {
            $data->builderQuery = $this->builderQuery;
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
