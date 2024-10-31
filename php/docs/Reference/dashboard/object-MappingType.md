---
title: <span class="badge object-type-enum"></span> MappingType
---
# <span class="badge object-type-enum"></span> MappingType

Supported value mapping types

`value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

`range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.

`regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

`special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```php
final class MappingType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MappingType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function valueToText(): self
    {
        if (!isset(self::$instances["ValueToText"])) {
            self::$instances["ValueToText"] = new self("value");
        }

        return self::$instances["ValueToText"];
    }

    public static function rangeToText(): self
    {
        if (!isset(self::$instances["RangeToText"])) {
            self::$instances["RangeToText"] = new self("range");
        }

        return self::$instances["RangeToText"];
    }

    public static function regexToText(): self
    {
        if (!isset(self::$instances["RegexToText"])) {
            self::$instances["RegexToText"] = new self("regex");
        }

        return self::$instances["RegexToText"];
    }

    public static function specialValue(): self
    {
        if (!isset(self::$instances["SpecialValue"])) {
            self::$instances["SpecialValue"] = new self("special");
        }

        return self::$instances["SpecialValue"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "value") {
            return self::valueToText();
        }

        if ($value === "range") {
            return self::rangeToText();
        }

        if ($value === "regex") {
            return self::regexToText();
        }

        if ($value === "special") {
            return self::specialValue();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MappingType");
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
