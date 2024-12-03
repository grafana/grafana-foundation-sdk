---
title: <span class="badge object-type-enum"></span> QueryEditorPropertyType
---
# <span class="badge object-type-enum"></span> QueryEditorPropertyType

## Definition

```php
final class QueryEditorPropertyType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryEditorPropertyType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function string(): self
    {
        if (!isset(self::$instances["string"])) {
            self::$instances["string"] = new self("string");
        }

        return self::$instances["string"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "string") {
            return self::string();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryEditorPropertyType");
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
