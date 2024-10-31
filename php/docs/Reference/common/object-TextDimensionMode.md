---
title: <span class="badge object-type-enum"></span> TextDimensionMode
---
# <span class="badge object-type-enum"></span> TextDimensionMode

## Definition

```php
final class TextDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TextDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["fixed"])) {
            self::$instances["fixed"] = new self("fixed");
        }

        return self::$instances["fixed"];
    }

    public static function field(): self
    {
        if (!isset(self::$instances["field"])) {
            self::$instances["field"] = new self("field");
        }

        return self::$instances["field"];
    }

    public static function template(): self
    {
        if (!isset(self::$instances["template"])) {
            self::$instances["template"] = new self("template");
        }

        return self::$instances["template"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fixed") {
            return self::fixed();
        }

        if ($value === "field") {
            return self::field();
        }

        if ($value === "template") {
            return self::template();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TextDimensionMode");
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
