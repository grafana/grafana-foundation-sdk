---
title: <span class="badge object-type-enum"></span> TypeResampleDownsampler
---
# <span class="badge object-type-enum"></span> TypeResampleDownsampler

## Definition

```php
final class TypeResampleDownsampler implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeResampleDownsampler>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function sum(): self
    {
        if (!isset(self::$instances["Sum"])) {
            self::$instances["Sum"] = new self("sum");
        }

        return self::$instances["Sum"];
    }

    public static function mean(): self
    {
        if (!isset(self::$instances["Mean"])) {
            self::$instances["Mean"] = new self("mean");
        }

        return self::$instances["Mean"];
    }

    public static function min(): self
    {
        if (!isset(self::$instances["Min"])) {
            self::$instances["Min"] = new self("min");
        }

        return self::$instances["Min"];
    }

    public static function max(): self
    {
        if (!isset(self::$instances["Max"])) {
            self::$instances["Max"] = new self("max");
        }

        return self::$instances["Max"];
    }

    public static function count(): self
    {
        if (!isset(self::$instances["Count"])) {
            self::$instances["Count"] = new self("count");
        }

        return self::$instances["Count"];
    }

    public static function last(): self
    {
        if (!isset(self::$instances["Last"])) {
            self::$instances["Last"] = new self("last");
        }

        return self::$instances["Last"];
    }

    public static function median(): self
    {
        if (!isset(self::$instances["Median"])) {
            self::$instances["Median"] = new self("median");
        }

        return self::$instances["Median"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "sum") {
            return self::sum();
        }

        if ($value === "mean") {
            return self::mean();
        }

        if ($value === "min") {
            return self::min();
        }

        if ($value === "max") {
            return self::max();
        }

        if ($value === "count") {
            return self::count();
        }

        if ($value === "last") {
            return self::last();
        }

        if ($value === "median") {
            return self::median();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeResampleDownsampler");
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
