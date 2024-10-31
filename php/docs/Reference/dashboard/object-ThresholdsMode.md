---
title: <span class="badge object-type-enum"></span> ThresholdsMode
---
# <span class="badge object-type-enum"></span> ThresholdsMode

Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).

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
        if (!isset(self::$instances["Absolute"])) {
            self::$instances["Absolute"] = new self("absolute");
        }

        return self::$instances["Absolute"];
    }

    public static function percentage(): self
    {
        if (!isset(self::$instances["Percentage"])) {
            self::$instances["Percentage"] = new self("percentage");
        }

        return self::$instances["Percentage"];
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
