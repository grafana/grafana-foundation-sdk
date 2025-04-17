<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorWhereExpressionArray implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression>
     */
    public array $expressions;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression>|null $expressions
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
    	/** @var array{type?: string, expressions?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpression::fromArray($val);
    }), $data["expressions"] ?? [])),
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->expressions = $this->expressions;
        $data->type = $this->type;
        return $data;
    }
}
