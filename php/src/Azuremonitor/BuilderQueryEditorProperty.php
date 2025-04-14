<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorProperty implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType $type;

    public string $name;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType|null $type
     * @param string|null $name
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType $type = null, ?string $name = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType::Number();
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType::fromValue($input); })($data["type"]) : null,
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
            "name" => $this->name,
        ];
        return $data;
    }
}
