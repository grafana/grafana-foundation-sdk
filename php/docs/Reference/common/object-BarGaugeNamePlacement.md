---
title: <span class="badge object-type-enum"></span> BarGaugeNamePlacement
---
# <span class="badge object-type-enum"></span> BarGaugeNamePlacement

Allows for the bar gauge name to be placed explicitly

## Definition

```php
final class BarGaugeNamePlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BarGaugeNamePlacement>
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

    public static function top(): self
    {
        if (!isset(self::$instances["top"])) {
            self::$instances["top"] = new self("top");
        }

        return self::$instances["top"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["left"])) {
            self::$instances["left"] = new self("left");
        }

        return self::$instances["left"];
    }

    public static function hidden(): self
    {
        if (!isset(self::$instances["hidden"])) {
            self::$instances["hidden"] = new self("hidden");
        }

        return self::$instances["hidden"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "top") {
            return self::top();
        }

        if ($value === "left") {
            return self::left();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarGaugeNamePlacement");
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