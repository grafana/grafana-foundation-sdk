<?php

namespace Grafana\Foundation\Testdata;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public ?string $alias;

    /**
     * Used for live query
     */
    public ?string $channel;

    public ?string $csvContent;

    public ?string $csvFileName;

    /**
     * @var array<\Grafana\Foundation\Testdata\CSVWave>|null
     */
    public ?array $csvWave;

    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public ?float $dropPercent;

    /**
     * Possible enum values:
     *  - `"plugin"` 
     *  - `"downstream"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource;

    /**
     * Possible enum values:
     *  - `"frontend_exception"` 
     *  - `"frontend_observable"` 
     *  - `"server_panic"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType;

    public ?bool $flamegraphDiff;

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

    public ?string $labels;

    public ?bool $levelColumn;

    public ?int $lines;

    public ?float $max;

    /**
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public ?int $maxDataPoints;

    public ?float $min;

    public ?\Grafana\Foundation\Testdata\NodesQuery $nodes;

    public ?float $noise;

    /**
     * @var array<array<mixed>>|null
     */
    public ?array $points;

    public ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    public ?string $rawFrameContent;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Testdata\ResultAssertions $resultAssertions;

    /**
     * Possible enum values:
     *  - `"annotations"` 
     *  - `"arrow"` 
     *  - `"csv_content"` 
     *  - `"csv_file"` 
     *  - `"csv_metric_values"` 
     *  - `"datapoints_outside_range"` 
     *  - `"error_with_source"` 
     *  - `"exponential_heatmap_bucket_data"` 
     *  - `"flame_graph"` 
     *  - `"grafana_api"` 
     *  - `"linear_heatmap_bucket_data"` 
     *  - `"live"` 
     *  - `"logs"` 
     *  - `"manual_entry"` 
     *  - `"no_data_points"` 
     *  - `"node_graph"` 
     *  - `"predictable_csv_wave"` 
     *  - `"predictable_pulse"` 
     *  - `"random_walk"` 
     *  - `"random_walk_table"` 
     *  - `"random_walk_with_error"` 
     *  - `"raw_frame"` 
     *  - `"server_error_500"` 
     *  - `"simulation"` 
     *  - `"slow_query"` 
     *  - `"streaming_client"` 
     *  - `"table_static"` 
     *  - `"trace"` 
     *  - `"usa"` 
     *  - `"variables-query"` 
     */
    public ?\Grafana\Foundation\Testdata\DataqueryScenarioId $scenarioId;

    public ?int $seriesCount;

    public ?\Grafana\Foundation\Testdata\SimulationQuery $sim;

    public ?int $spanCount;

    public ?float $spread;

    public ?float $startValue;

    public ?\Grafana\Foundation\Testdata\StreamingQuery $stream;

    /**
     * common parameter used by many query types
     */
    public ?string $stringInput;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Testdata\TimeRange $timeRange;

    public ?\Grafana\Foundation\Testdata\USAQuery $usa;

    public ?bool $withNil;

    /**
     * @param string|null $alias
     * @param string|null $channel
     * @param string|null $csvContent
     * @param string|null $csvFileName
     * @param array<\Grafana\Foundation\Testdata\CSVWave>|null $csvWave
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param float|null $dropPercent
     * @param \Grafana\Foundation\Testdata\DataqueryErrorSource|null $errorSource
     * @param \Grafana\Foundation\Testdata\DataqueryErrorType|null $errorType
     * @param bool|null $flamegraphDiff
     * @param bool|null $hide
     * @param float|null $intervalMs
     * @param string|null $labels
     * @param bool|null $levelColumn
     * @param int|null $lines
     * @param float|null $max
     * @param int|null $maxDataPoints
     * @param float|null $min
     * @param \Grafana\Foundation\Testdata\NodesQuery|null $nodes
     * @param float|null $noise
     * @param array<array<mixed>>|null $points
     * @param \Grafana\Foundation\Testdata\PulseWaveQuery|null $pulseWave
     * @param string|null $queryType
     * @param string|null $rawFrameContent
     * @param string|null $refId
     * @param \Grafana\Foundation\Testdata\ResultAssertions|null $resultAssertions
     * @param \Grafana\Foundation\Testdata\DataqueryScenarioId|null $scenarioId
     * @param int|null $seriesCount
     * @param \Grafana\Foundation\Testdata\SimulationQuery|null $sim
     * @param int|null $spanCount
     * @param float|null $spread
     * @param float|null $startValue
     * @param \Grafana\Foundation\Testdata\StreamingQuery|null $stream
     * @param string|null $stringInput
     * @param \Grafana\Foundation\Testdata\TimeRange|null $timeRange
     * @param \Grafana\Foundation\Testdata\USAQuery|null $usa
     * @param bool|null $withNil
     */
    public function __construct(?string $alias = null, ?string $channel = null, ?string $csvContent = null, ?string $csvFileName = null, ?array $csvWave = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?float $dropPercent = null, ?\Grafana\Foundation\Testdata\DataqueryErrorSource $errorSource = null, ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType = null, ?bool $flamegraphDiff = null, ?bool $hide = null, ?float $intervalMs = null, ?string $labels = null, ?bool $levelColumn = null, ?int $lines = null, ?float $max = null, ?int $maxDataPoints = null, ?float $min = null, ?\Grafana\Foundation\Testdata\NodesQuery $nodes = null, ?float $noise = null, ?array $points = null, ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave = null, ?string $queryType = null, ?string $rawFrameContent = null, ?string $refId = null, ?\Grafana\Foundation\Testdata\ResultAssertions $resultAssertions = null, ?\Grafana\Foundation\Testdata\DataqueryScenarioId $scenarioId = null, ?int $seriesCount = null, ?\Grafana\Foundation\Testdata\SimulationQuery $sim = null, ?int $spanCount = null, ?float $spread = null, ?float $startValue = null, ?\Grafana\Foundation\Testdata\StreamingQuery $stream = null, ?string $stringInput = null, ?\Grafana\Foundation\Testdata\TimeRange $timeRange = null, ?\Grafana\Foundation\Testdata\USAQuery $usa = null, ?bool $withNil = null)
    {
        $this->alias = $alias;
        $this->channel = $channel;
        $this->csvContent = $csvContent;
        $this->csvFileName = $csvFileName;
        $this->csvWave = $csvWave;
        $this->datasource = $datasource;
        $this->dropPercent = $dropPercent;
        $this->errorSource = $errorSource;
        $this->errorType = $errorType;
        $this->flamegraphDiff = $flamegraphDiff;
        $this->hide = $hide;
        $this->intervalMs = $intervalMs;
        $this->labels = $labels;
        $this->levelColumn = $levelColumn;
        $this->lines = $lines;
        $this->max = $max;
        $this->maxDataPoints = $maxDataPoints;
        $this->min = $min;
        $this->nodes = $nodes;
        $this->noise = $noise;
        $this->points = $points;
        $this->pulseWave = $pulseWave;
        $this->queryType = $queryType;
        $this->rawFrameContent = $rawFrameContent;
        $this->refId = $refId;
        $this->resultAssertions = $resultAssertions;
        $this->scenarioId = $scenarioId;
        $this->seriesCount = $seriesCount;
        $this->sim = $sim;
        $this->spanCount = $spanCount;
        $this->spread = $spread;
        $this->startValue = $startValue;
        $this->stream = $stream;
        $this->stringInput = $stringInput;
        $this->timeRange = $timeRange;
        $this->usa = $usa;
        $this->withNil = $withNil;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{alias?: string, channel?: string, csvContent?: string, csvFileName?: string, csvWave?: array<mixed>, datasource?: mixed, dropPercent?: float, errorSource?: string, errorType?: string, flamegraphDiff?: bool, hide?: bool, intervalMs?: float, labels?: string, levelColumn?: bool, lines?: int, max?: float, maxDataPoints?: int, min?: float, nodes?: mixed, noise?: float, points?: array<array<mixed>>, pulseWave?: mixed, queryType?: string, rawFrameContent?: string, refId?: string, resultAssertions?: mixed, scenarioId?: string, seriesCount?: int, sim?: mixed, spanCount?: int, spread?: float, startValue?: float, stream?: mixed, stringInput?: string, timeRange?: mixed, usa?: mixed, withNil?: bool} $inputData */
        $data = $inputData;
        return new self(
            alias: $data["alias"] ?? null,
            channel: $data["channel"] ?? null,
            csvContent: $data["csvContent"] ?? null,
            csvFileName: $data["csvFileName"] ?? null,
            csvWave: array_filter(array_map((function($input) {
    	/** @var array{labels?: string, name?: string, timeStep?: int, valuesCSV?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\CSVWave::fromArray($val);
    }), $data["csvWave"] ?? [])),
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            dropPercent: $data["dropPercent"] ?? null,
            errorSource: isset($data["errorSource"]) ? (function($input) { return \Grafana\Foundation\Testdata\DataqueryErrorSource::fromValue($input); })($data["errorSource"]) : null,
            errorType: isset($data["errorType"]) ? (function($input) { return \Grafana\Foundation\Testdata\DataqueryErrorType::fromValue($input); })($data["errorType"]) : null,
            flamegraphDiff: $data["flamegraphDiff"] ?? null,
            hide: $data["hide"] ?? null,
            intervalMs: $data["intervalMs"] ?? null,
            labels: $data["labels"] ?? null,
            levelColumn: $data["levelColumn"] ?? null,
            lines: $data["lines"] ?? null,
            max: $data["max"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            min: $data["min"] ?? null,
            nodes: isset($data["nodes"]) ? (function($input) {
    	/** @var array{count?: int, seed?: int, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\NodesQuery::fromArray($val);
    })($data["nodes"]) : null,
            noise: $data["noise"] ?? null,
            points: $data["points"] ?? null,
            pulseWave: isset($data["pulseWave"]) ? (function($input) {
    	/** @var array{offCount?: int, offValue?: float, onCount?: int, onValue?: float, timeStep?: int} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\PulseWaveQuery::fromArray($val);
    })($data["pulseWave"]) : null,
            queryType: $data["queryType"] ?? null,
            rawFrameContent: $data["rawFrameContent"] ?? null,
            refId: $data["refId"] ?? null,
            resultAssertions: isset($data["resultAssertions"]) ? (function($input) {
    	/** @var array{maxFrames?: int, type?: string, typeVersion?: array<int>} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\ResultAssertions::fromArray($val);
    })($data["resultAssertions"]) : null,
            scenarioId: isset($data["scenarioId"]) ? (function($input) { return \Grafana\Foundation\Testdata\DataqueryScenarioId::fromValue($input); })($data["scenarioId"]) : null,
            seriesCount: $data["seriesCount"] ?? null,
            sim: isset($data["sim"]) ? (function($input) {
    	/** @var array{config?: mixed, key?: mixed, last?: bool, stream?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\SimulationQuery::fromArray($val);
    })($data["sim"]) : null,
            spanCount: $data["spanCount"] ?? null,
            spread: $data["spread"] ?? null,
            startValue: $data["startValue"] ?? null,
            stream: isset($data["stream"]) ? (function($input) {
    	/** @var array{bands?: int, noise?: float, speed?: float, spread?: float, type?: string, url?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\StreamingQuery::fromArray($val);
    })($data["stream"]) : null,
            stringInput: $data["stringInput"] ?? null,
            timeRange: isset($data["timeRange"]) ? (function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\TimeRange::fromArray($val);
    })($data["timeRange"]) : null,
            usa: isset($data["usa"]) ? (function($input) {
    	/** @var array{fields?: array<string>, mode?: string, period?: string, states?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\USAQuery::fromArray($val);
    })($data["usa"]) : null,
            withNil: $data["withNil"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->alias)) {
            $data["alias"] = $this->alias;
        }
        if (isset($this->channel)) {
            $data["channel"] = $this->channel;
        }
        if (isset($this->csvContent)) {
            $data["csvContent"] = $this->csvContent;
        }
        if (isset($this->csvFileName)) {
            $data["csvFileName"] = $this->csvFileName;
        }
        if (isset($this->csvWave)) {
            $data["csvWave"] = $this->csvWave;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->dropPercent)) {
            $data["dropPercent"] = $this->dropPercent;
        }
        if (isset($this->errorSource)) {
            $data["errorSource"] = $this->errorSource;
        }
        if (isset($this->errorType)) {
            $data["errorType"] = $this->errorType;
        }
        if (isset($this->flamegraphDiff)) {
            $data["flamegraphDiff"] = $this->flamegraphDiff;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->intervalMs)) {
            $data["intervalMs"] = $this->intervalMs;
        }
        if (isset($this->labels)) {
            $data["labels"] = $this->labels;
        }
        if (isset($this->levelColumn)) {
            $data["levelColumn"] = $this->levelColumn;
        }
        if (isset($this->lines)) {
            $data["lines"] = $this->lines;
        }
        if (isset($this->max)) {
            $data["max"] = $this->max;
        }
        if (isset($this->maxDataPoints)) {
            $data["maxDataPoints"] = $this->maxDataPoints;
        }
        if (isset($this->min)) {
            $data["min"] = $this->min;
        }
        if (isset($this->nodes)) {
            $data["nodes"] = $this->nodes;
        }
        if (isset($this->noise)) {
            $data["noise"] = $this->noise;
        }
        if (isset($this->points)) {
            $data["points"] = $this->points;
        }
        if (isset($this->pulseWave)) {
            $data["pulseWave"] = $this->pulseWave;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->rawFrameContent)) {
            $data["rawFrameContent"] = $this->rawFrameContent;
        }
        if (isset($this->refId)) {
            $data["refId"] = $this->refId;
        }
        if (isset($this->resultAssertions)) {
            $data["resultAssertions"] = $this->resultAssertions;
        }
        if (isset($this->scenarioId)) {
            $data["scenarioId"] = $this->scenarioId;
        }
        if (isset($this->seriesCount)) {
            $data["seriesCount"] = $this->seriesCount;
        }
        if (isset($this->sim)) {
            $data["sim"] = $this->sim;
        }
        if (isset($this->spanCount)) {
            $data["spanCount"] = $this->spanCount;
        }
        if (isset($this->spread)) {
            $data["spread"] = $this->spread;
        }
        if (isset($this->startValue)) {
            $data["startValue"] = $this->startValue;
        }
        if (isset($this->stream)) {
            $data["stream"] = $this->stream;
        }
        if (isset($this->stringInput)) {
            $data["stringInput"] = $this->stringInput;
        }
        if (isset($this->timeRange)) {
            $data["timeRange"] = $this->timeRange;
        }
        if (isset($this->usa)) {
            $data["usa"] = $this->usa;
        }
        if (isset($this->withNil)) {
            $data["withNil"] = $this->withNil;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "";
    }
}
