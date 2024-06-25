<?php

namespace Grafana\Foundation\Dashboard;

class AnnotationPanelFilter implements \JsonSerializable
{
    /**
     * Should the specified panels be included or excluded
     */
    public ?bool $exclude;

    /**
     * Panel IDs that should be included or excluded
     * @var array<int>
     */
    public array $ids;

    /**
     * @param bool|null $exclude
     * @param array<int>|null $ids
     */
    public function __construct(?bool $exclude = null, ?array $ids = null)
    {
        $this->exclude = $exclude;
        $this->ids = $ids ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{exclude?: bool, ids?: array<int>} $inputData */
        $data = $inputData;
        return new self(
            exclude: $data["exclude"] ?? null,
            ids: $data["ids"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "ids" => $this->ids,
        ];
        if (isset($this->exclude)) {
            $data["exclude"] = $this->exclude;
        }
        return $data;
    }
}
