<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingVariableSpec implements \JsonSerializable
{
    public string $variable;

    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator $operator;

    public string $value;

    /**
     * @param string|null $variable
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator|null $operator
     * @param string|null $value
     */
    public function __construct(?string $variable = null, ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator $operator = null, ?string $value = null)
    {
        $this->variable = $variable ?: "";
        $this->operator = $operator ?: \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator::Equals();
        $this->value = $value ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{variable?: string, operator?: string, value?: string} $inputData */
        $data = $inputData;
        return new self(
            variable: $data["variable"] ?? null,
            operator: isset($data["operator"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpecOperator::fromValue($input); })($data["operator"]) : null,
            value: $data["value"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->variable = $this->variable;
        $data->operator = $this->operator;
        $data->value = $this->value;
        return $data;
    }
}
