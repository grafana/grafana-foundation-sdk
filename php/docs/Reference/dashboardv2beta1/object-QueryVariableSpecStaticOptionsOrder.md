---
title: <span class="badge object-type-enum"></span> QueryVariableSpecStaticOptionsOrder
---
# <span class="badge object-type-enum"></span> QueryVariableSpecStaticOptionsOrder

## Definition

```php
final class QueryVariableSpecStaticOptionsOrder implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryVariableSpecStaticOptionsOrder>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function before(): self
    {
        if (!isset(self::$instances["Before"])) {
            self::$instances["Before"] = new self("before");
        }

        return self::$instances["Before"];
    }

    public static function after(): self
    {
        if (!isset(self::$instances["After"])) {
            self::$instances["After"] = new self("after");
        }

        return self::$instances["After"];
    }

    public static function sorted(): self
    {
        if (!isset(self::$instances["Sorted"])) {
            self::$instances["Sorted"] = new self("sorted");
        }

        return self::$instances["Sorted"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "before") {
            return self::before();
        }

        if ($value === "after") {
            return self::after();
        }

        if ($value === "sorted") {
            return self::sorted();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryVariableSpecStaticOptionsOrder");
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
