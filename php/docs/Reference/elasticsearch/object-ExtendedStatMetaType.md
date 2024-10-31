---
title: <span class="badge object-type-enum"></span> ExtendedStatMetaType
---
# <span class="badge object-type-enum"></span> ExtendedStatMetaType

## Definition

```php
final class ExtendedStatMetaType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ExtendedStatMetaType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function avg(): self
    {
        if (!isset(self::$instances["avg"])) {
            self::$instances["avg"] = new self("avg");
        }

        return self::$instances["avg"];
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

    public static function sum(): self
    {
        if (!isset(self::$instances["sum"])) {
            self::$instances["sum"] = new self("sum");
        }

        return self::$instances["sum"];
    }

    public static function count(): self
    {
        if (!isset(self::$instances["count"])) {
            self::$instances["count"] = new self("count");
        }

        return self::$instances["count"];
    }

    public static function stdDeviation(): self
    {
        if (!isset(self::$instances["std_deviation"])) {
            self::$instances["std_deviation"] = new self("std_deviation");
        }

        return self::$instances["std_deviation"];
    }

    public static function stdDeviationBoundsUpper(): self
    {
        if (!isset(self::$instances["std_deviation_bounds_upper"])) {
            self::$instances["std_deviation_bounds_upper"] = new self("std_deviation_bounds_upper");
        }

        return self::$instances["std_deviation_bounds_upper"];
    }

    public static function stdDeviationBoundsLower(): self
    {
        if (!isset(self::$instances["std_deviation_bounds_lower"])) {
            self::$instances["std_deviation_bounds_lower"] = new self("std_deviation_bounds_lower");
        }

        return self::$instances["std_deviation_bounds_lower"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "avg") {
            return self::avg();
        }

        if ($value === "min") {
            return self::min();
        }

        if ($value === "max") {
            return self::max();
        }

        if ($value === "sum") {
            return self::sum();
        }

        if ($value === "count") {
            return self::count();
        }

        if ($value === "std_deviation") {
            return self::stdDeviation();
        }

        if ($value === "std_deviation_bounds_upper") {
            return self::stdDeviationBoundsUpper();
        }

        if ($value === "std_deviation_bounds_lower") {
            return self::stdDeviationBoundsLower();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ExtendedStatMetaType");
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
