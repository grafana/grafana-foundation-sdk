<?php

namespace Grafana\Foundation\Prometheus;

class AdhocFilters implements \JsonSerializable
{
    public string $key;

    public string $operator;

    public string $value;

    /**
     * @var array<string>|null
     */
    public ?array $values;

    /**
     * @param string|null $key
     * @param string|null $operator
     * @param string|null $value
     * @param array<string>|null $values
     */
    public function __construct(?string $key = null, ?string $operator = null, ?string $value = null, ?array $values = null)
    {
        $this->key = $key ?: "";
        $this->operator = $operator ?: "";
        $this->value = $value ?: "";
        $this->values = $values;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{key?: string, operator?: string, value?: string, values?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            key: $data["key"] ?? null,
            operator: $data["operator"] ?? null,
            value: $data["value"] ?? null,
            values: $data["values"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->key = $this->key;
        $data->operator = $this->operator;
        $data->value = $this->value;
        if (isset($this->values)) {
            $data->values = $this->values;
        }
        return $data;
    }
}
