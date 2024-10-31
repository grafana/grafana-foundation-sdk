---
title: <span class="badge object-type-enum"></span> HeatmapColorScale
---
# <span class="badge object-type-enum"></span> HeatmapColorScale

Controls the color scale of the heatmap

## Definition

```php
final class HeatmapColorScale implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapColorScale>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function linear(): self
    {
        if (!isset(self::$instances["linear"])) {
            self::$instances["linear"] = new self("linear");
        }

        return self::$instances["linear"];
    }

    public static function exponential(): self
    {
        if (!isset(self::$instances["exponential"])) {
            self::$instances["exponential"] = new self("exponential");
        }

        return self::$instances["exponential"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "linear") {
            return self::linear();
        }

        if ($value === "exponential") {
            return self::exponential();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapColorScale");
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
