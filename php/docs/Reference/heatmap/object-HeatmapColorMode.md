---
title: <span class="badge object-type-enum"></span> HeatmapColorMode
---
# <span class="badge object-type-enum"></span> HeatmapColorMode

Controls the color mode of the heatmap

## Definition

```php
final class HeatmapColorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapColorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function opacity(): self
    {
        if (!isset(self::$instances["opacity"])) {
            self::$instances["opacity"] = new self("opacity");
        }

        return self::$instances["opacity"];
    }

    public static function scheme(): self
    {
        if (!isset(self::$instances["scheme"])) {
            self::$instances["scheme"] = new self("scheme");
        }

        return self::$instances["scheme"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "opacity") {
            return self::opacity();
        }

        if ($value === "scheme") {
            return self::scheme();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapColorMode");
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
