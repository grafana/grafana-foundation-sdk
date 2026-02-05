---
title: <span class="badge object-type-enum"></span> SearchStreamingState
---
# <span class="badge object-type-enum"></span> SearchStreamingState

The state of the TraceQL streaming search query

## Definition

```php
final class SearchStreamingState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SearchStreamingState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function pending(): self
    {
        if (!isset(self::$instances["pending"])) {
            self::$instances["pending"] = new self("pending");
        }

        return self::$instances["pending"];
    }

    public static function streaming(): self
    {
        if (!isset(self::$instances["streaming"])) {
            self::$instances["streaming"] = new self("streaming");
        }

        return self::$instances["streaming"];
    }

    public static function done(): self
    {
        if (!isset(self::$instances["done"])) {
            self::$instances["done"] = new self("done");
        }

        return self::$instances["done"];
    }

    public static function error(): self
    {
        if (!isset(self::$instances["error"])) {
            self::$instances["error"] = new self("error");
        }

        return self::$instances["error"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "pending") {
            return self::pending();
        }

        if ($value === "streaming") {
            return self::streaming();
        }

        if ($value === "done") {
            return self::done();
        }

        if ($value === "error") {
            return self::error();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SearchStreamingState");
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
