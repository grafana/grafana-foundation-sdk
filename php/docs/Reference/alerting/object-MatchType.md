---
title: <span class="badge object-type-enum"></span> MatchType
---
# <span class="badge object-type-enum"></span> MatchType

## Definition

```php
final class MatchType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MatchType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function equal(): self
    {
        if (!isset(self::$instances["equal"])) {
            self::$instances["equal"] = new self("=");
        }

        return self::$instances["equal"];
    }

    public static function notEqual(): self
    {
        if (!isset(self::$instances["not-equal"])) {
            self::$instances["not-equal"] = new self("!=");
        }

        return self::$instances["not-equal"];
    }

    public static function equalRegex(): self
    {
        if (!isset(self::$instances["equal-regex"])) {
            self::$instances["equal-regex"] = new self("=~");
        }

        return self::$instances["equal-regex"];
    }

    public static function notEqualRegex(): self
    {
        if (!isset(self::$instances["not-equal-regex"])) {
            self::$instances["not-equal-regex"] = new self("!~");
        }

        return self::$instances["not-equal-regex"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "=") {
            return self::equal();
        }

        if ($value === "!=") {
            return self::notEqual();
        }

        if ($value === "=~") {
            return self::equalRegex();
        }

        if ($value === "!~") {
            return self::notEqualRegex();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MatchType");
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
