---
title: <span class="badge object-type-enum"></span> CandleStyle
---
# <span class="badge object-type-enum"></span> CandleStyle

## Definition

```php
final class CandleStyle implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, CandleStyle>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function candles(): self
    {
        if (!isset(self::$instances["Candles"])) {
            self::$instances["Candles"] = new self("candles");
        }

        return self::$instances["Candles"];
    }

    public static function oHLCBars(): self
    {
        if (!isset(self::$instances["OHLCBars"])) {
            self::$instances["OHLCBars"] = new self("ohlcbars");
        }

        return self::$instances["OHLCBars"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "candles") {
            return self::candles();
        }

        if ($value === "ohlcbars") {
            return self::oHLCBars();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum CandleStyle");
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
