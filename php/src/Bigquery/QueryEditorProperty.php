<?php

namespace Grafana\Foundation\Bigquery;

class QueryEditorProperty implements \JsonSerializable
{
    public \Grafana\Foundation\Bigquery\QueryEditorPropertyType $type;

    public ?string $name;

    /**
     * @param \Grafana\Foundation\Bigquery\QueryEditorPropertyType|null $type
     * @param string|null $name
     */
    public function __construct(?\Grafana\Foundation\Bigquery\QueryEditorPropertyType $type = null, ?string $name = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Bigquery\QueryEditorPropertyType::String();
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
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Bigquery\QueryEditorPropertyType::fromValue($input); })($data["type"]) : null,
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
