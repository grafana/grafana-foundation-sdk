---
title: <span class="badge object-type-enum"></span> BigValueGraphMode
---
# <span class="badge object-type-enum"></span> BigValueGraphMode

TODO docs

## Definition

```php
final class BigValueGraphMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BigValueGraphMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["none"])) {
            self::$instances["none"] = new self("none");
        }

        return self::$instances["none"];
    }

    public static function line(): self
    {
        if (!isset(self::$instances["line"])) {
            self::$instances["line"] = new self("line");
        }

        return self::$instances["line"];
    }

    public static function area(): self
    {
        if (!isset(self::$instances["area"])) {
            self::$instances["area"] = new self("area");
        }

        return self::$instances["area"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "line") {
            return self::line();
        }

        if ($value === "area") {
            return self::area();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BigValueGraphMode");
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
