<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditionsOperator implements \JsonSerializable
{
    public \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorType $type;

    /**
     * @param \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorType|null $type
     */
    public function __construct(?\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorType $type = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorType::And();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        return $data;
    }
}
