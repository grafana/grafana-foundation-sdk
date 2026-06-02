<?php

namespace Grafana\Foundation\Common;

final class HeatmapCalculationMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, HeatmapCalculationMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function size(): self
    {
        if (!isset(self::$instances["Size"])) {
            self::$instances["Size"] = new self("size");
        }

        return self::$instances["Size"];
    }

    public static function count(): self
    {
        if (!isset(self::$instances["Count"])) {
            self::$instances["Count"] = new self("count");
        }

        return self::$instances["Count"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "size") {
            return self::size();
        }

        if ($value === "count") {
            return self::count();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum HeatmapCalculationMode");
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

