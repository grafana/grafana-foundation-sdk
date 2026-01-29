<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * FIXME: should we introduce this? --- Variable value option
 */
class VariableValueOption implements \JsonSerializable
{
    public string $label;

    /**
     * @var string|bool|float|\Grafana\Foundation\Dashboardv2beta1\CustomVariableValue
     */
    public $value;

    public ?string $group;

    /**
     * @param string|null $label
     * @param string|bool|float|\Grafana\Foundation\Dashboardv2beta1\CustomVariableValue|null $value
     * @param string|null $group
     */
    public function __construct(?string $label = null,  $value = null, ?string $group = null)
    {
        $this->label = $label ?: "";
        $this->value = $value ?: "";
        $this->group = $group;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{label?: string, value?: string|bool|float|mixed, group?: string} $inputData */
        $data = $inputData;
        return new self(
            label: $data["label"] ?? null,
            value: isset($data["value"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        case is_bool($input):
            return $input;
        case is_float($input):
            return $input;
        default:
            /** @var array{formatter?: string} $input */
            return \Grafana\Foundation\Dashboardv2beta1\CustomVariableValue::fromArray($input);
    }
    })($data["value"]) : null,
            group: $data["group"] ?? null,
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
        if (isset($this->group)) {
            $data->group = $this->group;
        }
        return $data;
    }
}
