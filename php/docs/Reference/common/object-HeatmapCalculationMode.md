---
title: <span class="badge object-type-enum"></span> HeatmapCalculationMode
---
# <span class="badge object-type-enum"></span> HeatmapCalculationMode

## Definition

```php
final class HeatmapCalculationMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapCalculationMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function size(): self
    {
        if (!isset(self::$instances["size"])) {
            self::$instances["size"] = new self("size");
        }

        return self::$instances["size"];
    }

    public static function count(): self
    {
        if (!isset(self::$instances["count"])) {
            self::$instances["count"] = new self("count");
        }

        return self::$instances["count"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "size") {
            return self::size();
        }

        if ($value === "count") {
            return self::count();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapCalculationMode");
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
