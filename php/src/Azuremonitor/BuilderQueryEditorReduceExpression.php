<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorReduceExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $reduce;

    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression>|null
     */
    public ?array $parameters;

    public ?bool $focus;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $property
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty|null $reduce
     * @param array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression>|null $parameters
     * @param bool|null $focus
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $reduce = null, ?array $parameters = null, ?bool $focus = null)
    {
        $this->property = $property;
        $this->reduce = $reduce;
        $this->parameters = $parameters;
        $this->focus = $focus;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: mixed, reduce?: mixed, parameters?: array<mixed>, focus?: bool} $inputData */
        $data = $inputData;
        return new self(
            property: isset($data["property"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["property"]) : null,
            reduce: isset($data["reduce"]) ? (function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty::fromArray($val);
    })($data["reduce"]) : null,
            parameters: array_filter(array_map((function($input) {
    	/** @var array{value?: string, fieldType?: string, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression::fromArray($val);
    }), $data["parameters"] ?? [])),
            focus: $data["focus"] ?? null,
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
        if (isset($this->reduce)) {
            $data->reduce = $this->reduce;
        }
        if (isset($this->parameters)) {
            $data->parameters = $this->parameters;
        }
        if (isset($this->focus)) {
            $data->focus = $this->focus;
        }
        return $data;
    }
}
