<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridLayoutSpecRowHeightMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AutoGridLayoutSpecRowHeightMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function short(): self
    {
        if (!isset(self::$instances["Short"])) {
            self::$instances["Short"] = new self("short");
        }

        return self::$instances["Short"];
    }

    public static function standard(): self
    {
        if (!isset(self::$instances["Standard"])) {
            self::$instances["Standard"] = new self("standard");
        }

        return self::$instances["Standard"];
    }

    public static function tall(): self
    {
        if (!isset(self::$instances["Tall"])) {
            self::$instances["Tall"] = new self("tall");
        }

        return self::$instances["Tall"];
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
        if ($value === "short") {
            return self::short();
        }

        if ($value === "standard") {
            return self::standard();
        }

        if ($value === "tall") {
            return self::tall();
        }

        if ($value === "custom") {
            return self::custom();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AutoGridLayoutSpecRowHeightMode");
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

