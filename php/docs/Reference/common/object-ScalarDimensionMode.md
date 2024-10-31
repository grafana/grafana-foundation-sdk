---
title: <span class="badge object-type-enum"></span> ScalarDimensionMode
---
# <span class="badge object-type-enum"></span> ScalarDimensionMode

## Definition

```php
final class ScalarDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ScalarDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function mod(): self
    {
        if (!isset(self::$instances["mod"])) {
            self::$instances["mod"] = new self("mod");
        }

        return self::$instances["mod"];
    }

    public static function clamped(): self
    {
        if (!isset(self::$instances["clamped"])) {
            self::$instances["clamped"] = new self("clamped");
        }

        return self::$instances["clamped"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "mod") {
            return self::mod();
        }

        if ($value === "clamped") {
            return self::clamped();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScalarDimensionMode");
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
