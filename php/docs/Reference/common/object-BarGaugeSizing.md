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
        if (!isset(self::$instances["auto"])) {
            self::$instances["auto"] = new self("auto");
        }

        return self::$instances["auto"];
    }

    public static function manual(): self
    {
        if (!isset(self::$instances["manual"])) {
            self::$instances["manual"] = new self("manual");
        }

        return self::$instances["manual"];
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
