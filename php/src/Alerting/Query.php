<?php

namespace Grafana\Foundation\Alerting;

class Query implements \JsonSerializable
{
    /**
     * Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
     */
    public ?string $datasourceUid;

    /**
     * JSON is the raw JSON query and includes the above properties as well as custom properties.
     * @var \Grafana\Foundation\Cog\Dataquery|null
     */
    public ?\Grafana\Foundation\Cog\Dataquery $model;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * RelativeTimeRange is the per query start and end time
     * for requests.
     */
    public ?\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange;

    /**
     * @param string|null $datasourceUid
     * @param \Grafana\Foundation\Cog\Dataquery|null $model
     * @param string|null $queryType
     * @param string|null $refId
     * @param \Grafana\Foundation\Alerting\RelativeTimeRange|null $relativeTimeRange
     */
    public function __construct(?string $datasourceUid = null, ?\Grafana\Foundation\Cog\Dataquery $model = null, ?string $queryType = null, ?string $refId = null, ?\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange = null)
    {
        $this->datasourceUid = $datasourceUid;
        $this->model = $model;
        $this->queryType = $queryType;
        $this->refId = $refId;
        $this->relativeTimeRange = $relativeTimeRange;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{datasourceUid?: string, model?: mixed, queryType?: string, refId?: string, relativeTimeRange?: mixed} $inputData */
        $data = $inputData;
        return new self(
            datasourceUid: $data["datasourceUid"] ?? null,
            model: isset($data["model"]) ? (function ($in) {
    	/** @var array{datasource?: array{type?: mixed}} $in */
        $hint = "";
    
        /** @var array<string, mixed> $in */
        return \Grafana\Foundation\Cog\Runtime::get()->dataqueryFromArray($in, $hint);
    })($data["model"]) : null,
            queryType: $data["queryType"] ?? null,
            refId: $data["refId"] ?? null,
            relativeTimeRange: isset($data["relativeTimeRange"]) ? (function($input) {
    	/** @var array{from?: int, to?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\RelativeTimeRange::fromArray($val);
    })($data["relativeTimeRange"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->datasourceUid)) {
            $data->datasourceUid = $this->datasourceUid;
        }
        if (isset($this->model)) {
            $data->model = $this->model;
        }
        if (isset($this->queryType)) {
            $data->queryType = $this->queryType;
        }
        if (isset($this->refId)) {
            $data->refId = $this->refId;
        }
        if (isset($this->relativeTimeRange)) {
            $data->relativeTimeRange = $this->relativeTimeRange;
        }
        return $data;
    }
}
