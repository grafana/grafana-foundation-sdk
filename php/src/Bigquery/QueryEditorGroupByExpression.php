<?php

namespace Grafana\Foundation\Bigquery;

class QueryEditorGroupByExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Bigquery\QueryEditorExpressionType $type;

    public \Grafana\Foundation\Bigquery\QueryEditorProperty $property;

    /**
     * @param \Grafana\Foundation\Bigquery\QueryEditorProperty|null $property
     */
    public function __construct(?\Grafana\Foundation\Bigquery\QueryEditorProperty $property = null)
    {
        $this->type = \Grafana\Foundation\Bigquery\QueryEditorExpressionType::property();
        $this->property = $property ?: new \Grafana\Foundation\Bigquery\QueryEditorProperty();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "groupBy", property?: mixed} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: "string", name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\QueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
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
        return $data;
    }
}
