<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
final class AxisPlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AxisPlacement>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function top(): self
    {
        if (!isset(self::$instances["Top"])) {
            self::$instances["Top"] = new self("top");
        }

        return self::$instances["Top"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["Right"])) {
            self::$instances["Right"] = new self("right");
        }

        return self::$instances["Right"];
    }

    public static function bottom(): self
    {
        if (!isset(self::$instances["Bottom"])) {
            self::$instances["Bottom"] = new self("bottom");
        }

        return self::$instances["Bottom"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["Left"])) {
            self::$instances["Left"] = new self("left");
        }

        return self::$instances["Left"];
    }

    public static function hidden(): self
    {
        if (!isset(self::$instances["Hidden"])) {
            self::$instances["Hidden"] = new self("hidden");
        }

        return self::$instances["Hidden"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "top") {
            return self::top();
        }

        if ($value === "right") {
            return self::right();
        }

        if ($value === "bottom") {
            return self::bottom();
        }

        if ($value === "left") {
            return self::left();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AxisPlacement");
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

