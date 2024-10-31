---
title: <span class="badge object-type-enum"></span> ScaleDistribution
---
# <span class="badge object-type-enum"></span> ScaleDistribution

TODO docs

## Definition

```php
final class ScaleDistribution implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ScaleDistribution>
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

    public static function log(): self
    {
        if (!isset(self::$instances["log"])) {
            self::$instances["log"] = new self("log");
        }

        return self::$instances["log"];
    }

    public static function ordinal(): self
    {
        if (!isset(self::$instances["ordinal"])) {
            self::$instances["ordinal"] = new self("ordinal");
        }

        return self::$instances["ordinal"];
    }

    public static function symlog(): self
    {
        if (!isset(self::$instances["symlog"])) {
            self::$instances["symlog"] = new self("symlog");
        }

        return self::$instances["symlog"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "linear") {
            return self::linear();
        }

        if ($value === "log") {
            return self::log();
        }

        if ($value === "ordinal") {
            return self::ordinal();
        }

        if ($value === "symlog") {
            return self::symlog();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScaleDistribution");
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
