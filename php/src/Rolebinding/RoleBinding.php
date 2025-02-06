<?php

namespace Grafana\Foundation\Rolebinding;

class RoleBinding implements \JsonSerializable
{
    /**
     * The role we are discussing
     * @var \Grafana\Foundation\Rolebinding\BuiltinRoleRef|\Grafana\Foundation\Rolebinding\CustomRoleRef
     */
    public $role;

    /**
     * The team or user that has the specified role
     */
    public \Grafana\Foundation\Rolebinding\RoleBindingSubject $subject;

    /**
     * @param \Grafana\Foundation\Rolebinding\BuiltinRoleRef|\Grafana\Foundation\Rolebinding\CustomRoleRef|null $role
     * @param \Grafana\Foundation\Rolebinding\RoleBindingSubject|null $subject
     */
    public function __construct( $role = null, ?\Grafana\Foundation\Rolebinding\RoleBindingSubject $subject = null)
    {
        $this->role = $role ?: new \Grafana\Foundation\Rolebinding\BuiltinRoleRef();
        $this->subject = $subject ?: new \Grafana\Foundation\Rolebinding\RoleBindingSubject();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{role?: mixed|mixed, subject?: mixed} $inputData */
        $data = $inputData;
        return new self(
            role: isset($data["role"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "BuiltinRole":
            return BuiltinRoleRef::fromArray($input);
        case "Role":
            return CustomRoleRef::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["role"]) : null,
            subject: isset($data["subject"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Rolebinding\RoleBindingSubject::fromArray($val);
    })($data["subject"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "role" => $this->role,
            "subject" => $this->subject,
        ];
        return $data;
    }
}
