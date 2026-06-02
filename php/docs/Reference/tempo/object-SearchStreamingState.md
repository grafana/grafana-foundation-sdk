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
        if (!isset(self::$instances["Pending"])) {
            self::$instances["Pending"] = new self("pending");
        }

        return self::$instances["Pending"];
    }

    public static function streaming(): self
    {
        if (!isset(self::$instances["Streaming"])) {
            self::$instances["Streaming"] = new self("streaming");
        }

        return self::$instances["Streaming"];
    }

    public static function done(): self
    {
        if (!isset(self::$instances["Done"])) {
            self::$instances["Done"] = new self("done");
        }

        return self::$instances["Done"];
    }

    public static function error(): self
    {
        if (!isset(self::$instances["Error"])) {
            self::$instances["Error"] = new self("error");
        }

        return self::$instances["Error"];
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
