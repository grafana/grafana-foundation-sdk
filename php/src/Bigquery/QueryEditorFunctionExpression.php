<?php

namespace Grafana\Foundation\Bigquery;

class QueryEditorFunctionExpression implements \JsonSerializable
{
    public string $type;

    public ?string $name;

    /**
     * @var array<\Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression>|null
     */
    public ?array $parameters;

    /**
     * @param string|null $name
     * @param array<\Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression>|null $parameters
     */
    public function __construct(?string $name = null, ?array $parameters = null)
    {
        $this->type = "function";
    
        $this->name = $name;
        $this->parameters = $parameters;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, name?: string, parameters?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            parameters: array_filter(array_map((function($input) {
    	/** @var array{type?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression::fromArray($val);
    }), $data["parameters"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->parameters)) {
            $data["parameters"] = $this->parameters;
        }
        return $data;
    }
}
