<?php

namespace Grafana\Foundation\Alertingv0alpha1;

class Expression implements \JsonSerializable
{
    /**
     * The type of query if this is a query expression
     */
    public ?string $queryType;

    public ?\Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange $relativeTimeRange;

    /**
     * The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
     */
    public string $datasourceUID;

    /**
     * @var mixed
     */
    public $model;

    /**
     * Used to mark the expression to be used as the final source for the rule evaluation
     * Only one expression in a rule can be marked as the source
     * For AlertRules, this is the expression that will be evaluated against the alerting condition
     * For RecordingRules, this is the expression that will be recorded
     */
    public ?bool $source;

    /**
     * @param string|null $queryType
     * @param \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange|null $relativeTimeRange
     * @param string|null $datasourceUID
     * @param mixed|null $model
     * @param bool|null $source
     */
    public function __construct(?string $queryType = null, ?\Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange $relativeTimeRange = null, ?string $datasourceUID = null,  $model = null, ?bool $source = null)
    {
        $this->queryType = $queryType;
        $this->relativeTimeRange = $relativeTimeRange;
        $this->datasourceUID = $datasourceUID ?: "";
        $this->model = $model ?: null;
        $this->source = $source;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{queryType?: string, relativeTimeRange?: mixed, datasourceUID?: string, model?: mixed, source?: bool} $inputData */
        $data = $inputData;
        return new self(
            queryType: $data["queryType"] ?? null,
            relativeTimeRange: isset($data["relativeTimeRange"]) ? (function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange::fromArray($val);
    })($data["relativeTimeRange"]) : null,
            datasourceUID: $data["datasourceUID"] ?? null,
            model: $data["model"] ?? null,
            source: $data["source"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->datasourceUID = $this->datasourceUID;
        $data->model = $this->model;
        if (isset($this->queryType)) {
            $data->queryType = $this->queryType;
        }
        if (isset($this->relativeTimeRange)) {
            $data->relativeTimeRange = $this->relativeTimeRange;
        }
        if (isset($this->source)) {
            $data->source = $this->source;
        }
        return $data;
    }
}
