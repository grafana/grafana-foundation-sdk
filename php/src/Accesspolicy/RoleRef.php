<?php

namespace Grafana\Foundation\Accesspolicy;

class RoleRef implements \JsonSerializable
{
    /**
     * Policies can apply to roles, teams, or users
     * Applying policies to individual users is supported, but discouraged
     */
    public \Grafana\Foundation\Accesspolicy\RoleRefKind $kind;

    public string $name;

    public string $xname;

    /**
     * @param \Grafana\Foundation\Accesspolicy\RoleRefKind|null $kind
     * @param string|null $name
     * @param string|null $xname
     */
    public function __construct(?\Grafana\Foundation\Accesspolicy\RoleRefKind $kind = null, ?string $name = null, ?string $xname = null)
    {
        $this->kind = $kind ?: \Grafana\Foundation\Accesspolicy\RoleRefKind::Role();
        $this->name = $name ?: "";
        $this->xname = $xname ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, name?: string, xname?: string} $inputData */
        $data = $inputData;
        return new self(
            kind: isset($data["kind"]) ? (function($input) { return \Grafana\Foundation\Accesspolicy\RoleRefKind::fromValue($input); })($data["kind"]) : null,
            name: $data["name"] ?? null,
            xname: $data["xname"] ?? null,
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
        $data->xname = $this->xname;
        return $data;
    }
}
