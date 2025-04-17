<?php

namespace Grafana\Foundation\Rolebinding;

class RoleBindingSubject implements \JsonSerializable
{
    public \Grafana\Foundation\Rolebinding\RoleBindingSubjectKind $kind;

    /**
     * The team/user identifier name
     */
    public string $name;

    /**
     * @param \Grafana\Foundation\Rolebinding\RoleBindingSubjectKind|null $kind
     * @param string|null $name
     */
    public function __construct(?\Grafana\Foundation\Rolebinding\RoleBindingSubjectKind $kind = null, ?string $name = null)
    {
        $this->kind = $kind ?: \Grafana\Foundation\Rolebinding\RoleBindingSubjectKind::Team();
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            kind: isset($data["kind"]) ? (function($input) { return \Grafana\Foundation\Rolebinding\RoleBindingSubjectKind::fromValue($input); })($data["kind"]) : null,
            name: $data["name"] ?? null,
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
        return $data;
    }
}
