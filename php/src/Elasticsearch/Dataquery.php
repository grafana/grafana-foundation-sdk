<?php

namespace Grafana\Foundation\Elasticsearch;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Alias pattern
     */
    public ?string $alias;

    /**
     * Lucene query
     */
    public ?string $query;

    /**
     * Name of time field
     */
    public ?string $timeField;

    /**
     * List of bucket aggregations
     * @var array<\Grafana\Foundation\Elasticsearch\DateHistogram|\Grafana\Foundation\Elasticsearch\Histogram|\Grafana\Foundation\Elasticsearch\Terms|\Grafana\Foundation\Elasticsearch\Filters|\Grafana\Foundation\Elasticsearch\GeoHashGrid|\Grafana\Foundation\Elasticsearch\Nested>|null
     */
    public ?array $bucketAggs;

    /**
     * List of metric aggregations
     * @var array<\Grafana\Foundation\Elasticsearch\Count|\Grafana\Foundation\Elasticsearch\MovingAverage|\Grafana\Foundation\Elasticsearch\Derivative|\Grafana\Foundation\Elasticsearch\CumulativeSum|\Grafana\Foundation\Elasticsearch\BucketScript|\Grafana\Foundation\Elasticsearch\SerialDiff|\Grafana\Foundation\Elasticsearch\RawData|\Grafana\Foundation\Elasticsearch\RawDocument|\Grafana\Foundation\Elasticsearch\UniqueCount|\Grafana\Foundation\Elasticsearch\Percentiles|\Grafana\Foundation\Elasticsearch\ExtendedStats|\Grafana\Foundation\Elasticsearch\Min|\Grafana\Foundation\Elasticsearch\Max|\Grafana\Foundation\Elasticsearch\Sum|\Grafana\Foundation\Elasticsearch\Average|\Grafana\Foundation\Elasticsearch\MovingFunction|\Grafana\Foundation\Elasticsearch\Logs|\Grafana\Foundation\Elasticsearch\Rate|\Grafana\Foundation\Elasticsearch\TopMetrics>|null
     */
    public ?array $metrics;

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
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @param string|null $alias
     * @param string|null $query
     * @param string|null $timeField
     * @param array<\Grafana\Foundation\Elasticsearch\DateHistogram|\Grafana\Foundation\Elasticsearch\Histogram|\Grafana\Foundation\Elasticsearch\Terms|\Grafana\Foundation\Elasticsearch\Filters|\Grafana\Foundation\Elasticsearch\GeoHashGrid|\Grafana\Foundation\Elasticsearch\Nested>|null $bucketAggs
     * @param array<\Grafana\Foundation\Elasticsearch\Count|\Grafana\Foundation\Elasticsearch\MovingAverage|\Grafana\Foundation\Elasticsearch\Derivative|\Grafana\Foundation\Elasticsearch\CumulativeSum|\Grafana\Foundation\Elasticsearch\BucketScript|\Grafana\Foundation\Elasticsearch\SerialDiff|\Grafana\Foundation\Elasticsearch\RawData|\Grafana\Foundation\Elasticsearch\RawDocument|\Grafana\Foundation\Elasticsearch\UniqueCount|\Grafana\Foundation\Elasticsearch\Percentiles|\Grafana\Foundation\Elasticsearch\ExtendedStats|\Grafana\Foundation\Elasticsearch\Min|\Grafana\Foundation\Elasticsearch\Max|\Grafana\Foundation\Elasticsearch\Sum|\Grafana\Foundation\Elasticsearch\Average|\Grafana\Foundation\Elasticsearch\MovingFunction|\Grafana\Foundation\Elasticsearch\Logs|\Grafana\Foundation\Elasticsearch\Rate|\Grafana\Foundation\Elasticsearch\TopMetrics>|null $metrics
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?string $alias = null, ?string $query = null, ?string $timeField = null, ?array $bucketAggs = null, ?array $metrics = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->alias = $alias;
        $this->query = $query;
        $this->timeField = $timeField;
        $this->bucketAggs = $bucketAggs;
        $this->metrics = $metrics;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{alias?: string, query?: string, timeField?: string, bucketAggs?: array<mixed|mixed|mixed|mixed|mixed|mixed>, metrics?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            alias: $data["alias"] ?? null,
            query: $data["query"] ?? null,
            timeField: $data["timeField"] ?? null,
            bucketAggs: !empty($data["bucketAggs"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
    
        switch ($input["type"]) {
        case "date_histogram":
            return DateHistogram::fromArray($input);
        case "filters":
            return Filters::fromArray($input);
        case "geohash_grid":
            return GeoHashGrid::fromArray($input);
        case "histogram":
            return Histogram::fromArray($input);
        case "nested":
            return Nested::fromArray($input);
        case "terms":
            return Terms::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["bucketAggs"]) : null,
            metrics: !empty($data["metrics"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
    
        switch ($input["type"]) {
        case "avg":
            return Average::fromArray($input);
        case "bucket_script":
            return BucketScript::fromArray($input);
        case "cardinality":
            return UniqueCount::fromArray($input);
        case "count":
            return Count::fromArray($input);
        case "cumulative_sum":
            return CumulativeSum::fromArray($input);
        case "derivative":
            return Derivative::fromArray($input);
        case "extended_stats":
            return ExtendedStats::fromArray($input);
        case "logs":
            return Logs::fromArray($input);
        case "max":
            return Max::fromArray($input);
        case "min":
            return Min::fromArray($input);
        case "moving_avg":
            return MovingAverage::fromArray($input);
        case "moving_fn":
            return MovingFunction::fromArray($input);
        case "percentiles":
            return Percentiles::fromArray($input);
        case "rate":
            return Rate::fromArray($input);
        case "raw_data":
            return RawData::fromArray($input);
        case "raw_document":
            return RawDocument::fromArray($input);
        case "serial_diff":
            return SerialDiff::fromArray($input);
        case "sum":
            return Sum::fromArray($input);
        case "top_metrics":
            return TopMetrics::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["metrics"]) : null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
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
            "refId" => $this->refId,
        ];
        if (isset($this->alias)) {
            $data["alias"] = $this->alias;
        }
        if (isset($this->query)) {
            $data["query"] = $this->query;
        }
        if (isset($this->timeField)) {
            $data["timeField"] = $this->timeField;
        }
        if (isset($this->bucketAggs)) {
            $data["bucketAggs"] = $this->bucketAggs;
        }
        if (isset($this->metrics)) {
            $data["metrics"] = $this->metrics;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "elasticsearch";
    }
}
