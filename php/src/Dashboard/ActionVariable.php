<?php

namespace Grafana\Foundation\Dashboard;

class ActionVariable implements \JsonSerializable
{
    public string $key;

    public string $name;

    public \Grafana\Foundation\Dashboard\ActionVariableType $type;

    /**
     * @param string|null $key
     * @param string|null $name
     */
    public function __construct(?string $key = null, ?string $name = null)
    {
        $this->key = $key ?: "";
        $this->name = $name ?: "";
        $this->type = \Grafana\Foundation\Dashboard\ActionVariableType::string();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{key?: string, name?: string, type?: "string"} $inputData */
        $data = $inputData;
        return new self(
            key: $data["key"] ?? null,
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->key = $this->key;
        $data->name = $this->name;
        $data->type = $this->type;
        return $data;
    }
}
