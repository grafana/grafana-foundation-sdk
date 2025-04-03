---
title: <span class="badge object-type-enum"></span> HorizontalConstraint
---
# <span class="badge object-type-enum"></span> HorizontalConstraint

## Definition

```php
final class HorizontalConstraint implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HorizontalConstraint>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function left(): self
    {
        if (!isset(self::$instances["Left"])) {
            self::$instances["Left"] = new self("left");
        }

        return self::$instances["Left"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["Right"])) {
            self::$instances["Right"] = new self("right");
        }

        return self::$instances["Right"];
    }

    public static function leftRight(): self
    {
        if (!isset(self::$instances["LeftRight"])) {
            self::$instances["LeftRight"] = new self("leftright");
        }

        return self::$instances["LeftRight"];
    }

    public static function center(): self
    {
        if (!isset(self::$instances["Center"])) {
            self::$instances["Center"] = new self("center");
        }

        return self::$instances["Center"];
    }

    public static function scale(): self
    {
        if (!isset(self::$instances["Scale"])) {
            self::$instances["Scale"] = new self("scale");
        }

        return self::$instances["Scale"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "left") {
            return self::left();
        }

        if ($value === "right") {
            return self::right();
        }

        if ($value === "leftright") {
            return self::leftRight();
        }

        if ($value === "center") {
            return self::center();
        }

        if ($value === "scale") {
            return self::scale();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HorizontalConstraint");
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
