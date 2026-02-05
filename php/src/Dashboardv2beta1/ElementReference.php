<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ElementReference implements \JsonSerializable
{
    public string $kind;

    public string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->kind = "ElementReference";
    
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
