<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorPropertyExpression implements \JsonSerializable
{
    public string $type;

    public \Grafana\Foundation\Cloudwatch\QueryEditorProperty $property;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorProperty|null $property
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorProperty $property = null)
    {
        $this->type = "property";
    
        $this->property = $property ?: new \Grafana\Foundation\Cloudwatch\QueryEditorProperty();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, property?: mixed} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "property" => $this->property,
        ];
        return $data;
    }
}
