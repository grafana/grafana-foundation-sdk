<?php

namespace Grafana\Foundation\Elasticsearch;

class Count implements \JsonSerializable
{
    public \Grafana\Foundation\Elasticsearch\MetricAggregationType $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?string $id = null, ?bool $hide = null)
    {
        $this->type = \Grafana\Foundation\Elasticsearch\MetricAggregationType::count();
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "count", id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
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
