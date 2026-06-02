---
title: <span class="badge object-type-enum"></span> SearchTableType
---
# <span class="badge object-type-enum"></span> SearchTableType

The type of the table that is used to display the search results

## Definition

```php
final class SearchTableType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SearchTableType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function traces(): self
    {
        if (!isset(self::$instances["Traces"])) {
            self::$instances["Traces"] = new self("traces");
        }

        return self::$instances["Traces"];
    }

    public static function spans(): self
    {
        if (!isset(self::$instances["Spans"])) {
            self::$instances["Spans"] = new self("spans");
        }

        return self::$instances["Spans"];
    }

    public static function raw(): self
    {
        if (!isset(self::$instances["Raw"])) {
            self::$instances["Raw"] = new self("raw");
        }

        return self::$instances["Raw"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "traces") {
            return self::traces();
        }

        if ($value === "spans") {
            return self::spans();
        }

        if ($value === "raw") {
            return self::raw();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SearchTableType");
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
