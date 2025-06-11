<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorOperatorExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType $type;

    public \Grafana\Foundation\Cloudwatch\QueryEditorProperty $property;

    /**
     * TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
     */
    public \Grafana\Foundation\Cloudwatch\QueryEditorOperator $operator;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorProperty|null $property
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorOperator|null $operator
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorProperty $property = null, ?\Grafana\Foundation\Cloudwatch\QueryEditorOperator $operator = null)
    {
        $this->type = \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType::property();
        $this->property = $property ?: new \Grafana\Foundation\Cloudwatch\QueryEditorProperty();
        $this->operator = $operator ?: new \Grafana\Foundation\Cloudwatch\QueryEditorOperator();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "operator", property?: mixed, operator?: mixed} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: "string", name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            operator: isset($data["operator"]) ? (function($input) {
    	/** @var array{name?: string, value?: string|bool|int|array<string|bool|int>} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorOperator::fromArray($val);
    })($data["operator"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->property = $this->property;
        $data->operator = $this->operator;
        return $data;
    }
}
