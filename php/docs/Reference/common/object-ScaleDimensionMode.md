---
title: <span class="badge object-type-enum"></span> ScaleDimensionMode
---
# <span class="badge object-type-enum"></span> ScaleDimensionMode

## Definition

```php
final class ScaleDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ScaleDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function linear(): self
    {
        if (!isset(self::$instances["linear"])) {
            self::$instances["linear"] = new self("linear");
        }

        return self::$instances["linear"];
    }

    public static function quad(): self
    {
        if (!isset(self::$instances["quad"])) {
            self::$instances["quad"] = new self("quad");
        }

        return self::$instances["quad"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "linear") {
            return self::linear();
        }

        if ($value === "quad") {
            return self::quad();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScaleDimensionMode");
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
