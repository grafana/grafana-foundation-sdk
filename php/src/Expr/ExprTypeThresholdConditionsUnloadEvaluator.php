<?php

namespace Grafana\Foundation\Expr;

class ExprTypeThresholdConditionsUnloadEvaluator implements \JsonSerializable
{
    /**
     * @var array<float>
     */
    public array $params;

    /**
     * e.g. "gt"
     */
    public \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType $type;

    /**
     * @param array<float>|null $params
     * @param \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType|null $type
     */
    public function __construct(?array $params = null, ?\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType $type = null)
    {
        $this->params = $params ?: [];
        $this->type = $type ?: \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType::Gt();
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
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->params = $this->params;
        $data->type = $this->type;
        return $data;
    }
}
