<?php

namespace Grafana\Foundation\Elasticsearch;

class RawData implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings $settings;

    public ?bool $hide;

    /**
     * @param string|null $id
     * @param \Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings|null $settings
     * @param bool|null $hide
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings $settings = null, ?bool $hide = null)
    {
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
        /** @var array{type?: "raw_data", id?: string, settings?: mixed, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            settings: isset($data["settings"]) ? (function($input) {
    	/** @var array{size?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\ElasticsearchRawDataSettings::fromArray($val);
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
        if (isset($this->settings)) {
            $data->settings = $this->settings;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        return $data;
    }
}
