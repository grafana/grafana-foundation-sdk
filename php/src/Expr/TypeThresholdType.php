<?php

namespace Grafana\Foundation\Expr;

final class TypeThresholdType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeThresholdType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function gt(): self
    {
        if (!isset(self::$instances["Gt"])) {
            self::$instances["Gt"] = new self("gt");
        }

        return self::$instances["Gt"];
    }

    public static function lt(): self
    {
        if (!isset(self::$instances["Lt"])) {
            self::$instances["Lt"] = new self("lt");
        }

        return self::$instances["Lt"];
    }

    public static function withinRange(): self
    {
        if (!isset(self::$instances["WithinRange"])) {
            self::$instances["WithinRange"] = new self("within_range");
        }

        return self::$instances["WithinRange"];
    }

    public static function outsideRange(): self
    {
        if (!isset(self::$instances["OutsideRange"])) {
            self::$instances["OutsideRange"] = new self("outside_range");
        }

        return self::$instances["OutsideRange"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "gt") {
            return self::gt();
        }

        if ($value === "lt") {
            return self::lt();
        }

        if ($value === "within_range") {
            return self::withinRange();
        }

        if ($value === "outside_range") {
            return self::outsideRange();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeThresholdType");
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

