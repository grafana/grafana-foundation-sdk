---
title: <span class="badge object-type-enum"></span> GraphGradientMode
---
# <span class="badge object-type-enum"></span> GraphGradientMode

TODO docs

## Definition

```php
final class GraphGradientMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GraphGradientMode>
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

    public static function opacity(): self
    {
        if (!isset(self::$instances["opacity"])) {
            self::$instances["opacity"] = new self("opacity");
        }

        return self::$instances["opacity"];
    }

    public static function hue(): self
    {
        if (!isset(self::$instances["hue"])) {
            self::$instances["hue"] = new self("hue");
        }

        return self::$instances["hue"];
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
        if ($value === "none") {
            return self::none();
        }

        if ($value === "opacity") {
            return self::opacity();
        }

        if ($value === "hue") {
            return self::hue();
        }

        if ($value === "scheme") {
            return self::scheme();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GraphGradientMode");
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
