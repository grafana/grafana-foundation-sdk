---
title: <span class="badge object-type-enum"></span> GraphDrawStyle
---
# <span class="badge object-type-enum"></span> GraphDrawStyle

TODO docs

## Definition

```php
final class GraphDrawStyle implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GraphDrawStyle>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function line(): self
    {
        if (!isset(self::$instances["Line"])) {
            self::$instances["Line"] = new self("line");
        }

        return self::$instances["Line"];
    }

    public static function bars(): self
    {
        if (!isset(self::$instances["Bars"])) {
            self::$instances["Bars"] = new self("bars");
        }

        return self::$instances["Bars"];
    }

    public static function points(): self
    {
        if (!isset(self::$instances["Points"])) {
            self::$instances["Points"] = new self("points");
        }

        return self::$instances["Points"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "line") {
            return self::line();
        }

        if ($value === "bars") {
            return self::bars();
        }

        if ($value === "points") {
            return self::points();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GraphDrawStyle");
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
