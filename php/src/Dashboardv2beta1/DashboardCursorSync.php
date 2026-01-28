<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * "Off" for no shared crosshair or tooltip (default).
 * "Crosshair" for shared crosshair.
 * "Tooltip" for shared crosshair AND shared tooltip.
 */
final class DashboardCursorSync implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DashboardCursorSync>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function crosshair(): self
    {
        if (!isset(self::$instances["Crosshair"])) {
            self::$instances["Crosshair"] = new self("Crosshair");
        }

        return self::$instances["Crosshair"];
    }

    public static function tooltip(): self
    {
        if (!isset(self::$instances["Tooltip"])) {
            self::$instances["Tooltip"] = new self("Tooltip");
        }

        return self::$instances["Tooltip"];
    }

    public static function off(): self
    {
        if (!isset(self::$instances["Off"])) {
            self::$instances["Off"] = new self("Off");
        }

        return self::$instances["Off"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Crosshair") {
            return self::crosshair();
        }

        if ($value === "Tooltip") {
            return self::tooltip();
        }

        if ($value === "Off") {
            return self::off();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DashboardCursorSync");
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

