---
title: <span class="badge object-type-enum"></span> HeatmapCellLayout
---
# <span class="badge object-type-enum"></span> HeatmapCellLayout

## Definition

```php
final class HeatmapCellLayout implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapCellLayout>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function le(): self
    {
        if (!isset(self::$instances["Le"])) {
            self::$instances["Le"] = new self("le");
        }

        return self::$instances["Le"];
    }

    public static function ge(): self
    {
        if (!isset(self::$instances["Ge"])) {
            self::$instances["Ge"] = new self("ge");
        }

        return self::$instances["Ge"];
    }

    public static function unknown(): self
    {
        if (!isset(self::$instances["Unknown"])) {
            self::$instances["Unknown"] = new self("unknown");
        }

        return self::$instances["Unknown"];
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "le") {
            return self::le();
        }

        if ($value === "ge") {
            return self::ge();
        }

        if ($value === "unknown") {
            return self::unknown();
        }

        if ($value === "auto") {
            return self::auto();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapCellLayout");
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
