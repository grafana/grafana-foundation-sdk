<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridLayoutSpecColumnWidthMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AutoGridLayoutSpecColumnWidthMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function narrow(): self
    {
        if (!isset(self::$instances["Narrow"])) {
            self::$instances["Narrow"] = new self("narrow");
        }

        return self::$instances["Narrow"];
    }

    public static function standard(): self
    {
        if (!isset(self::$instances["Standard"])) {
            self::$instances["Standard"] = new self("standard");
        }

        return self::$instances["Standard"];
    }

    public static function wide(): self
    {
        if (!isset(self::$instances["Wide"])) {
            self::$instances["Wide"] = new self("wide");
        }

        return self::$instances["Wide"];
    }

    public static function custom(): self
    {
        if (!isset(self::$instances["Custom"])) {
            self::$instances["Custom"] = new self("custom");
        }

        return self::$instances["Custom"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "narrow") {
            return self::narrow();
        }

        if ($value === "standard") {
            return self::standard();
        }

        if ($value === "wide") {
            return self::wide();
        }

        if ($value === "custom") {
            return self::custom();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AutoGridLayoutSpecColumnWidthMode");
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

