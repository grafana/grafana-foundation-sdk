<?php

namespace Grafana\Foundation\Elasticsearch;

class Nested implements \JsonSerializable
{
    public ?string $field;

    public string $id;

    public string $type;

    /**
     * @var mixed|null
     */
    public $settings;

    /**
     * @param string|null $field
     * @param string|null $id
     * @param mixed|null $settings
     */
    public function __construct(?string $field = null, ?string $id = null,  $settings = null)
    {
        $this->field = $field;
        $this->id = $id ?: "";
        $this->type = "nested";
    
        $this->settings = $settings;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{field?: string, id?: string, type?: string, settings?: mixed} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            id: $data["id"] ?? null,
            settings: $data["settings"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
            "type" => $this->type,
        ];
        if (isset($this->field)) {
            $data["field"] = $this->field;
        }
        if (isset($this->settings)) {
            $data["settings"] = $this->settings;
        }
        return $data;
    }
}