---
title: <span class="badge object-type-enum"></span> MovingAverageModel
---
# <span class="badge object-type-enum"></span> MovingAverageModel

## Definition

```php
final class MovingAverageModel implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MovingAverageModel>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function simple(): self
    {
        if (!isset(self::$instances["simple"])) {
            self::$instances["simple"] = new self("simple");
        }

        return self::$instances["simple"];
    }

    public static function linear(): self
    {
        if (!isset(self::$instances["linear"])) {
            self::$instances["linear"] = new self("linear");
        }

        return self::$instances["linear"];
    }

    public static function ewma(): self
    {
        if (!isset(self::$instances["ewma"])) {
            self::$instances["ewma"] = new self("ewma");
        }

        return self::$instances["ewma"];
    }

    public static function holt(): self
    {
        if (!isset(self::$instances["holt"])) {
            self::$instances["holt"] = new self("holt");
        }

        return self::$instances["holt"];
    }

    public static function holtWinters(): self
    {
        if (!isset(self::$instances["holt_winters"])) {
            self::$instances["holt_winters"] = new self("holt_winters");
        }

        return self::$instances["holt_winters"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "simple") {
            return self::simple();
        }

        if ($value === "linear") {
            return self::linear();
        }

        if ($value === "ewma") {
            return self::ewma();
        }

        if ($value === "holt") {
            return self::holt();
        }

        if ($value === "holt_winters") {
            return self::holtWinters();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MovingAverageModel");
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
