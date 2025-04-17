<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorOrderByExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions $order;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $property
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions|null $order
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions $order = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->property = $property ?: new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty();
        $this->order = $order ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions::Asc();
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: mixed, order?: string, type?: string} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            order: isset($data["order"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByOptions::fromValue($input); })($data["order"]) : null,
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
        $data->order = $this->order;
        $data->type = $this->type;
        return $data;
    }
}
