<?php

namespace Grafana\Foundation\Debug;

class UpdateConfig implements \JsonSerializable
{
    public bool $render;

    public bool $dataChanged;

    public bool $schemaChanged;

    /**
     * @param bool|null $render
     * @param bool|null $dataChanged
     * @param bool|null $schemaChanged
     */
    public function __construct(?bool $render = null, ?bool $dataChanged = null, ?bool $schemaChanged = null)
    {
        $this->render = $render ?: false;
        $this->dataChanged = $dataChanged ?: false;
        $this->schemaChanged = $schemaChanged ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{render?: bool, dataChanged?: bool, schemaChanged?: bool} $inputData */
        $data = $inputData;
        return new self(
            render: $data["render"] ?? null,
            dataChanged: $data["dataChanged"] ?? null,
            schemaChanged: $data["schemaChanged"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "render" => $this->render,
            "dataChanged" => $this->dataChanged,
            "schemaChanged" => $this->schemaChanged,
        ];
        return $data;
    }
}
