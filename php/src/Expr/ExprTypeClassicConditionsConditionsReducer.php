<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditionsReducer implements \JsonSerializable
{
    public string $type;

    /**
     * @param string|null $type
     */
    public function __construct(?string $type = null)
    {
        $this->type = $type ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
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
