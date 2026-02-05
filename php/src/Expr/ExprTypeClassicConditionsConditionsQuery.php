<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditionsQuery implements \JsonSerializable
{
    /**
     * @var array<string>
     */
    public array $params;

    /**
     * @param array<string>|null $params
     */
    public function __construct(?array $params = null)
    {
        $this->params = $params ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{params?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            params: $data["params"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->params = $this->params;
        return $data;
    }
}
