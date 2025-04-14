<?php

namespace Grafana\Foundation\Accesspolicy;

class AccessRule implements \JsonSerializable
{
    /**
     * The kind this rule applies to (dashboards, alert, etc)
     */
    public string $kind;

    /**
     * READ, WRITE, CREATE, DELETE, ...
     * should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
     * @var string|string|string
     */
    public $verb;

    /**
     * Specific sub-elements like "alert.rules" or "dashboard.permissions"????
     */
    public ?string $target;

    /**
     * @param string|null $kind
     * @param string|string|string|null $verb
     * @param string|null $target
     */
    public function __construct(?string $kind = null,  $verb = null, ?string $target = null)
    {
        $this->kind = $kind ?: "*";
        $this->verb = $verb ?: "*";
        $this->target = $target;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, verb?: string|string|string, target?: string} $inputData */
        $data = $inputData;
        return new self(
            kind: $data["kind"] ?? null,
            verb: isset($data["verb"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        case is_string($input):
            return $input;
        case is_string($input):
            return $input;
        default:
            throw new \ValueError('incorrect value for disjunction');
    }
    })($data["verb"]) : null,
            target: $data["target"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "kind" => $this->kind,
            "verb" => $this->verb,
        ];
        if (isset($this->target)) {
            $data["target"] = $this->target;
        }
        return $data;
    }
}
