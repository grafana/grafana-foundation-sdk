---
title: <span class="badge object-type-enum"></span> FieldColorModeId
---
# <span class="badge object-type-enum"></span> FieldColorModeId

Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.

Continuous color interpolates a color using the percentage of a value relative to min and max.

Accepted values are:

`thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold

`palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations

`palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations

`continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode

`continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode

`continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode

`continuous-YlRd`: Continuous Yellow-Red palette mode

`continuous-BlPu`: Continuous Blue-Purple palette mode

`continuous-YlBl`: Continuous Yellow-Blue palette mode

`continuous-blues`: Continuous Blue palette mode

`continuous-reds`: Continuous Red palette mode

`continuous-greens`: Continuous Green palette mode

`continuous-purples`: Continuous Purple palette mode

`shades`: Shades of a single color. Specify a single color, useful in an override rule.

`fixed`: Fixed color mode. Specify a single color, useful in an override rule.

## Definition

```php
final class FieldColorModeId implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, FieldColorModeId>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function thresholds(): self
    {
        if (!isset(self::$instances["Thresholds"])) {
            self::$instances["Thresholds"] = new self("thresholds");
        }

        return self::$instances["Thresholds"];
    }

    public static function paletteClassic(): self
    {
        if (!isset(self::$instances["PaletteClassic"])) {
            self::$instances["PaletteClassic"] = new self("palette-classic");
        }

        return self::$instances["PaletteClassic"];
    }

    public static function paletteClassicByName(): self
    {
        if (!isset(self::$instances["PaletteClassicByName"])) {
            self::$instances["PaletteClassicByName"] = new self("palette-classic-by-name");
        }

        return self::$instances["PaletteClassicByName"];
    }

    public static function continuousGrYlRd(): self
    {
        if (!isset(self::$instances["ContinuousGrYlRd"])) {
            self::$instances["ContinuousGrYlRd"] = new self("continuous-GrYlRd");
        }

        return self::$instances["ContinuousGrYlRd"];
    }

    public static function continuousRdYlGr(): self
    {
        if (!isset(self::$instances["ContinuousRdYlGr"])) {
            self::$instances["ContinuousRdYlGr"] = new self("continuous-RdYlGr");
        }

        return self::$instances["ContinuousRdYlGr"];
    }

    public static function continuousBlYlRd(): self
    {
        if (!isset(self::$instances["ContinuousBlYlRd"])) {
            self::$instances["ContinuousBlYlRd"] = new self("continuous-BlYlRd");
        }

        return self::$instances["ContinuousBlYlRd"];
    }

    public static function continuousYlRd(): self
    {
        if (!isset(self::$instances["ContinuousYlRd"])) {
            self::$instances["ContinuousYlRd"] = new self("continuous-YlRd");
        }

        return self::$instances["ContinuousYlRd"];
    }

    public static function continuousBlPu(): self
    {
        if (!isset(self::$instances["ContinuousBlPu"])) {
            self::$instances["ContinuousBlPu"] = new self("continuous-BlPu");
        }

        return self::$instances["ContinuousBlPu"];
    }

    public static function continuousYlBl(): self
    {
        if (!isset(self::$instances["ContinuousYlBl"])) {
            self::$instances["ContinuousYlBl"] = new self("continuous-YlBl");
        }

        return self::$instances["ContinuousYlBl"];
    }

    public static function continuousBlues(): self
    {
        if (!isset(self::$instances["ContinuousBlues"])) {
            self::$instances["ContinuousBlues"] = new self("continuous-blues");
        }

        return self::$instances["ContinuousBlues"];
    }

    public static function continuousReds(): self
    {
        if (!isset(self::$instances["ContinuousReds"])) {
            self::$instances["ContinuousReds"] = new self("continuous-reds");
        }

        return self::$instances["ContinuousReds"];
    }

    public static function continuousGreens(): self
    {
        if (!isset(self::$instances["ContinuousGreens"])) {
            self::$instances["ContinuousGreens"] = new self("continuous-greens");
        }

        return self::$instances["ContinuousGreens"];
    }

    public static function continuousPurples(): self
    {
        if (!isset(self::$instances["ContinuousPurples"])) {
            self::$instances["ContinuousPurples"] = new self("continuous-purples");
        }

        return self::$instances["ContinuousPurples"];
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["Fixed"])) {
            self::$instances["Fixed"] = new self("fixed");
        }

        return self::$instances["Fixed"];
    }

    public static function shades(): self
    {
        if (!isset(self::$instances["Shades"])) {
            self::$instances["Shades"] = new self("shades");
        }

        return self::$instances["Shades"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "thresholds") {
            return self::thresholds();
        }

        if ($value === "palette-classic") {
            return self::paletteClassic();
        }

        if ($value === "palette-classic-by-name") {
            return self::paletteClassicByName();
        }

        if ($value === "continuous-GrYlRd") {
            return self::continuousGrYlRd();
        }

        if ($value === "continuous-RdYlGr") {
            return self::continuousRdYlGr();
        }

        if ($value === "continuous-BlYlRd") {
            return self::continuousBlYlRd();
        }

        if ($value === "continuous-YlRd") {
            return self::continuousYlRd();
        }

        if ($value === "continuous-BlPu") {
            return self::continuousBlPu();
        }

        if ($value === "continuous-YlBl") {
            return self::continuousYlBl();
        }

        if ($value === "continuous-blues") {
            return self::continuousBlues();
        }

        if ($value === "continuous-reds") {
            return self::continuousReds();
        }

        if ($value === "continuous-greens") {
            return self::continuousGreens();
        }

        if ($value === "continuous-purples") {
            return self::continuousPurples();
        }

        if ($value === "fixed") {
            return self::fixed();
        }

        if ($value === "shades") {
            return self::shades();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FieldColorModeId");
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
