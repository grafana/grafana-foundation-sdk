---
title: <span class="badge object-type-enum"></span> HeatmapSelectionMode
---
# <span class="badge object-type-enum"></span> HeatmapSelectionMode

Controls which axis to allow selection on

## Definition

```php
final class HeatmapSelectionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapSelectionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function x(): self
    {
        if (!isset(self::$instances["x"])) {
            self::$instances["x"] = new self("x");
        }

        return self::$instances["x"];
    }

    public static function y(): self
    {
        if (!isset(self::$instances["y"])) {
            self::$instances["y"] = new self("y");
        }

        return self::$instances["y"];
    }

    public static function xy(): self
    {
        if (!isset(self::$instances["xy"])) {
            self::$instances["xy"] = new self("xy");
        }

        return self::$instances["xy"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "x") {
            return self::x();
        }

        if ($value === "y") {
            return self::y();
        }

        if ($value === "xy") {
            return self::xy();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapSelectionMode");
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
