---
title: <span class="badge object-type-enum"></span> OrderByDirection
---
# <span class="badge object-type-enum"></span> OrderByDirection

## Definition

```php
final class OrderByDirection implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, OrderByDirection>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function aSC(): self
    {
        if (!isset(self::$instances["ASC"])) {
            self::$instances["ASC"] = new self("ASC");
        }

        return self::$instances["ASC"];
    }

    public static function dESC(): self
    {
        if (!isset(self::$instances["DESC"])) {
            self::$instances["DESC"] = new self("DESC");
        }

        return self::$instances["DESC"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "ASC") {
            return self::aSC();
        }

        if ($value === "DESC") {
            return self::dESC();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum OrderByDirection");
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
