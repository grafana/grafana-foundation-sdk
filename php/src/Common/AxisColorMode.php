<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
final class AxisColorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AxisColorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function text(): self
    {
        if (!isset(self::$instances["Text"])) {
            self::$instances["Text"] = new self("text");
        }

        return self::$instances["Text"];
    }

    public static function series(): self
    {
        if (!isset(self::$instances["Series"])) {
            self::$instances["Series"] = new self("series");
        }

        return self::$instances["Series"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "text") {
            return self::text();
        }

        if ($value === "series") {
            return self::series();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AxisColorMode");
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

