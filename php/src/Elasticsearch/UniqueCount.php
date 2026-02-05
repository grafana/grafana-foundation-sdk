<?php

namespace Grafana\Foundation\Elasticsearch;

class UniqueCount implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public ?string $field;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $field
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $field = null, ?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings $settings = null, ?bool $hide = null)
    {
        $this->type = \Grafana\Foundation\Elasticsearch\MetricAggregationType::count();
        $this->field = $field;
        $this->id = $id ?: "";
        $this->settings = $settings;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "cardinality", field?: string, id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{precision_threshold?: string, missing?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchUniqueCountSettings::fromArray($val);
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
