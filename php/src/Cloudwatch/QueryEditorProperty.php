<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorProperty implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType $type;

    public ?string $name;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType|null $type
     * @param string|null $name
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorPropertyType $type = null, ?string $name = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType::String();
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
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType::fromValue($input); })($data["type"]) : null,
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
