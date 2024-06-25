<?php

namespace Grafana\Foundation\Librarypanel;

class LibraryElementDTOMetaUser implements \JsonSerializable
{
    public int $id;

    public string $name;

    public string $avatarUrl;

    /**
     * @param int|null $id
     * @param string|null $name
     * @param string|null $avatarUrl
     */
    public function __construct(?int $id = null, ?string $name = null, ?string $avatarUrl = null)
    {
        $this->id = $id ?: 0;
        $this->name = $name ?: "";
        $this->avatarUrl = $avatarUrl ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: int, name?: string, avatarUrl?: string} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            name: $data["name"] ?? null,
            avatarUrl: $data["avatarUrl"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
            "name" => $this->name,
            "avatarUrl" => $this->avatarUrl,
        ];
        return $data;
    }
}
