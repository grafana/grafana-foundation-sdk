---
title: <span class="badge object-type-enum"></span> MetricsQueryType
---
# <span class="badge object-type-enum"></span> MetricsQueryType

## Definition

```php
final class MetricsQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MetricsQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function range(): self
    {
        if (!isset(self::$instances["range"])) {
            self::$instances["range"] = new self("range");
        }

        return self::$instances["range"];
    }

    public static function instant(): self
    {
        if (!isset(self::$instances["instant"])) {
            self::$instances["instant"] = new self("instant");
        }

        return self::$instances["instant"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "range") {
            return self::range();
        }

        if ($value === "instant") {
            return self::instant();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricsQueryType");
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
