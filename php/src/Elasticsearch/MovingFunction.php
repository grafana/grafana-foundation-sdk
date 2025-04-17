<?php

namespace Grafana\Foundation\Elasticsearch;

class MovingFunction implements \JsonSerializable
{
    public ?string $pipelineAgg;

    public ?string $field;

    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $pipelineAgg
     * @param string|null $field
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $pipelineAgg = null, ?string $field = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettings $settings = null, ?bool $hide = null)
    {
        $this->pipelineAgg = $pipelineAgg;
        $this->field = $field;
        $this->type = \Grafana\Foundation\Elasticsearch\MetricAggregationType::count();
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pipelineAgg?: string, field?: string, type?: "moving_fn", id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            pipelineAgg: $data["pipelineAgg"] ?? null,
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{window?: string, script?: string|mixed, shift?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchMovingFunctionSettings::fromArray($val);
    })($data["settings"]) : null,
            hide: $data["hide"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->id = $this->id;
        if (isset($this->pipelineAgg)) {
            $data->pipelineAgg = $this->pipelineAgg;
        }
        if (isset($this->field)) {
            $data->field = $this->field;
        }
        if (isset($this->settings)) {
            $data->settings = $this->settings;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        return $data;
    }
}
