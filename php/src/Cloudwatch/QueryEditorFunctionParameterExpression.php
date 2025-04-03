<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorFunctionParameterExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType $type;

    public ?string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->type = \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType::property();
        $this->name = $name;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "functionParameter", name?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
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
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        return $data;
    }
}
