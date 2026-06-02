<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
final class GraphGradientMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GraphGradientMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function opacity(): self
    {
        if (!isset(self::$instances["Opacity"])) {
            self::$instances["Opacity"] = new self("opacity");
        }

        return self::$instances["Opacity"];
    }

    public static function hue(): self
    {
        if (!isset(self::$instances["Hue"])) {
            self::$instances["Hue"] = new self("hue");
        }

        return self::$instances["Hue"];
    }

    public static function scheme(): self
    {
        if (!isset(self::$instances["Scheme"])) {
            self::$instances["Scheme"] = new self("scheme");
        }

        return self::$instances["Scheme"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "opacity") {
            return self::opacity();
        }

        if ($value === "hue") {
            return self::hue();
        }

        if ($value === "scheme") {
            return self::scheme();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GraphGradientMode");
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

