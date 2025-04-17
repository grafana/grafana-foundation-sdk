<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Query filter representation.
 */
class Filter implements \JsonSerializable
{
    /**
     * Filter key.
     */
    public string $key;

    /**
     * Filter operator.
     */
    public string $operator;

    /**
     * Filter value.
     */
    public string $value;

    /**
     * Filter condition.
     */
    public ?string $condition;

    /**
     * @param string|null $key
     * @param string|null $operator
     * @param string|null $value
     * @param string|null $condition
     */
    public function __construct(?string $key = null, ?string $operator = null, ?string $value = null, ?string $condition = null)
    {
        $this->key = $key ?: "";
        $this->operator = $operator ?: "";
        $this->value = $value ?: "";
        $this->condition = $condition;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{key?: string, operator?: string, value?: string, condition?: string} $inputData */
        $data = $inputData;
        return new self(
            key: $data["key"] ?? null,
            operator: $data["operator"] ?? null,
            value: $data["value"] ?? null,
            condition: $data["condition"] ?? null,
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
        if (isset($this->condition)) {
            $data->condition = $this->condition;
        }
        return $data;
    }
}
