<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditionsEvaluator implements \JsonSerializable
{
    /**
     * @var array<float>
     */
    public array $params;

    /**
     * e.g. "gt"
     */
    public string $type;

    /**
     * @param array<float>|null $params
     * @param string|null $type
     */
    public function __construct(?array $params = null, ?string $type = null)
    {
        $this->params = $params ?: [];
        $this->type = $type ?: "";
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
            type: $data["type"] ?? null,
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
