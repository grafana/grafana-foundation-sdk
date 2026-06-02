<?php

namespace Grafana\Foundation\Piechart;

/**
 * Select values to display in the legend.
 *  - Percent: The percentage of the whole.
 *  - Value: The raw numerical value.
 */
final class PieChartLegendValues implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PieChartLegendValues>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function value(): self
    {
        if (!isset(self::$instances["Value"])) {
            self::$instances["Value"] = new self("value");
        }

        return self::$instances["Value"];
    }

    public static function percent(): self
    {
        if (!isset(self::$instances["Percent"])) {
            self::$instances["Percent"] = new self("percent");
        }

        return self::$instances["Percent"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "value") {
            return self::value();
        }

        if ($value === "percent") {
            return self::percent();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PieChartLegendValues");
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

