---
title: <span class="badge object-type-enum"></span> QueryEditorArrayExpressionType
---
# <span class="badge object-type-enum"></span> QueryEditorArrayExpressionType

## Definition

```php
final class QueryEditorArrayExpressionType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryEditorArrayExpressionType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function and(): self
    {
        if (!isset(self::$instances["And"])) {
            self::$instances["And"] = new self("and");
        }

        return self::$instances["And"];
    }

    public static function or(): self
    {
        if (!isset(self::$instances["Or"])) {
            self::$instances["Or"] = new self("or");
        }

        return self::$instances["Or"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "and") {
            return self::and();
        }

        if ($value === "or") {
            return self::or();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryEditorArrayExpressionType");
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
