<?php

namespace Grafana\Foundation\Elasticsearch;

class MetricAggregationWithMissingSupport implements \JsonSerializable
{
    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $settings;

    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings|null $settings
     * @param \Grafana\Foundation\Elasticsearch\MetricAggregationType|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $settings = null, ?\Grafana\Foundation\Elasticsearch\MetricAggregationType $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->settings = $settings;
        $this->type = $type ?: \Grafana\Foundation\Elasticsearch\MetricAggregationType::Count();
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{settings?: mixed, type?: string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{missing?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings::fromArray($val);
    })($data["settings"]) : null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\MetricAggregationType::fromValue($input); })($data["type"]) : null,
            id: $data["id"] ?? null,
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
        if (isset($this->settings)) {
            $data->settings = $this->settings;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        return $data;
    }
}
