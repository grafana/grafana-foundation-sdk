<?php

namespace Grafana\Foundation\Elasticsearch;

class Count implements \JsonSerializable
{
    public string $type;

    public string $id;

    public ?bool $hide;

    /**
     * @param string|null $id
     * @param bool|null $hide
     */
    public function __construct(?string $id = null, ?bool $hide = null)
    {
        $this->type = "count";
    
        $this->id = $id ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, id?: string, hide?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            hide: $data["hide"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "id" => $this->id,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }
}
