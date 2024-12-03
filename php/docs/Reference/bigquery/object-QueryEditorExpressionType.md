---
title: <span class="badge object-type-enum"></span> QueryEditorExpressionType
---
# <span class="badge object-type-enum"></span> QueryEditorExpressionType

## Definition

```php
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

    public static function groupBy(): self
    {
        if (!isset(self::$instances["groupBy"])) {
            self::$instances["groupBy"] = new self("groupBy");
        }

        return self::$instances["groupBy"];
    }

    public static function function(): self
    {
        if (!isset(self::$instances["function"])) {
            self::$instances["function"] = new self("function");
        }

        return self::$instances["function"];
    }

    public static function functionParameter(): self
    {
        if (!isset(self::$instances["functionParameter"])) {
            self::$instances["functionParameter"] = new self("functionParameter");
        }

        return self::$instances["functionParameter"];
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

```
