---
title: <span class="badge object-type-enum"></span> PointShape
---
# <span class="badge object-type-enum"></span> PointShape

## Definition

```php
final class PointShape implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PointShape>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function circle(): self
    {
        if (!isset(self::$instances["circle"])) {
            self::$instances["circle"] = new self("circle");
        }

        return self::$instances["circle"];
    }

    public static function square(): self
    {
        if (!isset(self::$instances["square"])) {
            self::$instances["square"] = new self("square");
        }

        return self::$instances["square"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "circle") {
            return self::circle();
        }

        if ($value === "square") {
            return self::square();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PointShape");
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
