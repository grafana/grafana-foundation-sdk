---
title: <span class="badge object-type-enum"></span> PieChartLabels
---
# <span class="badge object-type-enum"></span> PieChartLabels

Select labels to display on the pie chart.

 - Name - The series or field name.

 - Percent - The percentage of the whole.

 - Value - The raw numerical value.

## Definition

```php
final class PieChartLabels implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PieChartLabels>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function name(): self
    {
        if (!isset(self::$instances["Name"])) {
            self::$instances["Name"] = new self("name");
        }

        return self::$instances["Name"];
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
        if ($value === "name") {
            return self::name();
        }

        if ($value === "value") {
            return self::value();
        }

        if ($value === "percent") {
            return self::percent();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PieChartLabels");
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
