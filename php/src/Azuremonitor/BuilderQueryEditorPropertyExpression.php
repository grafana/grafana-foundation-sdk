<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorPropertyExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $property
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->property = $property ?: new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty();
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: mixed, type?: string} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "property" => $this->property,
            "type" => $this->type,
        ];
        return $data;
    }
}
