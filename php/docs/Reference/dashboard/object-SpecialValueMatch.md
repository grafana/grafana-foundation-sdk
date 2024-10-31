---
title: <span class="badge object-type-enum"></span> SpecialValueMatch
---
# <span class="badge object-type-enum"></span> SpecialValueMatch

Special value types supported by the `SpecialValueMap`

## Definition

```php
final class SpecialValueMatch implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SpecialValueMatch>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function true(): self
    {
        if (!isset(self::$instances["True"])) {
            self::$instances["True"] = new self("true");
        }

        return self::$instances["True"];
    }

    public static function false(): self
    {
        if (!isset(self::$instances["False"])) {
            self::$instances["False"] = new self("false");
        }

        return self::$instances["False"];
    }

    public static function null(): self
    {
        if (!isset(self::$instances["Null"])) {
            self::$instances["Null"] = new self("null");
        }

        return self::$instances["Null"];
    }

    public static function naN(): self
    {
        if (!isset(self::$instances["NaN"])) {
            self::$instances["NaN"] = new self("nan");
        }

        return self::$instances["NaN"];
    }

    public static function nullAndNan(): self
    {
        if (!isset(self::$instances["NullAndNan"])) {
            self::$instances["NullAndNan"] = new self("null+nan");
        }

        return self::$instances["NullAndNan"];
    }

    public static function empty(): self
    {
        if (!isset(self::$instances["Empty"])) {
            self::$instances["Empty"] = new self("empty");
        }

        return self::$instances["Empty"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "true") {
            return self::true();
        }

        if ($value === "false") {
            return self::false();
        }

        if ($value === "null") {
            return self::null();
        }

        if ($value === "nan") {
            return self::naN();
        }

        if ($value === "null+nan") {
            return self::nullAndNan();
        }

        if ($value === "empty") {
            return self::empty();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SpecialValueMatch");
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
