---
title: <span class="badge object-type-enum"></span> SeriesMapping
---
# <span class="badge object-type-enum"></span> SeriesMapping

Auto is "table" in the UI

## Definition

```php
final class SeriesMapping implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SeriesMapping>
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

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SeriesMapping");
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
