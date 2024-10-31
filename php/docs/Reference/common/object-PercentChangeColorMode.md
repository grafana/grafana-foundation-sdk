---
title: <span class="badge object-type-enum"></span> PercentChangeColorMode
---
# <span class="badge object-type-enum"></span> PercentChangeColorMode

TODO docs

## Definition

```php
final class PercentChangeColorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PercentChangeColorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function standard(): self
    {
        if (!isset(self::$instances["Standard"])) {
            self::$instances["Standard"] = new self("standard");
        }

        return self::$instances["Standard"];
    }

    public static function inverted(): self
    {
        if (!isset(self::$instances["Inverted"])) {
            self::$instances["Inverted"] = new self("inverted");
        }

        return self::$instances["Inverted"];
    }

    public static function sameAsValue(): self
    {
        if (!isset(self::$instances["SameAsValue"])) {
            self::$instances["SameAsValue"] = new self("same_as_value");
        }

        return self::$instances["SameAsValue"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "standard") {
            return self::standard();
        }

        if ($value === "inverted") {
            return self::inverted();
        }

        if ($value === "same_as_value") {
            return self::sameAsValue();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PercentChangeColorMode");
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
