---
title: <span class="badge object-type-enum"></span> TableCellDisplayMode
---
# <span class="badge object-type-enum"></span> TableCellDisplayMode

Internally, this is the "type" of cell that's being displayed

in the table such as colored text, JSON, gauge, etc.

The color-background-solid, gradient-gauge, and lcd-gauge

modes are deprecated in favor of new cell subOptions

## Definition

```php
final class TableCellDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TableCellDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function colorText(): self
    {
        if (!isset(self::$instances["ColorText"])) {
            self::$instances["ColorText"] = new self("color-text");
        }

        return self::$instances["ColorText"];
    }

    public static function colorBackground(): self
    {
        if (!isset(self::$instances["ColorBackground"])) {
            self::$instances["ColorBackground"] = new self("color-background");
        }

        return self::$instances["ColorBackground"];
    }

    public static function colorBackgroundSolid(): self
    {
        if (!isset(self::$instances["ColorBackgroundSolid"])) {
            self::$instances["ColorBackgroundSolid"] = new self("color-background-solid");
        }

        return self::$instances["ColorBackgroundSolid"];
    }

    public static function gradientGauge(): self
    {
        if (!isset(self::$instances["GradientGauge"])) {
            self::$instances["GradientGauge"] = new self("gradient-gauge");
        }

        return self::$instances["GradientGauge"];
    }

    public static function lcdGauge(): self
    {
        if (!isset(self::$instances["LcdGauge"])) {
            self::$instances["LcdGauge"] = new self("lcd-gauge");
        }

        return self::$instances["LcdGauge"];
    }

    public static function jSONView(): self
    {
        if (!isset(self::$instances["JSONView"])) {
            self::$instances["JSONView"] = new self("json-view");
        }

        return self::$instances["JSONView"];
    }

    public static function basicGauge(): self
    {
        if (!isset(self::$instances["BasicGauge"])) {
            self::$instances["BasicGauge"] = new self("basic");
        }

        return self::$instances["BasicGauge"];
    }

    public static function image(): self
    {
        if (!isset(self::$instances["Image"])) {
            self::$instances["Image"] = new self("image");
        }

        return self::$instances["Image"];
    }

    public static function gauge(): self
    {
        if (!isset(self::$instances["Gauge"])) {
            self::$instances["Gauge"] = new self("gauge");
        }

        return self::$instances["Gauge"];
    }

    public static function sparkline(): self
    {
        if (!isset(self::$instances["Sparkline"])) {
            self::$instances["Sparkline"] = new self("sparkline");
        }

        return self::$instances["Sparkline"];
    }

    public static function dataLinks(): self
    {
        if (!isset(self::$instances["DataLinks"])) {
            self::$instances["DataLinks"] = new self("data-links");
        }

        return self::$instances["DataLinks"];
    }

    public static function custom(): self
    {
        if (!isset(self::$instances["Custom"])) {
            self::$instances["Custom"] = new self("custom");
        }

        return self::$instances["Custom"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "color-text") {
            return self::colorText();
        }

        if ($value === "color-background") {
            return self::colorBackground();
        }

        if ($value === "color-background-solid") {
            return self::colorBackgroundSolid();
        }

        if ($value === "gradient-gauge") {
            return self::gradientGauge();
        }

        if ($value === "lcd-gauge") {
            return self::lcdGauge();
        }

        if ($value === "json-view") {
            return self::jSONView();
        }

        if ($value === "basic") {
            return self::basicGauge();
        }

        if ($value === "image") {
            return self::image();
        }

        if ($value === "gauge") {
            return self::gauge();
        }

        if ($value === "sparkline") {
            return self::sparkline();
        }

        if ($value === "data-links") {
            return self::dataLinks();
        }

        if ($value === "custom") {
            return self::custom();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TableCellDisplayMode");
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
