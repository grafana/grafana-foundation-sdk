<?php

namespace Grafana\Foundation\Prometheus;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public string $expr;

    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public ?bool $instant;

    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public ?bool $range;

    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public ?bool $exemplar;

    /**
     * Specifies which editor is being used to prepare the query. It can be "code" or "builder"
     */
    public ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode;

    /**
     * Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
     */
    public ?\Grafana\Foundation\Prometheus\PromQueryFormat $format;

    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public ?string $legendFormat;

    /**
     * @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     */
    public ?float $intervalFactor;

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
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public ?string $interval;

    /**
     * @param string|null $expr
     * @param bool|null $instant
     * @param bool|null $range
     * @param bool|null $exemplar
     * @param \Grafana\Foundation\Prometheus\QueryEditorMode|null $editorMode
     * @param \Grafana\Foundation\Prometheus\PromQueryFormat|null $format
     * @param string|null $legendFormat
     * @param float|null $intervalFactor
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param string|null $interval
     */
    public function __construct(?string $expr = null, ?bool $instant = null, ?bool $range = null, ?bool $exemplar = null, ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode = null, ?\Grafana\Foundation\Prometheus\PromQueryFormat $format = null, ?string $legendFormat = null, ?float $intervalFactor = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?string $interval = null)
    {
        $this->expr = $expr ?: "";
        $this->instant = $instant;
        $this->range = $range;
        $this->exemplar = $exemplar;
        $this->editorMode = $editorMode;
        $this->format = $format;
        $this->legendFormat = $legendFormat;
        $this->intervalFactor = $intervalFactor;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->datasource = $datasource;
        $this->interval = $interval;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{expr?: string, instant?: bool, range?: bool, exemplar?: bool, editorMode?: string, format?: string, legendFormat?: string, intervalFactor?: float, refId?: string, hide?: bool, queryType?: string, datasource?: mixed, interval?: string} $inputData */
        $data = $inputData;
        return new self(
            expr: $data["expr"] ?? null,
            instant: $data["instant"] ?? null,
            range: $data["range"] ?? null,
            exemplar: $data["exemplar"] ?? null,
            editorMode: isset($data["editorMode"]) ? (function($input) { return \Grafana\Foundation\Prometheus\QueryEditorMode::fromValue($input); })($data["editorMode"]) : null,
            format: isset($data["format"]) ? (function($input) { return \Grafana\Foundation\Prometheus\PromQueryFormat::fromValue($input); })($data["format"]) : null,
            legendFormat: $data["legendFormat"] ?? null,
            intervalFactor: $data["intervalFactor"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            interval: $data["interval"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "expr" => $this->expr,
            "refId" => $this->refId,
        ];
        if (isset($this->instant)) {
            $data["instant"] = $this->instant;
        }
        if (isset($this->range)) {
            $data["range"] = $this->range;
        }
        if (isset($this->exemplar)) {
            $data["exemplar"] = $this->exemplar;
        }
        if (isset($this->editorMode)) {
            $data["editorMode"] = $this->editorMode;
        }
        if (isset($this->format)) {
            $data["format"] = $this->format;
        }
        if (isset($this->legendFormat)) {
            $data["legendFormat"] = $this->legendFormat;
        }
        if (isset($this->intervalFactor)) {
            $data["intervalFactor"] = $this->intervalFactor;
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
        if (isset($this->interval)) {
            $data["interval"] = $this->interval;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "prometheus";
    }
}
