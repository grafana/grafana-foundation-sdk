<?php

namespace Grafana\Foundation\Expr;

class ExprTypeThresholdConditionsEvaluator implements \JsonSerializable
{
    /**
     * @var array<float>
     */
    public array $params;

    /**
     * e.g. "gt"
     */
    public \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorType $type;

    /**
     * @param array<float>|null $params
     * @param \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorType|null $type
     */
    public function __construct(?array $params = null, ?\Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorType $type = null)
    {
        $this->params = $params ?: [];
        $this->type = $type ?: \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorType::Gt();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{params?: array<float>, type?: string} $inputData */
        $data = $inputData;
        return new self(
            params: $data["params"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "params" => $this->params,
            "type" => $this->type,
        ];
        return $data;
    }
}
