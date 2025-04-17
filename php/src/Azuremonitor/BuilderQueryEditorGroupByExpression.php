<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorGroupByExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $interval;

    public ?bool $focus;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $property
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $interval
     * @param bool|null $focus
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $interval = null, ?bool $focus = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->property = $property;
        $this->interval = $interval;
        $this->focus = $focus;
        $this->type = $type;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: mixed, interval?: mixed, focus?: bool, type?: string} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            interval: isset($data["interval"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["interval"]) : null,
            focus: $data["focus"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->property)) {
            $data->property = $this->property;
        }
        if (isset($this->interval)) {
            $data->interval = $this->interval;
        }
        if (isset($this->focus)) {
            $data->focus = $this->focus;
        }
        if (isset($this->type)) {
            $data->type = $this->type;
        }
        return $data;
    }
}
