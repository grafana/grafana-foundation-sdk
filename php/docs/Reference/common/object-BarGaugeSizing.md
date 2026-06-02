---
title: <span class="badge object-type-enum"></span> BarGaugeSizing
---
# <span class="badge object-type-enum"></span> BarGaugeSizing

Allows for the bar gauge size to be set explicitly

## Definition

```php
final class BarGaugeSizing implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BarGaugeSizing>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function manual(): self
    {
        if (!isset(self::$instances["Manual"])) {
            self::$instances["Manual"] = new self("manual");
        }

        return self::$instances["Manual"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "manual") {
            return self::manual();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarGaugeSizing");
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
