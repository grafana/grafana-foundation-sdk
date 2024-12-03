<?php

namespace Grafana\Foundation\Bigquery;

class QueryEditorFunctionParameterExpression implements \JsonSerializable
{
    public string $type;

    public ?string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->type = "functionParameter";
    
        $this->name = $name;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
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
        return $data;
    }
}
