---
title: <span class="badge object-type-enum"></span> TooltipDisplayMode
---
# <span class="badge object-type-enum"></span> TooltipDisplayMode

TODO docs

## Definition

```php
final class TooltipDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TooltipDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function single(): self
    {
        if (!isset(self::$instances["single"])) {
            self::$instances["single"] = new self("single");
        }

        return self::$instances["single"];
    }

    public static function multi(): self
    {
        if (!isset(self::$instances["multi"])) {
            self::$instances["multi"] = new self("multi");
        }

        return self::$instances["multi"];
    }

    public static function none(): self
    {
        if (!isset(self::$instances["none"])) {
            self::$instances["none"] = new self("none");
        }

        return self::$instances["none"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "single") {
            return self::single();
        }

        if ($value === "multi") {
            return self::multi();
        }

        if ($value === "none") {
            return self::none();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TooltipDisplayMode");
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
