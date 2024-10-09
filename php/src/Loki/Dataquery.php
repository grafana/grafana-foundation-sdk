<?php

namespace Grafana\Foundation\Loki;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * The LogQL query.
     */
    public string $expr;

    /**
     * Used to override the name of the series.
     */
    public ?string $legendFormat;

    /**
     * Used to limit the number of log rows returned.
     */
    public ?int $maxLines;

    /**
     * @deprecated, now use step.
     */
    public ?int $resolution;

    public ?\Grafana\Foundation\Loki\QueryEditorMode $editorMode;

    /**
     * @deprecated, now use queryType.
     */
    public ?bool $range;

    /**
     * @deprecated, now use queryType.
     */
    public ?bool $instant;

    /**
     * Used to set step value for range queries.
     */
    public ?string $step;

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
     * @param string|null $expr
     * @param string|null $legendFormat
     * @param int|null $maxLines
     * @param int|null $resolution
     * @param \Grafana\Foundation\Loki\QueryEditorMode|null $editorMode
     * @param bool|null $range
     * @param bool|null $instant
     * @param string|null $step
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?string $expr = null, ?string $legendFormat = null, ?int $maxLines = null, ?int $resolution = null, ?\Grafana\Foundation\Loki\QueryEditorMode $editorMode = null, ?bool $range = null, ?bool $instant = null, ?string $step = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->expr = $expr ?: "";
        $this->legendFormat = $legendFormat;
        $this->maxLines = $maxLines;
        $this->resolution = $resolution;
        $this->editorMode = $editorMode;
        $this->range = $range;
        $this->instant = $instant;
        $this->step = $step;
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
        /** @var array{expr?: string, legendFormat?: string, maxLines?: int, resolution?: int, editorMode?: string, range?: bool, instant?: bool, step?: string, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            expr: $data["expr"] ?? null,
            legendFormat: $data["legendFormat"] ?? null,
            maxLines: $data["maxLines"] ?? null,
            resolution: $data["resolution"] ?? null,
            editorMode: isset($data["editorMode"]) ? (function($input) { return \Grafana\Foundation\Loki\QueryEditorMode::fromValue($input); })($data["editorMode"]) : null,
            range: $data["range"] ?? null,
            instant: $data["instant"] ?? null,
            step: $data["step"] ?? null,
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
            "expr" => $this->expr,
            "refId" => $this->refId,
        ];
        if (isset($this->legendFormat)) {
            $data["legendFormat"] = $this->legendFormat;
        }
        if (isset($this->maxLines)) {
            $data["maxLines"] = $this->maxLines;
        }
        if (isset($this->resolution)) {
            $data["resolution"] = $this->resolution;
        }
        if (isset($this->editorMode)) {
            $data["editorMode"] = $this->editorMode;
        }
        if (isset($this->range)) {
            $data["range"] = $this->range;
        }
        if (isset($this->instant)) {
            $data["instant"] = $this->instant;
        }
        if (isset($this->step)) {
            $data["step"] = $this->step;
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
        return "loki";
    }
}
