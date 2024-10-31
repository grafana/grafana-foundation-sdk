---
title: <span class="badge object-type-enum"></span> MetricQueryType
---
# <span class="badge object-type-enum"></span> MetricQueryType

## Definition

```php
final class MetricQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, MetricQueryType>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function search(): self
    {
        if (!isset(self::$instances["Search"])) {
            self::$instances["Search"] = new self(0);
        }

        return self::$instances["Search"];
    }

    public static function query(): self
    {
        if (!isset(self::$instances["Query"])) {
            self::$instances["Query"] = new self(1);
        }

        return self::$instances["Query"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::search();
        }

        if ($value === 1) {
            return self::query();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricQueryType");
    }

    public function jsonSerialize(): int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

```
