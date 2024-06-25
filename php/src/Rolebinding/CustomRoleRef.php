<?php

namespace Grafana\Foundation\Rolebinding;

class CustomRoleRef implements \JsonSerializable
{
    public string $kind;

    public string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->kind = "Role";
    
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
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "kind" => $this->kind,
            "name" => $this->name,
        ];
        return $data;
    }
}
