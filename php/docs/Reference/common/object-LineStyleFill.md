---
title: <span class="badge object-type-enum"></span> LineStyleFill
---
# <span class="badge object-type-enum"></span> LineStyleFill

## Definition

```php
final class LineStyleFill implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LineStyleFill>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function solid(): self
    {
        if (!isset(self::$instances["Solid"])) {
            self::$instances["Solid"] = new self("solid");
        }

        return self::$instances["Solid"];
    }

    public static function dash(): self
    {
        if (!isset(self::$instances["Dash"])) {
            self::$instances["Dash"] = new self("dash");
        }

        return self::$instances["Dash"];
    }

    public static function dot(): self
    {
        if (!isset(self::$instances["Dot"])) {
            self::$instances["Dot"] = new self("dot");
        }

        return self::$instances["Dot"];
    }

    public static function square(): self
    {
        if (!isset(self::$instances["Square"])) {
            self::$instances["Square"] = new self("square");
        }

        return self::$instances["Square"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "solid") {
            return self::solid();
        }

        if ($value === "dash") {
            return self::dash();
        }

        if ($value === "dot") {
            return self::dot();
        }

        if ($value === "square") {
            return self::square();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LineStyleFill");
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
