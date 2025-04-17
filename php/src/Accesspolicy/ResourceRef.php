<?php

namespace Grafana\Foundation\Accesspolicy;

class ResourceRef implements \JsonSerializable
{
    public string $kind;

    public string $name;

    /**
     * @param string|null $kind
     * @param string|null $name
     */
    public function __construct(?string $kind = null, ?string $name = null)
    {
        $this->kind = $kind ?: "";
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            kind: $data["kind"] ?? null,
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->name = $this->name;
        return $data;
    }
}
