<?php

namespace Grafana\Foundation\Common;

final class TextDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TextDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["Fixed"])) {
            self::$instances["Fixed"] = new self("fixed");
        }

        return self::$instances["Fixed"];
    }

    public static function field(): self
    {
        if (!isset(self::$instances["Field"])) {
            self::$instances["Field"] = new self("field");
        }

        return self::$instances["Field"];
    }

    public static function template(): self
    {
        if (!isset(self::$instances["Template"])) {
            self::$instances["Template"] = new self("template");
        }

        return self::$instances["Template"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fixed") {
            return self::fixed();
        }

        if ($value === "field") {
            return self::field();
        }

        if ($value === "template") {
            return self::template();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TextDimensionMode");
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

