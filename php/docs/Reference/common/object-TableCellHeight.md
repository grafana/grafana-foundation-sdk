---
title: <span class="badge object-type-enum"></span> TableCellHeight
---
# <span class="badge object-type-enum"></span> TableCellHeight

Height of a table cell

## Definition

```php
final class TableCellHeight implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TableCellHeight>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function sm(): self
    {
        if (!isset(self::$instances["Sm"])) {
            self::$instances["Sm"] = new self("sm");
        }

        return self::$instances["Sm"];
    }

    public static function md(): self
    {
        if (!isset(self::$instances["Md"])) {
            self::$instances["Md"] = new self("md");
        }

        return self::$instances["Md"];
    }

    public static function lg(): self
    {
        if (!isset(self::$instances["Lg"])) {
            self::$instances["Lg"] = new self("lg");
        }

        return self::$instances["Lg"];
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "sm") {
            return self::sm();
        }

        if ($value === "md") {
            return self::md();
        }

        if ($value === "lg") {
            return self::lg();
        }

        if ($value === "auto") {
            return self::auto();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TableCellHeight");
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
