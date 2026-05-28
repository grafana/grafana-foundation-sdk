---
title: <span class="badge object-type-enum"></span> FieldColorSeriesByMode
---
# <span class="badge object-type-enum"></span> FieldColorSeriesByMode

Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.

## Definition

```php
final class FieldColorSeriesByMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, FieldColorSeriesByMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function min(): self
    {
        if (!isset(self::$instances["min"])) {
            self::$instances["min"] = new self("min");
        }

        return self::$instances["min"];
    }

    public static function max(): self
    {
        if (!isset(self::$instances["max"])) {
            self::$instances["max"] = new self("max");
        }

        return self::$instances["max"];
    }

    public static function last(): self
    {
        if (!isset(self::$instances["last"])) {
            self::$instances["last"] = new self("last");
        }

        return self::$instances["last"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "min") {
            return self::min();
        }

        if ($value === "max") {
            return self::max();
        }

        if ($value === "last") {
            return self::last();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FieldColorSeriesByMode");
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
