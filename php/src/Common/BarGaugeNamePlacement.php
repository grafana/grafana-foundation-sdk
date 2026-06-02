<?php

namespace Grafana\Foundation\Common;

/**
 * Allows for the bar gauge name to be placed explicitly
 */
final class BarGaugeNamePlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BarGaugeNamePlacement>
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

        if ($value === "left") {
            return self::left();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarGaugeNamePlacement");
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

