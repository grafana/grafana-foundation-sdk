---
title: <span class="badge object-type-enum"></span> PreprocessorType
---
# <span class="badge object-type-enum"></span> PreprocessorType

Types of pre-processor available. Defined by the metric.

## Definition

```php
final class PreprocessorType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PreprocessorType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function rate(): self
    {
        if (!isset(self::$instances["Rate"])) {
            self::$instances["Rate"] = new self("rate");
        }

        return self::$instances["Rate"];
    }

    public static function delta(): self
    {
        if (!isset(self::$instances["Delta"])) {
            self::$instances["Delta"] = new self("delta");
        }

        return self::$instances["Delta"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "rate") {
            return self::rate();
        }

        if ($value === "delta") {
            return self::delta();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PreprocessorType");
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
