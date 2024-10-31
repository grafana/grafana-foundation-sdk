---
title: <span class="badge object-type-enum"></span> TypeResampleUpsampler
---
# <span class="badge object-type-enum"></span> TypeResampleUpsampler

## Definition

```php
final class TypeResampleUpsampler implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeResampleUpsampler>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function pad(): self
    {
        if (!isset(self::$instances["Pad"])) {
            self::$instances["Pad"] = new self("pad");
        }

        return self::$instances["Pad"];
    }

    public static function backfilling(): self
    {
        if (!isset(self::$instances["Backfilling"])) {
            self::$instances["Backfilling"] = new self("backfilling");
        }

        return self::$instances["Backfilling"];
    }

    public static function fillna(): self
    {
        if (!isset(self::$instances["Fillna"])) {
            self::$instances["Fillna"] = new self("fillna");
        }

        return self::$instances["Fillna"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "pad") {
            return self::pad();
        }

        if ($value === "backfilling") {
            return self::backfilling();
        }

        if ($value === "fillna") {
            return self::fillna();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeResampleUpsampler");
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
