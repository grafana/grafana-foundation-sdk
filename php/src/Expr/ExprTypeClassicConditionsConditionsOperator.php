<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditionsOperator implements \JsonSerializable
{
    public \Grafana\Foundation\Expr\TypeClassicConditionsType $type;

    /**
     * @param \Grafana\Foundation\Expr\TypeClassicConditionsType|null $type
     */
    public function __construct(?\Grafana\Foundation\Expr\TypeClassicConditionsType $type = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Expr\TypeClassicConditionsType::And();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\TypeClassicConditionsType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        return $data;
    }
}
