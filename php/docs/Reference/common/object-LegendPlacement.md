---
title: <span class="badge object-type-enum"></span> LegendPlacement
---
# <span class="badge object-type-enum"></span> LegendPlacement

TODO docs

## Definition

```php
final class LegendPlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LegendPlacement>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function bottom(): self
    {
        if (!isset(self::$instances["bottom"])) {
            self::$instances["bottom"] = new self("bottom");
        }

        return self::$instances["bottom"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["right"])) {
            self::$instances["right"] = new self("right");
        }

        return self::$instances["right"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "bottom") {
            return self::bottom();
        }

        if ($value === "right") {
            return self::right();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LegendPlacement");
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
