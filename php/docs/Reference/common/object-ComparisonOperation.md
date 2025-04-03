---
title: <span class="badge object-type-enum"></span> ComparisonOperation
---
# <span class="badge object-type-enum"></span> ComparisonOperation

Compare two values

## Definition

```php
final class ComparisonOperation implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ComparisonOperation>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function eQ(): self
    {
        if (!isset(self::$instances["EQ"])) {
            self::$instances["EQ"] = new self("eq");
        }

        return self::$instances["EQ"];
    }

    public static function nEQ(): self
    {
        if (!isset(self::$instances["NEQ"])) {
            self::$instances["NEQ"] = new self("neq");
        }

        return self::$instances["NEQ"];
    }

    public static function lT(): self
    {
        if (!isset(self::$instances["LT"])) {
            self::$instances["LT"] = new self("lt");
        }

        return self::$instances["LT"];
    }

    public static function lTE(): self
    {
        if (!isset(self::$instances["LTE"])) {
            self::$instances["LTE"] = new self("lte");
        }

        return self::$instances["LTE"];
    }

    public static function gT(): self
    {
        if (!isset(self::$instances["GT"])) {
            self::$instances["GT"] = new self("gt");
        }

        return self::$instances["GT"];
    }

    public static function gTE(): self
    {
        if (!isset(self::$instances["GTE"])) {
            self::$instances["GTE"] = new self("gte");
        }

        return self::$instances["GTE"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "eq") {
            return self::eQ();
        }

        if ($value === "neq") {
            return self::nEQ();
        }

        if ($value === "lt") {
            return self::lT();
        }

        if ($value === "lte") {
            return self::lTE();
        }

        if ($value === "gt") {
            return self::gT();
        }

        if ($value === "gte") {
            return self::gTE();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ComparisonOperation");
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
