---
title: <span class="badge object-type-enum"></span> BarGaugeValueMode
---
# <span class="badge object-type-enum"></span> BarGaugeValueMode

Allows for the table cell gauge display type to set the gauge mode.

## Definition

```php
final class BarGaugeValueMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BarGaugeValueMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function color(): self
    {
        if (!isset(self::$instances["color"])) {
            self::$instances["color"] = new self("color");
        }

        return self::$instances["color"];
    }

    public static function text(): self
    {
        if (!isset(self::$instances["text"])) {
            self::$instances["text"] = new self("text");
        }

        return self::$instances["text"];
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
        if ($value === "color") {
            return self::color();
        }

        if ($value === "text") {
            return self::text();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarGaugeValueMode");
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
