<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorArrayExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType $type;

    /**
     * @var array<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression|\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression|\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression>
     */
    public array $expressions;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType|null $type
     * @param array<\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression|\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression|\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression>|null $expressions
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType $type = null, ?array $expressions = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType::And();
        $this->expressions = $expressions ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, expressions?: array<mixed|mixed|mixed|mixed|mixed|mixed>} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType::fromValue($input); })($data["type"]) : null,
            expressions: !empty($data["expressions"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
    
        switch ($input["type"]) {
        case "and":
            return QueryEditorArrayExpression::fromArray($input);
        case "function":
            return QueryEditorFunctionExpression::fromArray($input);
        case "functionParameter":
            return QueryEditorFunctionParameterExpression::fromArray($input);
        case "groupBy":
            return QueryEditorGroupByExpression::fromArray($input);
        case "operator":
            return QueryEditorOperatorExpression::fromArray($input);
        case "or":
            return QueryEditorArrayExpression::fromArray($input);
        case "property":
            return QueryEditorPropertyExpression::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["expressions"]) : null,
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
