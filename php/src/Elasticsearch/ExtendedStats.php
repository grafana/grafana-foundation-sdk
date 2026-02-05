<?php

namespace Grafana\Foundation\Elasticsearch;

class ExtendedStats implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings $settings;

    public ?string $field;

    public string $id;

    /**
     * @var mixed|null
     */
    public $meta;

    public ?bool $hide;

    /**
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings|null $settings
     * @param string|null $field
     * @param string|null $id
     * @param mixed|null $meta
     * @param bool|null $hide
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings $settings = null, ?string $field = null, ?string $id = null,  $meta = null, ?bool $hide = null)
    {
        $this->type = \Grafana\Foundation\Elasticsearch\MetricAggregationType::count();
        $this->settings = $settings;
        $this->field = $field;
        $this->id = $id ?: "";
        $this->meta = $meta;
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "extended_stats", settings?: mixed, field?: string, id?: string, meta?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{script?: string|mixed, missing?: string, sigma?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings::fromArray($val);
    })($data["settings"]) : null,
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            meta: $data["meta"] ?? null,
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
        if (isset($this->field)) {
            $data->field = $this->field;
        }
        if (isset($this->meta)) {
            $data->meta = $this->meta;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        return $data;
    }
}
