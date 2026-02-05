<?php

namespace Grafana\Foundation\Elasticsearch;

class BaseMetricAggregation implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param \Grafana\Foundation\Elasticsearch\MetricAggregationType|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\MetricAggregationType $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Elasticsearch\MetricAggregationType::Count();
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
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
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        return $data;
    }
}
