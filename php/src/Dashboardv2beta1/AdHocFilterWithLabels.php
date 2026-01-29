<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Define the AdHocFilterWithLabels type
 */
class AdHocFilterWithLabels implements \JsonSerializable
{
    public string $key;

    public string $operator;

    public string $value;

    /**
     * @var array<string>|null
     */
    public ?array $values;

    public ?string $keyLabel;

    /**
     * @var array<string>|null
     */
    public ?array $valueLabels;

    public ?bool $forceEdit;

    public ?string $origin;

    /**
     * @deprecated
     */
    public ?string $condition;

    /**
     * @param string|null $key
     * @param string|null $operator
     * @param string|null $value
     * @param array<string>|null $values
     * @param string|null $keyLabel
     * @param array<string>|null $valueLabels
     * @param bool|null $forceEdit
     * @param string|null $condition
     */
    public function __construct(?string $key = null, ?string $operator = null, ?string $value = null, ?array $values = null, ?string $keyLabel = null, ?array $valueLabels = null, ?bool $forceEdit = null, ?string $condition = null)
    {
        $this->key = $key ?: "";
        $this->operator = $operator ?: "";
        $this->value = $value ?: "";
        $this->values = $values;
        $this->keyLabel = $keyLabel;
        $this->valueLabels = $valueLabels;
        $this->forceEdit = $forceEdit;
        $this->origin = \Grafana\Foundation\Dashboardv2beta1\Constants::FILTER_ORIGIN;
        $this->condition = $condition;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{key?: string, operator?: string, value?: string, values?: array<string>, keyLabel?: string, valueLabels?: array<string>, forceEdit?: bool, origin?: "dashboard", condition?: string} $inputData */
        $data = $inputData;
        return new self(
            key: $data["key"] ?? null,
            operator: $data["operator"] ?? null,
            value: $data["value"] ?? null,
            values: $data["values"] ?? null,
            keyLabel: $data["keyLabel"] ?? null,
            valueLabels: $data["valueLabels"] ?? null,
            forceEdit: $data["forceEdit"] ?? null,
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
        if (isset($this->values)) {
            $data->values = $this->values;
        }
        if (isset($this->keyLabel)) {
            $data->keyLabel = $this->keyLabel;
        }
        if (isset($this->valueLabels)) {
            $data->valueLabels = $this->valueLabels;
        }
        if (isset($this->forceEdit)) {
            $data->forceEdit = $this->forceEdit;
        }
        if (isset($this->origin)) {
            $data->origin = $this->origin;
        }
        if (isset($this->condition)) {
            $data->condition = $this->condition;
        }
        return $data;
    }
}
