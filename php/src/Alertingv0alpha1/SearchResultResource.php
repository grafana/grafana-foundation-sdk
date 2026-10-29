<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultResource is the full identity of a hit.
 */
class SearchResultResource implements \JsonSerializable
{
    public string $group;

    public string $resource;

    public string $kind;

    public string $name;

    /**
     * @param string|null $group
     * @param string|null $resource
     * @param string|null $kind
     * @param string|null $name
     */
    public function __construct(?string $group = null, ?string $resource = null, ?string $kind = null, ?string $name = null)
    {
        $this->group = $group ?: "";
        $this->resource = $resource ?: "";
        $this->kind = $kind ?: "";
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{group?: string, resource?: string, kind?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
            resource: $data["resource"] ?? null,
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
        $data->group = $this->group;
        $data->resource = $this->resource;
        $data->kind = $this->kind;
        $data->name = $this->name;
        return $data;
    }
}
