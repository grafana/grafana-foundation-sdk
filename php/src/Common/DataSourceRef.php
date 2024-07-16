<?php

namespace Grafana\Foundation\Common;

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
     * Datasource API version
     */
    public ?string $apiVersion;

    /**
     * @param string|null $type
     * @param string|null $uid
     * @param string|null $apiVersion
     */
    public function __construct(?string $type = null, ?string $uid = null, ?string $apiVersion = null)
    {
        $this->type = $type;
        $this->uid = $uid;
        $this->apiVersion = $apiVersion;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, uid?: string, apiVersion?: string} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            uid: $data["uid"] ?? null,
            apiVersion: $data["apiVersion"] ?? null,
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
        if (isset($this->apiVersion)) {
            $data["apiVersion"] = $this->apiVersion;
        }
        return $data;
    }
}
