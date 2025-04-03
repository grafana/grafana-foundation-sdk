---
title: <span class="badge object-type-enum"></span> FieldTextAlignment
---
# <span class="badge object-type-enum"></span> FieldTextAlignment

TODO -- should not be table specific!

TODO docs

## Definition

```php
final class FieldTextAlignment implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, FieldTextAlignment>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["auto"])) {
            self::$instances["auto"] = new self("auto");
        }

        return self::$instances["auto"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["left"])) {
            self::$instances["left"] = new self("left");
        }

        return self::$instances["left"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["right"])) {
            self::$instances["right"] = new self("right");
        }

        return self::$instances["right"];
    }

    public static function center(): self
    {
        if (!isset(self::$instances["center"])) {
            self::$instances["center"] = new self("center");
        }

        return self::$instances["center"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "left") {
            return self::left();
        }

        if ($value === "right") {
            return self::right();
        }

        if ($value === "center") {
            return self::center();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FieldTextAlignment");
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
