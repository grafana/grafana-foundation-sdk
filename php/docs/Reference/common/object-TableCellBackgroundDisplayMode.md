---
title: <span class="badge object-type-enum"></span> TableCellBackgroundDisplayMode
---
# <span class="badge object-type-enum"></span> TableCellBackgroundDisplayMode

Display mode to the "Colored Background" display

mode for table cells. Either displays a solid color (basic mode)

or a gradient.

## Definition

```php
final class TableCellBackgroundDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TableCellBackgroundDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function basic(): self
    {
        if (!isset(self::$instances["Basic"])) {
            self::$instances["Basic"] = new self("basic");
        }

        return self::$instances["Basic"];
    }

    public static function gradient(): self
    {
        if (!isset(self::$instances["Gradient"])) {
            self::$instances["Gradient"] = new self("gradient");
        }

        return self::$instances["Gradient"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "basic") {
            return self::basic();
        }

        if ($value === "gradient") {
            return self::gradient();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TableCellBackgroundDisplayMode");
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
