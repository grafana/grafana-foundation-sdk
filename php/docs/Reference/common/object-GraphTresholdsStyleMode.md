---
title: <span class="badge object-type-enum"></span> GraphTresholdsStyleMode
---
# <span class="badge object-type-enum"></span> GraphTresholdsStyleMode

TODO docs

## Definition

```php
final class GraphTresholdsStyleMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GraphTresholdsStyleMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function off(): self
    {
        if (!isset(self::$instances["Off"])) {
            self::$instances["Off"] = new self("off");
        }

        return self::$instances["Off"];
    }

    public static function line(): self
    {
        if (!isset(self::$instances["Line"])) {
            self::$instances["Line"] = new self("line");
        }

        return self::$instances["Line"];
    }

    public static function dashed(): self
    {
        if (!isset(self::$instances["Dashed"])) {
            self::$instances["Dashed"] = new self("dashed");
        }

        return self::$instances["Dashed"];
    }

    public static function area(): self
    {
        if (!isset(self::$instances["Area"])) {
            self::$instances["Area"] = new self("area");
        }

        return self::$instances["Area"];
    }

    public static function lineAndArea(): self
    {
        if (!isset(self::$instances["LineAndArea"])) {
            self::$instances["LineAndArea"] = new self("line+area");
        }

        return self::$instances["LineAndArea"];
    }

    public static function dashedAndArea(): self
    {
        if (!isset(self::$instances["DashedAndArea"])) {
            self::$instances["DashedAndArea"] = new self("dashed+area");
        }

        return self::$instances["DashedAndArea"];
    }

    public static function series(): self
    {
        if (!isset(self::$instances["Series"])) {
            self::$instances["Series"] = new self("series");
        }

        return self::$instances["Series"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "off") {
            return self::off();
        }

        if ($value === "line") {
            return self::line();
        }

        if ($value === "dashed") {
            return self::dashed();
        }

        if ($value === "area") {
            return self::area();
        }

        if ($value === "line+area") {
            return self::lineAndArea();
        }

        if ($value === "dashed+area") {
            return self::dashedAndArea();
        }

        if ($value === "series") {
            return self::series();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GraphTresholdsStyleMode");
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
