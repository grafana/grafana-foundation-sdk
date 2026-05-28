---
title: <span class="badge object-type-enum"></span> ThresholdsMode
---
# <span class="badge object-type-enum"></span> ThresholdsMode

## Definition

```php
final class ThresholdsMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ThresholdsMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function absolute(): self
    {
        if (!isset(self::$instances["absolute"])) {
            self::$instances["absolute"] = new self("absolute");
        }

        return self::$instances["absolute"];
    }

    public static function percentage(): self
    {
        if (!isset(self::$instances["percentage"])) {
            self::$instances["percentage"] = new self("percentage");
        }

        return self::$instances["percentage"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "absolute") {
            return self::absolute();
        }

        if ($value === "percentage") {
            return self::percentage();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ThresholdsMode");
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
