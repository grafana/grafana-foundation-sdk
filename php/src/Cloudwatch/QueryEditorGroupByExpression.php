<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorGroupByExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType $type;

    public \Grafana\Foundation\Cloudwatch\QueryEditorProperty $property;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorProperty|null $property
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorProperty $property = null)
    {
        $this->type = \Grafana\Foundation\Cloudwatch\QueryEditorExpressionType::property();
        $this->property = $property ?: new \Grafana\Foundation\Cloudwatch\QueryEditorProperty();
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
    	return \Grafana\Foundation\Cloudwatch\QueryEditorProperty::fromArray($val);
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
