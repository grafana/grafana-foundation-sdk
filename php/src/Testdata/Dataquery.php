<?php

namespace Grafana\Foundation\Testdata;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public ?string $alias;

    public ?\Grafana\Foundation\Testdata\TestDataQueryType $scenarioId;

    public ?string $stringInput;

    public ?\Grafana\Foundation\Testdata\StreamingQuery $stream;

    public ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave;

    public ?\Grafana\Foundation\Testdata\SimulationQuery $sim;

    /**
     * @var array<\Grafana\Foundation\Testdata\CSVWave>|null
     */
    public ?array $csvWave;

    public ?string $labels;

    public ?int $lines;

    public ?bool $levelColumn;

    public ?string $channel;

    public ?\Grafana\Foundation\Testdata\NodesQuery $nodes;

    public ?string $csvFileName;

    public ?string $csvContent;

    public ?string $rawFrameContent;

    public ?int $seriesCount;

    public ?\Grafana\Foundation\Testdata\USAQuery $usa;

    public ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType;

    public ?int $spanCount;

    /**
     * @var array<array<string|int>>|null
     */
    public ?array $points;

    /**
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public ?float $dropPercent;

    public ?bool $flamegraphDiff;

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
     * @param \Grafana\Foundation\Testdata\TestDataQueryType|null $scenarioId
     * @param string|null $stringInput
     * @param \Grafana\Foundation\Testdata\StreamingQuery|null $stream
     * @param \Grafana\Foundation\Testdata\PulseWaveQuery|null $pulseWave
     * @param \Grafana\Foundation\Testdata\SimulationQuery|null $sim
     * @param array<\Grafana\Foundation\Testdata\CSVWave>|null $csvWave
     * @param string|null $labels
     * @param int|null $lines
     * @param bool|null $levelColumn
     * @param string|null $channel
     * @param \Grafana\Foundation\Testdata\NodesQuery|null $nodes
     * @param string|null $csvFileName
     * @param string|null $csvContent
     * @param string|null $rawFrameContent
     * @param int|null $seriesCount
     * @param \Grafana\Foundation\Testdata\USAQuery|null $usa
     * @param \Grafana\Foundation\Testdata\DataqueryErrorType|null $errorType
     * @param int|null $spanCount
     * @param array<array<string|int>>|null $points
     * @param float|null $dropPercent
     * @param bool|null $flamegraphDiff
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?string $alias = null, ?\Grafana\Foundation\Testdata\TestDataQueryType $scenarioId = null, ?string $stringInput = null, ?\Grafana\Foundation\Testdata\StreamingQuery $stream = null, ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave = null, ?\Grafana\Foundation\Testdata\SimulationQuery $sim = null, ?array $csvWave = null, ?string $labels = null, ?int $lines = null, ?bool $levelColumn = null, ?string $channel = null, ?\Grafana\Foundation\Testdata\NodesQuery $nodes = null, ?string $csvFileName = null, ?string $csvContent = null, ?string $rawFrameContent = null, ?int $seriesCount = null, ?\Grafana\Foundation\Testdata\USAQuery $usa = null, ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType = null, ?int $spanCount = null, ?array $points = null, ?float $dropPercent = null, ?bool $flamegraphDiff = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->alias = $alias;
        $this->scenarioId = $scenarioId;
        $this->stringInput = $stringInput;
        $this->stream = $stream;
        $this->pulseWave = $pulseWave;
        $this->sim = $sim;
        $this->csvWave = $csvWave;
        $this->labels = $labels;
        $this->lines = $lines;
        $this->levelColumn = $levelColumn;
        $this->channel = $channel;
        $this->nodes = $nodes;
        $this->csvFileName = $csvFileName;
        $this->csvContent = $csvContent;
        $this->rawFrameContent = $rawFrameContent;
        $this->seriesCount = $seriesCount;
        $this->usa = $usa;
        $this->errorType = $errorType;
        $this->spanCount = $spanCount;
        $this->points = $points;
        $this->dropPercent = $dropPercent;
        $this->flamegraphDiff = $flamegraphDiff;
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
        /** @var array{alias?: string, scenarioId?: string, stringInput?: string, stream?: mixed, pulseWave?: mixed, sim?: mixed, csvWave?: array<mixed>, labels?: string, lines?: int, levelColumn?: bool, channel?: string, nodes?: mixed, csvFileName?: string, csvContent?: string, rawFrameContent?: string, seriesCount?: int, usa?: mixed, errorType?: string, spanCount?: int, points?: array<array<string|int>>, dropPercent?: float, flamegraphDiff?: bool, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            alias: $data["alias"] ?? null,
            scenarioId: isset($data["scenarioId"]) ? (function($input) { return \Grafana\Foundation\Testdata\TestDataQueryType::fromValue($input); })($data["scenarioId"]) : null,
            stringInput: $data["stringInput"] ?? null,
            stream: isset($data["stream"]) ? (function($input) {
    	/** @var array{type?: string, speed?: int, spread?: int, noise?: int, bands?: int, url?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\StreamingQuery::fromArray($val);
    })($data["stream"]) : null,
            pulseWave: isset($data["pulseWave"]) ? (function($input) {
    	/** @var array{timeStep?: int, onCount?: int, offCount?: int, onValue?: float, offValue?: float} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\PulseWaveQuery::fromArray($val);
    })($data["pulseWave"]) : null,
            sim: isset($data["sim"]) ? (function($input) {
    	/** @var array{key?: mixed, config?: array<string, mixed>, stream?: bool, last?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\SimulationQuery::fromArray($val);
    })($data["sim"]) : null,
            csvWave: array_filter(array_map((function($input) {
    	/** @var array{timeStep?: int, name?: string, valuesCSV?: string, labels?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\CSVWave::fromArray($val);
    }), $data["csvWave"] ?? [])),
            labels: $data["labels"] ?? null,
            lines: $data["lines"] ?? null,
            levelColumn: $data["levelColumn"] ?? null,
            channel: $data["channel"] ?? null,
            nodes: isset($data["nodes"]) ? (function($input) {
    	/** @var array{type?: string, count?: int} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\NodesQuery::fromArray($val);
    })($data["nodes"]) : null,
            csvFileName: $data["csvFileName"] ?? null,
            csvContent: $data["csvContent"] ?? null,
            rawFrameContent: $data["rawFrameContent"] ?? null,
            seriesCount: $data["seriesCount"] ?? null,
            usa: isset($data["usa"]) ? (function($input) {
    	/** @var array{mode?: string, period?: string, fields?: array<string>, states?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\USAQuery::fromArray($val);
    })($data["usa"]) : null,
            errorType: isset($data["errorType"]) ? (function($input) { return \Grafana\Foundation\Testdata\DataqueryErrorType::fromValue($input); })($data["errorType"]) : null,
            spanCount: $data["spanCount"] ?? null,
            points: $data["points"] ?? null,
            dropPercent: $data["dropPercent"] ?? null,
            flamegraphDiff: $data["flamegraphDiff"] ?? null,
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
        if (isset($this->scenarioId)) {
            $data["scenarioId"] = $this->scenarioId;
        }
        if (isset($this->stringInput)) {
            $data["stringInput"] = $this->stringInput;
        }
        if (isset($this->stream)) {
            $data["stream"] = $this->stream;
        }
        if (isset($this->pulseWave)) {
            $data["pulseWave"] = $this->pulseWave;
        }
        if (isset($this->sim)) {
            $data["sim"] = $this->sim;
        }
        if (isset($this->csvWave)) {
            $data["csvWave"] = $this->csvWave;
        }
        if (isset($this->labels)) {
            $data["labels"] = $this->labels;
        }
        if (isset($this->lines)) {
            $data["lines"] = $this->lines;
        }
        if (isset($this->levelColumn)) {
            $data["levelColumn"] = $this->levelColumn;
        }
        if (isset($this->channel)) {
            $data["channel"] = $this->channel;
        }
        if (isset($this->nodes)) {
            $data["nodes"] = $this->nodes;
        }
        if (isset($this->csvFileName)) {
            $data["csvFileName"] = $this->csvFileName;
        }
        if (isset($this->csvContent)) {
            $data["csvContent"] = $this->csvContent;
        }
        if (isset($this->rawFrameContent)) {
            $data["rawFrameContent"] = $this->rawFrameContent;
        }
        if (isset($this->seriesCount)) {
            $data["seriesCount"] = $this->seriesCount;
        }
        if (isset($this->usa)) {
            $data["usa"] = $this->usa;
        }
        if (isset($this->errorType)) {
            $data["errorType"] = $this->errorType;
        }
        if (isset($this->spanCount)) {
            $data["spanCount"] = $this->spanCount;
        }
        if (isset($this->points)) {
            $data["points"] = $this->points;
        }
        if (isset($this->dropPercent)) {
            $data["dropPercent"] = $this->dropPercent;
        }
        if (isset($this->flamegraphDiff)) {
            $data["flamegraphDiff"] = $this->flamegraphDiff;
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
        return "testdata";
    }
}
