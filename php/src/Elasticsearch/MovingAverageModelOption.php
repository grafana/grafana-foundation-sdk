<?php

namespace Grafana\Foundation\Elasticsearch;

class MovingAverageModelOption implements \JsonSerializable
{
    public string $label;

    public \Grafana\Foundation\Elasticsearch\MovingAverageModel $value;

    /**
     * @param string|null $label
     * @param \Grafana\Foundation\Elasticsearch\MovingAverageModel|null $value
     */
    public function __construct(?string $label = null, ?\Grafana\Foundation\Elasticsearch\MovingAverageModel $value = null)
    {
        $this->label = $label ?: "";
        $this->value = $value ?: \Grafana\Foundation\Elasticsearch\MovingAverageModel::Simple();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{label?: string, value?: string} $inputData */
        $data = $inputData;
        return new self(
            label: $data["label"] ?? null,
            value: isset($data["value"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\MovingAverageModel::fromValue($input); })($data["value"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "label" => $this->label,
            "value" => $this->value,
        ];
        return $data;
    }
}
