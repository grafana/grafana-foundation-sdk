---
title: <span class="badge object-type-enum"></span> VizDisplayMode
---
# <span class="badge object-type-enum"></span> VizDisplayMode

## Definition

```php
final class VizDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VizDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function candlesVolume(): self
    {
        if (!isset(self::$instances["CandlesVolume"])) {
            self::$instances["CandlesVolume"] = new self("candles+volume");
        }

        return self::$instances["CandlesVolume"];
    }

    public static function candles(): self
    {
        if (!isset(self::$instances["Candles"])) {
            self::$instances["Candles"] = new self("candles");
        }

        return self::$instances["Candles"];
    }

    public static function volume(): self
    {
        if (!isset(self::$instances["Volume"])) {
            self::$instances["Volume"] = new self("volume");
        }

        return self::$instances["Volume"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "candles+volume") {
            return self::candlesVolume();
        }

        if ($value === "candles") {
            return self::candles();
        }

        if ($value === "volume") {
            return self::volume();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VizDisplayMode");
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
