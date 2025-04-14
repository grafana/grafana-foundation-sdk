---
title: <span class="badge object-type-enum"></span> DataqueryErrorSource
---
# <span class="badge object-type-enum"></span> DataqueryErrorSource

## Definition

```php
final class DataqueryErrorSource implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DataqueryErrorSource>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function plugin(): self
    {
        if (!isset(self::$instances["Plugin"])) {
            self::$instances["Plugin"] = new self("plugin");
        }

        return self::$instances["Plugin"];
    }

    public static function downstream(): self
    {
        if (!isset(self::$instances["Downstream"])) {
            self::$instances["Downstream"] = new self("downstream");
        }

        return self::$instances["Downstream"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "plugin") {
            return self::plugin();
        }

        if ($value === "downstream") {
            return self::downstream();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DataqueryErrorSource");
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
