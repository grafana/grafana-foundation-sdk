<?php

namespace Grafana\Foundation\Elasticsearch;

class MetricAggregationWithMissingSupport implements \JsonSerializable
{
    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $settings;

    /**
     * @var string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType
     */
    public $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings|null $settings
     * @param string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $settings = null,  $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->settings = $settings;
        $this->type = $type ?: "count";
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{settings?: mixed, type?: string|string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{missing?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings::fromArray($val);
    })($data["settings"]) : null,
            type: isset($data["type"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            /** @var string $input */
            return \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType::fromValue($input);
    }
    })($data["type"]) : null,
            id: $data["id"] ?? null,
            hide: $data["hide"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "id" => $this->id,
        ];
        if (isset($this->settings)) {
            $data["settings"] = $this->settings;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
