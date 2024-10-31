---
title: <span class="badge object-type-enum"></span> PieChartType
---
# <span class="badge object-type-enum"></span> PieChartType

Select the pie chart display style.

## Definition

```php
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
        if (!isset(self::$instances["pie"])) {
            self::$instances["pie"] = new self("pie");
        }

        return self::$instances["pie"];
    }

    public static function donut(): self
    {
        if (!isset(self::$instances["donut"])) {
            self::$instances["donut"] = new self("donut");
        }

        return self::$instances["donut"];
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

```
