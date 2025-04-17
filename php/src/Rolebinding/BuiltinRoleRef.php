<?php

namespace Grafana\Foundation\Rolebinding;

class BuiltinRoleRef implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Rolebinding\BuiltinRoleRefName $name;

    /**
     * @param \Grafana\Foundation\Rolebinding\BuiltinRoleRefName|null $name
     */
    public function __construct(?\Grafana\Foundation\Rolebinding\BuiltinRoleRefName $name = null)
    {
        $this->kind = "BuiltinRole";
    
        $this->name = $name ?: \Grafana\Foundation\Rolebinding\BuiltinRoleRefName::Viewer();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            name: isset($data["name"]) ? (function($input) { return \Grafana\Foundation\Rolebinding\BuiltinRoleRefName::fromValue($input); })($data["name"]) : null,
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
