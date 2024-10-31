---
title: <span class="badge object-type-enum"></span> ScatterShow
---
# <span class="badge object-type-enum"></span> ScatterShow

## Definition

```php
final class ScatterShow implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ScatterShow>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function points(): self
    {
        if (!isset(self::$instances["Points"])) {
            self::$instances["Points"] = new self("points");
        }

        return self::$instances["Points"];
    }

    public static function lines(): self
    {
        if (!isset(self::$instances["Lines"])) {
            self::$instances["Lines"] = new self("lines");
        }

        return self::$instances["Lines"];
    }

    public static function pointsAndLines(): self
    {
        if (!isset(self::$instances["PointsAndLines"])) {
            self::$instances["PointsAndLines"] = new self("points+lines");
        }

        return self::$instances["PointsAndLines"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "points") {
            return self::points();
        }

        if ($value === "lines") {
            return self::lines();
        }

        if ($value === "points+lines") {
            return self::pointsAndLines();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScatterShow");
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
