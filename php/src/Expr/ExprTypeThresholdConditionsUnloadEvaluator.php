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
    public \Grafana\Foundation\Expr\TypeThresholdType $type;

    /**
     * @param array<float>|null $params
     * @param \Grafana\Foundation\Expr\TypeThresholdType|null $type
     */
    public function __construct(?array $params = null, ?\Grafana\Foundation\Expr\TypeThresholdType $type = null)
    {
        $this->params = $params ?: [];
        $this->type = $type ?: \Grafana\Foundation\Expr\TypeThresholdType::None();
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
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\TypeThresholdType::fromValue($input); })($data["type"]) : null,
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
