<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeThresholdConditionsEvaluatorType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ExprTypeThresholdConditionsEvaluatorType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function gt(): self
    {
        if (!isset(self::$instances["Gt"])) {
            self::$instances["Gt"] = new self("gt");
        }

        return self::$instances["Gt"];
    }

    public static function lt(): self
    {
        if (!isset(self::$instances["Lt"])) {
            self::$instances["Lt"] = new self("lt");
        }

        return self::$instances["Lt"];
    }

    public static function eq(): self
    {
        if (!isset(self::$instances["Eq"])) {
            self::$instances["Eq"] = new self("eq");
        }

        return self::$instances["Eq"];
    }

    public static function ne(): self
    {
        if (!isset(self::$instances["Ne"])) {
            self::$instances["Ne"] = new self("ne");
        }

        return self::$instances["Ne"];
    }

    public static function gte(): self
    {
        if (!isset(self::$instances["Gte"])) {
            self::$instances["Gte"] = new self("gte");
        }

        return self::$instances["Gte"];
    }

    public static function lte(): self
    {
        if (!isset(self::$instances["Lte"])) {
            self::$instances["Lte"] = new self("lte");
        }

        return self::$instances["Lte"];
    }

    public static function withinRange(): self
    {
        if (!isset(self::$instances["WithinRange"])) {
            self::$instances["WithinRange"] = new self("within_range");
        }

        return self::$instances["WithinRange"];
    }

    public static function outsideRange(): self
    {
        if (!isset(self::$instances["OutsideRange"])) {
            self::$instances["OutsideRange"] = new self("outside_range");
        }

        return self::$instances["OutsideRange"];
    }

    public static function withinRangeIncluded(): self
    {
        if (!isset(self::$instances["WithinRangeIncluded"])) {
            self::$instances["WithinRangeIncluded"] = new self("within_range_included");
        }

        return self::$instances["WithinRangeIncluded"];
    }

    public static function outsideRangeIncluded(): self
    {
        if (!isset(self::$instances["OutsideRangeIncluded"])) {
            self::$instances["OutsideRangeIncluded"] = new self("outside_range_included");
        }

        return self::$instances["OutsideRangeIncluded"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "gt") {
            return self::gt();
        }

        if ($value === "lt") {
            return self::lt();
        }

        if ($value === "eq") {
            return self::eq();
        }

        if ($value === "ne") {
            return self::ne();
        }

        if ($value === "gte") {
            return self::gte();
        }

        if ($value === "lte") {
            return self::lte();
        }

        if ($value === "within_range") {
            return self::withinRange();
        }

        if ($value === "outside_range") {
            return self::outsideRange();
        }

        if ($value === "within_range_included") {
            return self::withinRangeIncluded();
        }

        if ($value === "outside_range_included") {
            return self::outsideRangeIncluded();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ExprTypeThresholdConditionsEvaluatorType");
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

