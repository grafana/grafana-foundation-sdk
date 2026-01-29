---
title: <span class="badge object-type-enum"></span> NullValueMode
---
# <span class="badge object-type-enum"></span> NullValueMode

How null values should be handled

## Definition

```php
final class NullValueMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, NullValueMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function null(): self
    {
        if (!isset(self::$instances["null"])) {
            self::$instances["null"] = new self("null");
        }

        return self::$instances["null"];
    }

    public static function connected(): self
    {
        if (!isset(self::$instances["connected"])) {
            self::$instances["connected"] = new self("connected");
        }

        return self::$instances["connected"];
    }

    public static function nullAsZero(): self
    {
        if (!isset(self::$instances["null as zero"])) {
            self::$instances["null as zero"] = new self("null as zero");
        }

        return self::$instances["null as zero"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "null") {
            return self::null();
        }

        if ($value === "connected") {
            return self::connected();
        }

        if ($value === "null as zero") {
            return self::nullAsZero();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum NullValueMode");
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
