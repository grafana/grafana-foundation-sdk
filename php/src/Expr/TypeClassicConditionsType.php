<?php

namespace Grafana\Foundation\Expr;

final class TypeClassicConditionsType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeClassicConditionsType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function and(): self
    {
        if (!isset(self::$instances["And"])) {
            self::$instances["And"] = new self("and");
        }

        return self::$instances["And"];
    }

    public static function or(): self
    {
        if (!isset(self::$instances["Or"])) {
            self::$instances["Or"] = new self("or");
        }

        return self::$instances["Or"];
    }

    public static function logicOr(): self
    {
        if (!isset(self::$instances["LogicOr"])) {
            self::$instances["LogicOr"] = new self("logic-or");
        }

        return self::$instances["LogicOr"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "and") {
            return self::and();
        }

        if ($value === "or") {
            return self::or();
        }

        if ($value === "logic-or") {
            return self::logicOr();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeClassicConditionsType");
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

