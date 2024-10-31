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
        if (!isset(self::$instances["line"])) {
            self::$instances["line"] = new self("line");
        }

        return self::$instances["line"];
    }

    public static function bars(): self
    {
        if (!isset(self::$instances["bars"])) {
            self::$instances["bars"] = new self("bars");
        }

        return self::$instances["bars"];
    }

    public static function points(): self
    {
        if (!isset(self::$instances["points"])) {
            self::$instances["points"] = new self("points");
        }

        return self::$instances["points"];
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
