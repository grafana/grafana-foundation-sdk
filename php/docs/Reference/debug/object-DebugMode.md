---
title: <span class="badge object-type-enum"></span> DebugMode
---
# <span class="badge object-type-enum"></span> DebugMode

## Definition

```php
final class DebugMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DebugMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function render(): self
    {
        if (!isset(self::$instances["render"])) {
            self::$instances["render"] = new self("render");
        }

        return self::$instances["render"];
    }

    public static function events(): self
    {
        if (!isset(self::$instances["events"])) {
            self::$instances["events"] = new self("events");
        }

        return self::$instances["events"];
    }

    public static function cursor(): self
    {
        if (!isset(self::$instances["cursor"])) {
            self::$instances["cursor"] = new self("cursor");
        }

        return self::$instances["cursor"];
    }

    public static function state(): self
    {
        if (!isset(self::$instances["State"])) {
            self::$instances["State"] = new self("State");
        }

        return self::$instances["State"];
    }

    public static function throwError(): self
    {
        if (!isset(self::$instances["ThrowError"])) {
            self::$instances["ThrowError"] = new self("ThrowError");
        }

        return self::$instances["ThrowError"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "render") {
            return self::render();
        }

        if ($value === "events") {
            return self::events();
        }

        if ($value === "cursor") {
            return self::cursor();
        }

        if ($value === "State") {
            return self::state();
        }

        if ($value === "ThrowError") {
            return self::throwError();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DebugMode");
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
