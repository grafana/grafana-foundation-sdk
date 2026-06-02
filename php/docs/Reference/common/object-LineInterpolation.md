---
title: <span class="badge object-type-enum"></span> LineInterpolation
---
# <span class="badge object-type-enum"></span> LineInterpolation

TODO docs

## Definition

```php
final class LineInterpolation implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LineInterpolation>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function linear(): self
    {
        if (!isset(self::$instances["Linear"])) {
            self::$instances["Linear"] = new self("linear");
        }

        return self::$instances["Linear"];
    }

    public static function smooth(): self
    {
        if (!isset(self::$instances["Smooth"])) {
            self::$instances["Smooth"] = new self("smooth");
        }

        return self::$instances["Smooth"];
    }

    public static function stepBefore(): self
    {
        if (!isset(self::$instances["StepBefore"])) {
            self::$instances["StepBefore"] = new self("stepBefore");
        }

        return self::$instances["StepBefore"];
    }

    public static function stepAfter(): self
    {
        if (!isset(self::$instances["StepAfter"])) {
            self::$instances["StepAfter"] = new self("stepAfter");
        }

        return self::$instances["StepAfter"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "linear") {
            return self::linear();
        }

        if ($value === "smooth") {
            return self::smooth();
        }

        if ($value === "stepBefore") {
            return self::stepBefore();
        }

        if ($value === "stepAfter") {
            return self::stepAfter();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LineInterpolation");
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
