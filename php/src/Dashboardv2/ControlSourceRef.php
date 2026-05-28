<?php

namespace Grafana\Foundation\Dashboardv2;

class ControlSourceRef implements \JsonSerializable
{
    public string $type;

    /**
     * The plugin type-id
     */
    public string $group;

    /**
     * @param string|null $group
     */
    public function __construct(?string $group = null)
    {
        $this->type = "datasource";
    
        $this->group = $group ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, group?: string} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->group = $this->group;
        return $data;
    }
}
