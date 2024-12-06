<?php

namespace Grafana\Foundation\Expr;

class TypeResample implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * The downsample function
     * Possible enum values:
     *  - `"sum"` 
     *  - `"mean"` 
     *  - `"min"` 
     *  - `"max"` 
     *  - `"count"` 
     *  - `"last"` 
     *  - `"median"` 
     */
    public \Grafana\Foundation\Expr\TypeResampleDownsampler $downsampler;

    /**
     * The math expression
     */
    public string $expression;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public ?float $intervalMs;

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
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Expr\ExprTypeResampleResultAssertions $resultAssertions;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Expr\ExprTypeResampleTimeRange $timeRange;

    public string $type;

    /**
     * The upsample function
     * Possible enum values:
     *  - `"pad"` Use the last seen value
     *  - `"backfilling"` backfill
     *  - `"fillna"` Do not fill values (nill)
     */
    public \Grafana\Foundation\Expr\TypeResampleUpsampler $upsampler;

    /**
     * The time duration
     */
    public string $window;

    /**
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param \Grafana\Foundation\Expr\TypeResampleDownsampler|null $downsampler
     * @param string|null $expression
     * @param bool|null $hide
     * @param float|null $intervalMs
     * @param int|null $maxDataPoints
     * @param string|null $queryType
     * @param string|null $refId
     * @param \Grafana\Foundation\Expr\ExprTypeResampleResultAssertions|null $resultAssertions
     * @param \Grafana\Foundation\Expr\ExprTypeResampleTimeRange|null $timeRange
     * @param \Grafana\Foundation\Expr\TypeResampleUpsampler|null $upsampler
     * @param string|null $window
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?\Grafana\Foundation\Expr\TypeResampleDownsampler $downsampler = null, ?string $expression = null, ?bool $hide = null, ?float $intervalMs = null, ?int $maxDataPoints = null, ?string $queryType = null, ?string $refId = null, ?\Grafana\Foundation\Expr\ExprTypeResampleResultAssertions $resultAssertions = null, ?\Grafana\Foundation\Expr\ExprTypeResampleTimeRange $timeRange = null, ?\Grafana\Foundation\Expr\TypeResampleUpsampler $upsampler = null, ?string $window = null)
    {
        $this->datasource = $datasource;
        $this->downsampler = $downsampler ?: \Grafana\Foundation\Expr\TypeResampleDownsampler::Sum();
        $this->expression = $expression ?: "";
        $this->hide = $hide;
        $this->intervalMs = $intervalMs;
        $this->maxDataPoints = $maxDataPoints;
        $this->queryType = $queryType;
        $this->refId = $refId ?: "";
        $this->resultAssertions = $resultAssertions;
        $this->timeRange = $timeRange;
        $this->type = "resample";
    
        $this->upsampler = $upsampler ?: \Grafana\Foundation\Expr\TypeResampleUpsampler::Pad();
        $this->window = $window ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{datasource?: mixed, downsampler?: string, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string, upsampler?: string, window?: string} $inputData */
        $data = $inputData;
        return new self(
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            downsampler: isset($data["downsampler"]) ? (function($input) { return \Grafana\Foundation\Expr\TypeResampleDownsampler::fromValue($input); })($data["downsampler"]) : null,
            expression: $data["expression"] ?? null,
            hide: $data["hide"] ?? null,
            intervalMs: $data["intervalMs"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            queryType: $data["queryType"] ?? null,
            refId: $data["refId"] ?? null,
            resultAssertions: isset($data["resultAssertions"]) ? (function($input) {
    	/** @var array{maxFrames?: int, type?: string, typeVersion?: array<int>} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeResampleResultAssertions::fromArray($val);
    })($data["resultAssertions"]) : null,
            timeRange: isset($data["timeRange"]) ? (function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeResampleTimeRange::fromArray($val);
    })($data["timeRange"]) : null,
            upsampler: isset($data["upsampler"]) ? (function($input) { return \Grafana\Foundation\Expr\TypeResampleUpsampler::fromValue($input); })($data["upsampler"]) : null,
            window: $data["window"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "downsampler" => $this->downsampler,
            "expression" => $this->expression,
            "refId" => $this->refId,
            "type" => $this->type,
            "upsampler" => $this->upsampler,
            "window" => $this->window,
        ];
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->intervalMs)) {
            $data["intervalMs"] = $this->intervalMs;
        }
        if (isset($this->maxDataPoints)) {
            $data["maxDataPoints"] = $this->maxDataPoints;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->resultAssertions)) {
            $data["resultAssertions"] = $this->resultAssertions;
        }
        if (isset($this->timeRange)) {
            $data["timeRange"] = $this->timeRange;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "__expr__";
    }
}
