<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorWhereExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems>
     */
    public array $expressions;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     * @param array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems>|null $expressions
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null, ?array $expressions = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
        $this->expressions = $expressions ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, expressions?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
            expressions: array_filter(array_map((function($input) {
    	/** @var array{property?: mixed, operator?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems::fromArray($val);
    }), $data["expressions"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "expressions" => $this->expressions,
        ];
        return $data;
    }
}
