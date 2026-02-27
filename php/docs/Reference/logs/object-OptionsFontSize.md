---
title: <span class="badge object-type-enum"></span> OptionsFontSize
---
# <span class="badge object-type-enum"></span> OptionsFontSize

## Definition

```php
final class OptionsFontSize implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, OptionsFontSize>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function default(): self
    {
        if (!isset(self::$instances["Default"])) {
            self::$instances["Default"] = new self("default");
        }

        return self::$instances["Default"];
    }

    public static function small(): self
    {
        if (!isset(self::$instances["Small"])) {
            self::$instances["Small"] = new self("small");
        }

        return self::$instances["Small"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "default") {
            return self::default();
        }

        if ($value === "small") {
            return self::small();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum OptionsFontSize");
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
