<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Options to config when to refresh a variable
 * `never`: Never refresh the variable
 * `onDashboardLoad`: Queries the data source every time the dashboard loads.
 * `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
 */
final class VariableRefresh implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableRefresh>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function never(): self
    {
        if (!isset(self::$instances["never"])) {
            self::$instances["never"] = new self("never");
        }

        return self::$instances["never"];
    }

    public static function onDashboardLoad(): self
    {
        if (!isset(self::$instances["onDashboardLoad"])) {
            self::$instances["onDashboardLoad"] = new self("onDashboardLoad");
        }

        return self::$instances["onDashboardLoad"];
    }

    public static function onTimeRangeChanged(): self
    {
        if (!isset(self::$instances["onTimeRangeChanged"])) {
            self::$instances["onTimeRangeChanged"] = new self("onTimeRangeChanged");
        }

        return self::$instances["onTimeRangeChanged"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "never") {
            return self::never();
        }

        if ($value === "onDashboardLoad") {
            return self::onDashboardLoad();
        }

        if ($value === "onTimeRangeChanged") {
            return self::onTimeRangeChanged();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableRefresh");
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

