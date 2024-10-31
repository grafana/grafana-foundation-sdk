---
title: <span class="badge object-type-enum"></span> LoadingState
---
# <span class="badge object-type-enum"></span> LoadingState

Loading status

Accepted values are `NotStarted` (the request is not started), `Loading` (waiting for response), `Streaming` (pulling continuous data), `Done` (response received successfully) or `Error` (failed request).

## Definition

```php
final class LoadingState implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LoadingState>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function notStarted(): self
    {
        if (!isset(self::$instances["NotStarted"])) {
            self::$instances["NotStarted"] = new self("NotStarted");
        }

        return self::$instances["NotStarted"];
    }

    public static function loading(): self
    {
        if (!isset(self::$instances["Loading"])) {
            self::$instances["Loading"] = new self("Loading");
        }

        return self::$instances["Loading"];
    }

    public static function streaming(): self
    {
        if (!isset(self::$instances["Streaming"])) {
            self::$instances["Streaming"] = new self("Streaming");
        }

        return self::$instances["Streaming"];
    }

    public static function done(): self
    {
        if (!isset(self::$instances["Done"])) {
            self::$instances["Done"] = new self("Done");
        }

        return self::$instances["Done"];
    }

    public static function error(): self
    {
        if (!isset(self::$instances["Error"])) {
            self::$instances["Error"] = new self("Error");
        }

        return self::$instances["Error"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "NotStarted") {
            return self::notStarted();
        }

        if ($value === "Loading") {
            return self::loading();
        }

        if ($value === "Streaming") {
            return self::streaming();
        }

        if ($value === "Done") {
            return self::done();
        }

        if ($value === "Error") {
            return self::error();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LoadingState");
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
