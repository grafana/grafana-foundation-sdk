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
        if (!isset(self::$instances["linear"])) {
            self::$instances["linear"] = new self("linear");
        }

        return self::$instances["linear"];
    }

    public static function smooth(): self
    {
        if (!isset(self::$instances["smooth"])) {
            self::$instances["smooth"] = new self("smooth");
        }

        return self::$instances["smooth"];
    }

    public static function stepBefore(): self
    {
        if (!isset(self::$instances["stepBefore"])) {
            self::$instances["stepBefore"] = new self("stepBefore");
        }

        return self::$instances["stepBefore"];
    }

    public static function stepAfter(): self
    {
        if (!isset(self::$instances["stepAfter"])) {
            self::$instances["stepAfter"] = new self("stepAfter");
        }

        return self::$instances["stepAfter"];
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
