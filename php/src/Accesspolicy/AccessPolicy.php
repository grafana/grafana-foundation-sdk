<?php

namespace Grafana\Foundation\Accesspolicy;

class AccessPolicy implements \JsonSerializable
{
    /**
     * The scope where these policies should apply
     */
    public \Grafana\Foundation\Accesspolicy\ResourceRef $scope;

    /**
     * The role that must apply this policy
     */
    public \Grafana\Foundation\Accesspolicy\RoleRef $role;

    /**
     * The set of rules to apply.  Note that * is required to modify
     * access policy rules, and that "none" will reject all actions
     * @var array<\Grafana\Foundation\Accesspolicy\AccessRule>
     */
    public array $rules;

    /**
     * @param \Grafana\Foundation\Accesspolicy\ResourceRef|null $scope
     * @param \Grafana\Foundation\Accesspolicy\RoleRef|null $role
     * @param array<\Grafana\Foundation\Accesspolicy\AccessRule>|null $rules
     */
    public function __construct(?\Grafana\Foundation\Accesspolicy\ResourceRef $scope = null, ?\Grafana\Foundation\Accesspolicy\RoleRef $role = null, ?array $rules = null)
    {
        $this->scope = $scope ?: new \Grafana\Foundation\Accesspolicy\ResourceRef();
        $this->role = $role ?: new \Grafana\Foundation\Accesspolicy\RoleRef();
        $this->rules = $rules ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{scope?: mixed, role?: mixed, rules?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            scope: isset($data["scope"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Accesspolicy\ResourceRef::fromArray($val);
    })($data["scope"]) : null,
            role: isset($data["role"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string, xname?: string} */
    $val = $input;
    	return \Grafana\Foundation\Accesspolicy\RoleRef::fromArray($val);
    })($data["role"]) : null,
            rules: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, verb?: string, target?: string} */
    $val = $input;
    	return \Grafana\Foundation\Accesspolicy\AccessRule::fromArray($val);
    }), $data["rules"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "scope" => $this->scope,
            "role" => $this->role,
            "rules" => $this->rules,
        ];
        return $data;
    }
}
