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
        if (!isset(self::$instances["Simple"])) {
            self::$instances["Simple"] = new self("simple");
        }

        return self::$instances["Simple"];
    }

    public static function linear(): self
    {
        if (!isset(self::$instances["Linear"])) {
            self::$instances["Linear"] = new self("linear");
        }

        return self::$instances["Linear"];
    }

    public static function ewma(): self
    {
        if (!isset(self::$instances["Ewma"])) {
            self::$instances["Ewma"] = new self("ewma");
        }

        return self::$instances["Ewma"];
    }

    public static function holt(): self
    {
        if (!isset(self::$instances["Holt"])) {
            self::$instances["Holt"] = new self("holt");
        }

        return self::$instances["Holt"];
    }

    public static function holtWinters(): self
    {
        if (!isset(self::$instances["HoltWinters"])) {
            self::$instances["HoltWinters"] = new self("holt_winters");
        }

        return self::$instances["HoltWinters"];
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
