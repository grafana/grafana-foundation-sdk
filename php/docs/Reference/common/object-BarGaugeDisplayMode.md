---
title: <span class="badge object-type-enum"></span> BarGaugeDisplayMode
---
# <span class="badge object-type-enum"></span> BarGaugeDisplayMode

Enum expressing the possible display modes

for the bar gauge component of Grafana UI

## Definition

```php
final class BarGaugeDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BarGaugeDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function basic(): self
    {
        if (!isset(self::$instances["basic"])) {
            self::$instances["basic"] = new self("basic");
        }

        return self::$instances["basic"];
    }

    public static function lcd(): self
    {
        if (!isset(self::$instances["lcd"])) {
            self::$instances["lcd"] = new self("lcd");
        }

        return self::$instances["lcd"];
    }

    public static function gradient(): self
    {
        if (!isset(self::$instances["gradient"])) {
            self::$instances["gradient"] = new self("gradient");
        }

        return self::$instances["gradient"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "basic") {
            return self::basic();
        }

        if ($value === "lcd") {
            return self::lcd();
        }

        if ($value === "gradient") {
            return self::gradient();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarGaugeDisplayMode");
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
