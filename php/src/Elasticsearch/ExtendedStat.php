<?php

namespace Grafana\Foundation\Elasticsearch;

class ExtendedStat implements \JsonSerializable
{
    public string $label;

    public \Grafana\Foundation\Elasticsearch\ExtendedStatMetaType $value;

    /**
     * @param string|null $label
     * @param \Grafana\Foundation\Elasticsearch\ExtendedStatMetaType|null $value
     */
    public function __construct(?string $label = null, ?\Grafana\Foundation\Elasticsearch\ExtendedStatMetaType $value = null)
    {
        $this->label = $label ?: "";
        $this->value = $value ?: \Grafana\Foundation\Elasticsearch\ExtendedStatMetaType::Avg();
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
            value: isset($data["value"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\ExtendedStatMetaType::fromValue($input); })($data["value"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->label = $this->label;
        $data->value = $this->value;
        return $data;
    }
}
