<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorExpressionType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryEditorExpressionType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function property(): self
    {
        if (!isset(self::$instances["Property"])) {
            self::$instances["Property"] = new self("property");
        }

        return self::$instances["Property"];
    }

    public static function operator(): self
    {
        if (!isset(self::$instances["Operator"])) {
            self::$instances["Operator"] = new self("operator");
        }

        return self::$instances["Operator"];
    }

    public static function or(): self
    {
        if (!isset(self::$instances["Or"])) {
            self::$instances["Or"] = new self("or");
        }

        return self::$instances["Or"];
    }

    public static function and(): self
    {
        if (!isset(self::$instances["And"])) {
            self::$instances["And"] = new self("and");
        }

        return self::$instances["And"];
    }

    public static function groupBy(): self
    {
        if (!isset(self::$instances["GroupBy"])) {
            self::$instances["GroupBy"] = new self("groupBy");
        }

        return self::$instances["GroupBy"];
    }

    public static function function(): self
    {
        if (!isset(self::$instances["Function"])) {
            self::$instances["Function"] = new self("function");
        }

        return self::$instances["Function"];
    }

    public static function functionParameter(): self
    {
        if (!isset(self::$instances["FunctionParameter"])) {
            self::$instances["FunctionParameter"] = new self("functionParameter");
        }

        return self::$instances["FunctionParameter"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "property") {
            return self::property();
        }

        if ($value === "operator") {
            return self::operator();
        }

        if ($value === "or") {
            return self::or();
        }

        if ($value === "and") {
            return self::and();
        }

        if ($value === "groupBy") {
            return self::groupBy();
        }

        if ($value === "function") {
            return self::function();
        }

        if ($value === "functionParameter") {
            return self::functionParameter();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryEditorExpressionType");
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

