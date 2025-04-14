<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorReduceExpressionArray implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression>
     */
    public array $expressions;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression>|null $expressions
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?array $expressions = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->expressions = $expressions ?: [];
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{expressions?: array<mixed>, type?: string} $inputData */
        $data = $inputData;
        return new self(
            expressions: array_filter(array_map((function($input) {
    	/** @var array{property?: mixed, reduce?: mixed, parameters?: array<mixed>, focus?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression::fromArray($val);
    }), $data["expressions"] ?? [])),
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "expressions" => $this->expressions,
            "type" => $this->type,
        ];
        return $data;
    }
}
