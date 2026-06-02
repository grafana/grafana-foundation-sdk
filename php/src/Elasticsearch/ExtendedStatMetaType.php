<?php

namespace Grafana\Foundation\Elasticsearch;

final class ExtendedStatMetaType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ExtendedStatMetaType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function avg(): self
    {
        if (!isset(self::$instances["Avg"])) {
            self::$instances["Avg"] = new self("avg");
        }

        return self::$instances["Avg"];
    }

    public static function min(): self
    {
        if (!isset(self::$instances["Min"])) {
            self::$instances["Min"] = new self("min");
        }

        return self::$instances["Min"];
    }

    public static function max(): self
    {
        if (!isset(self::$instances["Max"])) {
            self::$instances["Max"] = new self("max");
        }

        return self::$instances["Max"];
    }

    public static function sum(): self
    {
        if (!isset(self::$instances["Sum"])) {
            self::$instances["Sum"] = new self("sum");
        }

        return self::$instances["Sum"];
    }

    public static function count(): self
    {
        if (!isset(self::$instances["Count"])) {
            self::$instances["Count"] = new self("count");
        }

        return self::$instances["Count"];
    }

    public static function stdDeviation(): self
    {
        if (!isset(self::$instances["StdDeviation"])) {
            self::$instances["StdDeviation"] = new self("std_deviation");
        }

        return self::$instances["StdDeviation"];
    }

    public static function stdDeviationBoundsUpper(): self
    {
        if (!isset(self::$instances["StdDeviationBoundsUpper"])) {
            self::$instances["StdDeviationBoundsUpper"] = new self("std_deviation_bounds_upper");
        }

        return self::$instances["StdDeviationBoundsUpper"];
    }

    public static function stdDeviationBoundsLower(): self
    {
        if (!isset(self::$instances["StdDeviationBoundsLower"])) {
            self::$instances["StdDeviationBoundsLower"] = new self("std_deviation_bounds_lower");
        }

        return self::$instances["StdDeviationBoundsLower"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "avg") {
            return self::avg();
        }

        if ($value === "min") {
            return self::min();
        }

        if ($value === "max") {
            return self::max();
        }

        if ($value === "sum") {
            return self::sum();
        }

        if ($value === "count") {
            return self::count();
        }

        if ($value === "std_deviation") {
            return self::stdDeviation();
        }

        if ($value === "std_deviation_bounds_upper") {
            return self::stdDeviationBoundsUpper();
        }

        if ($value === "std_deviation_bounds_lower") {
            return self::stdDeviationBoundsLower();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ExtendedStatMetaType");
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

