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

`continuous-viridis`: Continuous Viridis palette mode

`continuous-magma`: Continuous Magma palette mode

`continuous-plasma`: Continuous Plasma palette mode

`continuous-inferno`: Continuous Inferno palette mode

`continuous-cividis`: Continuous Cividis palette mode

`continuous-GrYlRd`: Continuous Green-Yellow-Red palette mode

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
        if (!isset(self::$instances["thresholds"])) {
            self::$instances["thresholds"] = new self("thresholds");
        }

        return self::$instances["thresholds"];
    }

    public static function paletteClassic(): self
    {
        if (!isset(self::$instances["palette-classic"])) {
            self::$instances["palette-classic"] = new self("palette-classic");
        }

        return self::$instances["palette-classic"];
    }

    public static function paletteClassicByName(): self
    {
        if (!isset(self::$instances["palette-classic-by-name"])) {
            self::$instances["palette-classic-by-name"] = new self("palette-classic-by-name");
        }

        return self::$instances["palette-classic-by-name"];
    }

    public static function continuousViridis(): self
    {
        if (!isset(self::$instances["continuous-viridis"])) {
            self::$instances["continuous-viridis"] = new self("continuous-viridis");
        }

        return self::$instances["continuous-viridis"];
    }

    public static function continuousMagma(): self
    {
        if (!isset(self::$instances["continuous-magma"])) {
            self::$instances["continuous-magma"] = new self("continuous-magma");
        }

        return self::$instances["continuous-magma"];
    }

    public static function continuousPlasma(): self
    {
        if (!isset(self::$instances["continuous-plasma"])) {
            self::$instances["continuous-plasma"] = new self("continuous-plasma");
        }

        return self::$instances["continuous-plasma"];
    }

    public static function continuousInferno(): self
    {
        if (!isset(self::$instances["continuous-inferno"])) {
            self::$instances["continuous-inferno"] = new self("continuous-inferno");
        }

        return self::$instances["continuous-inferno"];
    }

    public static function continuousCividis(): self
    {
        if (!isset(self::$instances["continuous-cividis"])) {
            self::$instances["continuous-cividis"] = new self("continuous-cividis");
        }

        return self::$instances["continuous-cividis"];
    }

    public static function continuousGrYlRd(): self
    {
        if (!isset(self::$instances["continuous-GrYlRd"])) {
            self::$instances["continuous-GrYlRd"] = new self("continuous-GrYlRd");
        }

        return self::$instances["continuous-GrYlRd"];
    }

    public static function continuousRdYlGr(): self
    {
        if (!isset(self::$instances["continuous-RdYlGr"])) {
            self::$instances["continuous-RdYlGr"] = new self("continuous-RdYlGr");
        }

        return self::$instances["continuous-RdYlGr"];
    }

    public static function continuousBlYlRd(): self
    {
        if (!isset(self::$instances["continuous-BlYlRd"])) {
            self::$instances["continuous-BlYlRd"] = new self("continuous-BlYlRd");
        }

        return self::$instances["continuous-BlYlRd"];
    }

    public static function continuousYlRd(): self
    {
        if (!isset(self::$instances["continuous-YlRd"])) {
            self::$instances["continuous-YlRd"] = new self("continuous-YlRd");
        }

        return self::$instances["continuous-YlRd"];
    }

    public static function continuousBlPu(): self
    {
        if (!isset(self::$instances["continuous-BlPu"])) {
            self::$instances["continuous-BlPu"] = new self("continuous-BlPu");
        }

        return self::$instances["continuous-BlPu"];
    }

    public static function continuousYlBl(): self
    {
        if (!isset(self::$instances["continuous-YlBl"])) {
            self::$instances["continuous-YlBl"] = new self("continuous-YlBl");
        }

        return self::$instances["continuous-YlBl"];
    }

    public static function continuousBlues(): self
    {
        if (!isset(self::$instances["continuous-blues"])) {
            self::$instances["continuous-blues"] = new self("continuous-blues");
        }

        return self::$instances["continuous-blues"];
    }

    public static function continuousReds(): self
    {
        if (!isset(self::$instances["continuous-reds"])) {
            self::$instances["continuous-reds"] = new self("continuous-reds");
        }

        return self::$instances["continuous-reds"];
    }

    public static function continuousGreens(): self
    {
        if (!isset(self::$instances["continuous-greens"])) {
            self::$instances["continuous-greens"] = new self("continuous-greens");
        }

        return self::$instances["continuous-greens"];
    }

    public static function continuousPurples(): self
    {
        if (!isset(self::$instances["continuous-purples"])) {
            self::$instances["continuous-purples"] = new self("continuous-purples");
        }

        return self::$instances["continuous-purples"];
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["fixed"])) {
            self::$instances["fixed"] = new self("fixed");
        }

        return self::$instances["fixed"];
    }

    public static function shades(): self
    {
        if (!isset(self::$instances["shades"])) {
            self::$instances["shades"] = new self("shades");
        }

        return self::$instances["shades"];
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

        if ($value === "continuous-viridis") {
            return self::continuousViridis();
        }

        if ($value === "continuous-magma") {
            return self::continuousMagma();
        }

        if ($value === "continuous-plasma") {
            return self::continuousPlasma();
        }

        if ($value === "continuous-inferno") {
            return self::continuousInferno();
        }

        if ($value === "continuous-cividis") {
            return self::continuousCividis();
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
