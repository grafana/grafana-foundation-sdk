<?php

namespace Grafana\Foundation\Elasticsearch;

class BaseMetricAggregation implements \JsonSerializable
{
    /**
     * @var string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType
     */
    public $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType|null $type
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct( $type = null, ?string $id = null, ?bool $hide = null)
    {
        $this->type = $type ?: "count";
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string|string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
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
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
