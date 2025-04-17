<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorWhereExpressionItems implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator $operator;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $property
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator|null $operator
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator $operator = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->property = $property ?: new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty();
        $this->operator = $operator ?: new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator();
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: mixed, operator?: mixed, type?: string} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            operator: isset($data["operator"]) ? (function($input) {
    	/** @var array{name?: string, value?: string, labelValue?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOperator::fromArray($val);
    })($data["operator"]) : null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->property = $this->property;
        $data->operator = $this->operator;
        $data->type = $this->type;
        return $data;
    }
}
