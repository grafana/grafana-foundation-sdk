<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorExpressionType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BuilderQueryEditorExpressionType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function property(): self
    {
        if (!isset(self::$instances["property"])) {
            self::$instances["property"] = new self("property");
        }

        return self::$instances["property"];
    }

    public static function operator(): self
    {
        if (!isset(self::$instances["operator"])) {
            self::$instances["operator"] = new self("operator");
        }

        return self::$instances["operator"];
    }

    public static function reduce(): self
    {
        if (!isset(self::$instances["reduce"])) {
            self::$instances["reduce"] = new self("reduce");
        }

        return self::$instances["reduce"];
    }

    public static function functionParameter(): self
    {
        if (!isset(self::$instances["function_parameter"])) {
            self::$instances["function_parameter"] = new self("function_parameter");
        }

        return self::$instances["function_parameter"];
    }

    public static function groupBy(): self
    {
        if (!isset(self::$instances["group_by"])) {
            self::$instances["group_by"] = new self("group_by");
        }

        return self::$instances["group_by"];
    }

    public static function or(): self
    {
        if (!isset(self::$instances["or"])) {
            self::$instances["or"] = new self("or");
        }

        return self::$instances["or"];
    }

    public static function and(): self
    {
        if (!isset(self::$instances["and"])) {
            self::$instances["and"] = new self("and");
        }

        return self::$instances["and"];
    }

    public static function orderBy(): self
    {
        if (!isset(self::$instances["order_by"])) {
            self::$instances["order_by"] = new self("order_by");
        }

        return self::$instances["order_by"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "property") {
            return self::property();
        }

        if ($value === "operator") {
            return self::operator();
        }

        if ($value === "reduce") {
            return self::reduce();
        }

        if ($value === "function_parameter") {
            return self::functionParameter();
        }

        if ($value === "group_by") {
            return self::groupBy();
        }

        if ($value === "or") {
            return self::or();
        }

        if ($value === "and") {
            return self::and();
        }

        if ($value === "order_by") {
            return self::orderBy();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BuilderQueryEditorExpressionType");
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

