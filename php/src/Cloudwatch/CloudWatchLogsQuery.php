<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * Shape of a CloudWatch Logs query
 */
class CloudWatchLogsQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Whether a query is a Metrics, Logs, or Annotations query
     */
    public \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode;

    public string $id;

    /**
     * AWS region to query for the logs
     */
    public string $region;

    /**
     * The CloudWatch Logs Insights query to execute
     */
    public ?string $expression;

    /**
     * Fields to group the results by, this field is automatically populated whenever the query is updated
     * @var array<string>|null
     */
    public ?array $statsGroups;

    /**
     * Log groups to query
     * @var array<\Grafana\Foundation\Cloudwatch\LogGroup>|null
     */
    public ?array $logGroups;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * @deprecated use logGroups
     * @var array<string>|null
     */
    public ?array $logGroupNames;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @param \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode|null $queryMode
     * @param string|null $id
     * @param string|null $region
     * @param string|null $expression
     * @param array<string>|null $statsGroups
     * @param array<\Grafana\Foundation\Cloudwatch\LogGroup>|null $logGroups
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param array<string>|null $logGroupNames
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode $queryMode = null, ?string $id = null, ?string $region = null, ?string $expression = null, ?array $statsGroups = null, ?array $logGroups = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?array $logGroupNames = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->queryMode = $queryMode ?: \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::logs();
        $this->id = $id ?: "";
        $this->region = $region ?: "";
        $this->expression = $expression;
        $this->statsGroups = $statsGroups;
        $this->logGroups = $logGroups;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->logGroupNames = $logGroupNames;
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{queryMode?: string, id?: string, region?: string, expression?: string, statsGroups?: array<string>, logGroups?: array<mixed>, refId?: string, hide?: bool, queryType?: string, logGroupNames?: array<string>, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            queryMode: isset($data["queryMode"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::fromValue($input); })($data["queryMode"]) : null,
            id: $data["id"] ?? null,
            region: $data["region"] ?? null,
            expression: $data["expression"] ?? null,
            statsGroups: $data["statsGroups"] ?? null,
            logGroups: array_filter(array_map((function($input) {
    	/** @var array{arn?: string, name?: string, accountId?: string, accountLabel?: string} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\LogGroup::fromArray($val);
    }), $data["logGroups"] ?? [])),
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            logGroupNames: $data["logGroupNames"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "queryMode" => $this->queryMode,
            "id" => $this->id,
            "region" => $this->region,
            "refId" => $this->refId,
        ];
        if (isset($this->expression)) {
            $data["expression"] = $this->expression;
        }
        if (isset($this->statsGroups)) {
            $data["statsGroups"] = $this->statsGroups;
        }
        if (isset($this->logGroups)) {
            $data["logGroups"] = $this->logGroups;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->logGroupNames)) {
            $data["logGroupNames"] = $this->logGroupNames;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "cloudwatch";
    }
}
