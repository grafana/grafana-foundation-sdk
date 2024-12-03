<?php

namespace Grafana\Foundation\Bigquery;

class QueryEditorGroupByExpression implements \JsonSerializable
{
    public string $type;

    public \Grafana\Foundation\Bigquery\QueryEditorProperty $property;

    /**
     * @param \Grafana\Foundation\Bigquery\QueryEditorProperty|null $property
     */
    public function __construct(?\Grafana\Foundation\Bigquery\QueryEditorProperty $property = null)
    {
        $this->type = "groupBy";
    
        $this->property = $property ?: new \Grafana\Foundation\Bigquery\QueryEditorProperty();
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
    	return \Grafana\Foundation\Bigquery\QueryEditorProperty::fromArray($val);
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
