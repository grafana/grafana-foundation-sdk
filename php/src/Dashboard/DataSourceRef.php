<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Ref to a DataSource instance
 */
class DataSourceRef implements \JsonSerializable
{
    /**
     * The plugin type-id
     */
    public ?string $type;

    /**
     * Specific datasource instance
     */
    public ?string $uid;

    /**
     * @param string|null $type
     * @param string|null $uid
     */
    public function __construct(?string $type = null, ?string $uid = null)
    {
        $this->type = $type;
        $this->uid = $uid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            uid: $data["uid"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->type)) {
            $data["type"] = $this->type;
        }
        if (isset($this->uid)) {
            $data["uid"] = $this->uid;
        }
        return $data;
    }
}
