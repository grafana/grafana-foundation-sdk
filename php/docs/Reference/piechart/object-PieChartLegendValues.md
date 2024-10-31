---
title: <span class="badge object-type-enum"></span> PieChartLegendValues
---
# <span class="badge object-type-enum"></span> PieChartLegendValues

Select values to display in the legend.

 - Percent: The percentage of the whole.

 - Value: The raw numerical value.

## Definition

```php
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
        if (!isset(self::$instances["value"])) {
            self::$instances["value"] = new self("value");
        }

        return self::$instances["value"];
    }

    public static function percent(): self
    {
        if (!isset(self::$instances["percent"])) {
            self::$instances["percent"] = new self("percent");
        }

        return self::$instances["percent"];
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

```
