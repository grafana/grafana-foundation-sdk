<?php

namespace Grafana\Foundation\Starsv1alpha1;

class Resource implements \JsonSerializable
{
    public string $group;

    public string $kind;

    /**
     * The set of resources
     * +listType=set
     * @var array<string>
     */
    public array $names;

    /**
     * @param string|null $group
     * @param string|null $kind
     * @param array<string>|null $names
     */
    public function __construct(?string $group = null, ?string $kind = null, ?array $names = null)
    {
        $this->group = $group ?: "";
        $this->kind = $kind ?: "";
        $this->names = $names ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{group?: string, kind?: string, names?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
            kind: $data["kind"] ?? null,
            names: $data["names"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->group = $this->group;
        $data->kind = $this->kind;
        $data->names = $this->names;
        return $data;
    }
}
