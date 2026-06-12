<?php

namespace Grafana\Foundation\Prometheus;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Additional Ad-hoc filters that take precedence over Scope on conflict.
     * @var array<\Grafana\Foundation\Prometheus\AdhocFilters>|null
     */
    public ?array $adhocFilters;

    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Common\DataSourceRef $datasource;

    /**
     * what we should show in the editor
     * Possible enum values:
     *  - `"builder"` 
     *  - `"code"` 
     */
    public ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode;

    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public ?bool $exemplar;

    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public string $expr;

    /**
     * The response format
     * Possible enum values:
     *  - `"time_series"` 
     *  - `"table"` 
     *  - `"heatmap"` 
     */
    public ?\Grafana\Foundation\Prometheus\PromQueryFormat $format;

    /**
     * Group By parameters to apply to aggregate expressions in the query
     * @var array<string>|null
     */
    public ?array $groupByKeys;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public ?bool $instant;

    /**
     * Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     * Deprecated: use interval
     */
    public ?int $intervalFactor;

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public ?float $intervalMs;

    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public ?string $legendFormat;

    /**
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public ?int $maxDataPoints;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public ?bool $range;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Prometheus\ResultAssertions $resultAssertions;

    /**
     * A set of filters applied to apply to the query
     * @var array<\Grafana\Foundation\Prometheus\Scopes>|null
     */
    public ?array $scopes;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Prometheus\TimeRange $timeRange;

    /**
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public ?string $interval;

    /**
     * @param array<\Grafana\Foundation\Prometheus\AdhocFilters>|null $adhocFilters
     * @param \Grafana\Foundation\Common\DataSourceRef|null $datasource
     * @param \Grafana\Foundation\Prometheus\QueryEditorMode|null $editorMode
     * @param bool|null $exemplar
     * @param string|null $expr
     * @param \Grafana\Foundation\Prometheus\PromQueryFormat|null $format
     * @param array<string>|null $groupByKeys
     * @param bool|null $hide
     * @param bool|null $instant
     * @param int|null $intervalFactor
     * @param float|null $intervalMs
     * @param string|null $legendFormat
     * @param int|null $maxDataPoints
     * @param string|null $queryType
     * @param bool|null $range
     * @param string|null $refId
     * @param \Grafana\Foundation\Prometheus\ResultAssertions|null $resultAssertions
     * @param array<\Grafana\Foundation\Prometheus\Scopes>|null $scopes
     * @param \Grafana\Foundation\Prometheus\TimeRange|null $timeRange
     * @param string|null $interval
     */
    public function __construct(?array $adhocFilters = null, ?\Grafana\Foundation\Common\DataSourceRef $datasource = null, ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode = null, ?bool $exemplar = null, ?string $expr = null, ?\Grafana\Foundation\Prometheus\PromQueryFormat $format = null, ?array $groupByKeys = null, ?bool $hide = null, ?bool $instant = null, ?int $intervalFactor = null, ?float $intervalMs = null, ?string $legendFormat = null, ?int $maxDataPoints = null, ?string $queryType = null, ?bool $range = null, ?string $refId = null, ?\Grafana\Foundation\Prometheus\ResultAssertions $resultAssertions = null, ?array $scopes = null, ?\Grafana\Foundation\Prometheus\TimeRange $timeRange = null, ?string $interval = null)
    {
        $this->adhocFilters = $adhocFilters;
        $this->datasource = $datasource;
        $this->editorMode = $editorMode;
        $this->exemplar = $exemplar;
        $this->expr = $expr ?: "";
        $this->format = $format;
        $this->groupByKeys = $groupByKeys;
        $this->hide = $hide;
        $this->instant = $instant;
        $this->intervalFactor = $intervalFactor;
        $this->intervalMs = $intervalMs;
        $this->legendFormat = $legendFormat;
        $this->maxDataPoints = $maxDataPoints;
        $this->queryType = $queryType;
        $this->range = $range;
        $this->refId = $refId;
        $this->resultAssertions = $resultAssertions;
        $this->scopes = $scopes;
        $this->timeRange = $timeRange;
        $this->interval = $interval;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{adhocFilters?: array<mixed>, datasource?: mixed, editorMode?: string, exemplar?: bool, expr?: string, format?: string, groupByKeys?: array<string>, hide?: bool, instant?: bool, intervalFactor?: int, intervalMs?: float, legendFormat?: string, maxDataPoints?: int, queryType?: string, range?: bool, refId?: string, resultAssertions?: mixed, scopes?: array<mixed>, timeRange?: mixed, interval?: string} $inputData */
        $data = $inputData;
        return new self(
            adhocFilters: array_filter(array_map((function($input) {
    	/** @var array{key?: string, operator?: string, value?: string, values?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Prometheus\AdhocFilters::fromArray($val);
    }), $data["adhocFilters"] ?? [])),
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            editorMode: isset($data["editorMode"]) ? (function($input) { return \Grafana\Foundation\Prometheus\QueryEditorMode::fromValue($input); })($data["editorMode"]) : null,
            exemplar: $data["exemplar"] ?? null,
            expr: $data["expr"] ?? null,
            format: isset($data["format"]) ? (function($input) { return \Grafana\Foundation\Prometheus\PromQueryFormat::fromValue($input); })($data["format"]) : null,
            groupByKeys: $data["groupByKeys"] ?? null,
            hide: $data["hide"] ?? null,
            instant: $data["instant"] ?? null,
            intervalFactor: $data["intervalFactor"] ?? null,
            intervalMs: $data["intervalMs"] ?? null,
            legendFormat: $data["legendFormat"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            queryType: $data["queryType"] ?? null,
            range: $data["range"] ?? null,
            refId: $data["refId"] ?? null,
            resultAssertions: isset($data["resultAssertions"]) ? (function($input) {
    	/** @var array{maxFrames?: int, type?: string, typeVersion?: array<int>} */
    $val = $input;
    	return \Grafana\Foundation\Prometheus\ResultAssertions::fromArray($val);
    })($data["resultAssertions"]) : null,
            scopes: array_filter(array_map((function($input) {
    	/** @var array{defaultPath?: array<string>, filters?: array<mixed>, title?: string} */
    $val = $input;
    	return \Grafana\Foundation\Prometheus\Scopes::fromArray($val);
    }), $data["scopes"] ?? [])),
            timeRange: isset($data["timeRange"]) ? (function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Prometheus\TimeRange::fromArray($val);
    })($data["timeRange"]) : null,
            interval: $data["interval"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->expr = $this->expr;
        if (isset($this->adhocFilters)) {
            $data->adhocFilters = $this->adhocFilters;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        if (isset($this->editorMode)) {
            $data->editorMode = $this->editorMode;
        }
        if (isset($this->exemplar)) {
            $data->exemplar = $this->exemplar;
        }
        if (isset($this->format)) {
            $data->format = $this->format;
        }
        if (isset($this->groupByKeys)) {
            $data->groupByKeys = $this->groupByKeys;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        if (isset($this->instant)) {
            $data->instant = $this->instant;
        }
        if (isset($this->intervalFactor)) {
            $data->intervalFactor = $this->intervalFactor;
        }
        if (isset($this->intervalMs)) {
            $data->intervalMs = $this->intervalMs;
        }
        if (isset($this->legendFormat)) {
            $data->legendFormat = $this->legendFormat;
        }
        if (isset($this->maxDataPoints)) {
            $data->maxDataPoints = $this->maxDataPoints;
        }
        if (isset($this->queryType)) {
            $data->queryType = $this->queryType;
        }
        if (isset($this->range)) {
            $data->range = $this->range;
        }
        if (isset($this->refId)) {
            $data->refId = $this->refId;
        }
        if (isset($this->resultAssertions)) {
            $data->resultAssertions = $this->resultAssertions;
        }
        if (isset($this->scopes)) {
            $data->scopes = $this->scopes;
        }
        if (isset($this->timeRange)) {
            $data->timeRange = $this->timeRange;
        }
        if (isset($this->interval)) {
            $data->interval = $this->interval;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return 'prometheus';
    }
}
