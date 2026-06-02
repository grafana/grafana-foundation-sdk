<?php

namespace Grafana\Foundation\Piechart;

/**
 * Select the pie chart display style.
 */
final class PieChartType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PieChartType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function pie(): self
    {
        if (!isset(self::$instances["Pie"])) {
            self::$instances["Pie"] = new self("pie");
        }

        return self::$instances["Pie"];
    }

    public static function donut(): self
    {
        if (!isset(self::$instances["Donut"])) {
            self::$instances["Donut"] = new self("donut");
        }

        return self::$instances["Donut"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "pie") {
            return self::pie();
        }

        if ($value === "donut") {
            return self::donut();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PieChartType");
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

