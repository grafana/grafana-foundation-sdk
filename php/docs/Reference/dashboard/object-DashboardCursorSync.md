---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

0 for no shared crosshair or tooltip (default).

1 for shared crosshair.

2 for shared crosshair AND shared tooltip.

## Definition

```php
final class DashboardCursorSync implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, DashboardCursorSync>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function off(): self
    {
        if (!isset(self::$instances["Off"])) {
            self::$instances["Off"] = new self(0);
        }

        return self::$instances["Off"];
    }

    public static function crosshair(): self
    {
        if (!isset(self::$instances["Crosshair"])) {
            self::$instances["Crosshair"] = new self(1);
        }

        return self::$instances["Crosshair"];
    }

    public static function tooltip(): self
    {
        if (!isset(self::$instances["Tooltip"])) {
            self::$instances["Tooltip"] = new self(2);
        }

        return self::$instances["Tooltip"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::off();
        }

        if ($value === 1) {
            return self::crosshair();
        }

        if ($value === 2) {
            return self::tooltip();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DashboardCursorSync");
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
