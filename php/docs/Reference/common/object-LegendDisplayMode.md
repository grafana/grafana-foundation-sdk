---
title: <span class="badge object-type-enum"></span> LegendDisplayMode
---
# <span class="badge object-type-enum"></span> LegendDisplayMode

TODO docs

Note: "hidden" needs to remain as an option for plugins compatibility

## Definition

```php
final class LegendDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LegendDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function list(): self
    {
        if (!isset(self::$instances["list"])) {
            self::$instances["list"] = new self("list");
        }

        return self::$instances["list"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["table"])) {
            self::$instances["table"] = new self("table");
        }

        return self::$instances["table"];
    }

    public static function hidden(): self
    {
        if (!isset(self::$instances["hidden"])) {
            self::$instances["hidden"] = new self("hidden");
        }

        return self::$instances["hidden"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "list") {
            return self::list();
        }

        if ($value === "table") {
            return self::table();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LegendDisplayMode");
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
