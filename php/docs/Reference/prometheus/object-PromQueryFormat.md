---
title: <span class="badge object-type-enum"></span> PromQueryFormat
---
# <span class="badge object-type-enum"></span> PromQueryFormat

## Definition

```php
final class PromQueryFormat implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PromQueryFormat>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function timeSeries(): self
    {
        if (!isset(self::$instances["TimeSeries"])) {
            self::$instances["TimeSeries"] = new self("time_series");
        }

        return self::$instances["TimeSeries"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self("table");
        }

        return self::$instances["Table"];
    }

    public static function heatmap(): self
    {
        if (!isset(self::$instances["Heatmap"])) {
            self::$instances["Heatmap"] = new self("heatmap");
        }

        return self::$instances["Heatmap"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "time_series") {
            return self::timeSeries();
        }

        if ($value === "table") {
            return self::table();
        }

        if ($value === "heatmap") {
            return self::heatmap();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PromQueryFormat");
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
