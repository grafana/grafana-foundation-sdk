---
title: <span class="badge object-type-enum"></span> ResourceDimensionMode
---
# <span class="badge object-type-enum"></span> ResourceDimensionMode

## Definition

```php
final class ResourceDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ResourceDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["fixed"])) {
            self::$instances["fixed"] = new self("fixed");
        }

        return self::$instances["fixed"];
    }

    public static function field(): self
    {
        if (!isset(self::$instances["field"])) {
            self::$instances["field"] = new self("field");
        }

        return self::$instances["field"];
    }

    public static function mapping(): self
    {
        if (!isset(self::$instances["mapping"])) {
            self::$instances["mapping"] = new self("mapping");
        }

        return self::$instances["mapping"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fixed") {
            return self::fixed();
        }

        if ($value === "field") {
            return self::field();
        }

        if ($value === "mapping") {
            return self::mapping();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ResourceDimensionMode");
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
