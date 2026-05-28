<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Determine whether regex applies to variable value or display text
 * Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)
 */
final class VariableRegexApplyTo implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableRegexApplyTo>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function value(): self
    {
        if (!isset(self::$instances["value"])) {
            self::$instances["value"] = new self("value");
        }

        return self::$instances["value"];
    }

    public static function text(): self
    {
        if (!isset(self::$instances["text"])) {
            self::$instances["text"] = new self("text");
        }

        return self::$instances["text"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "value") {
            return self::value();
        }

        if ($value === "text") {
            return self::text();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableRegexApplyTo");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

