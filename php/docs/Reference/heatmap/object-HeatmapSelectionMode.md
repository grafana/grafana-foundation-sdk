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
        if (!isset(self::$instances["X"])) {
            self::$instances["X"] = new self("x");
        }

        return self::$instances["X"];
    }

    public static function y(): self
    {
        if (!isset(self::$instances["Y"])) {
            self::$instances["Y"] = new self("y");
        }

        return self::$instances["Y"];
    }

    public static function xy(): self
    {
        if (!isset(self::$instances["Xy"])) {
            self::$instances["Xy"] = new self("xy");
        }

        return self::$instances["Xy"];
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
