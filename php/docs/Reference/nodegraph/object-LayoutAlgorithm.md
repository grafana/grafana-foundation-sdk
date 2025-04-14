---
title: <span class="badge object-type-enum"></span> LayoutAlgorithm
---
# <span class="badge object-type-enum"></span> LayoutAlgorithm

## Definition

```php
final class LayoutAlgorithm implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LayoutAlgorithm>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function layered(): self
    {
        if (!isset(self::$instances["layered"])) {
            self::$instances["layered"] = new self("layered");
        }

        return self::$instances["layered"];
    }

    public static function force(): self
    {
        if (!isset(self::$instances["force"])) {
            self::$instances["force"] = new self("force");
        }

        return self::$instances["force"];
    }

    public static function grid(): self
    {
        if (!isset(self::$instances["grid"])) {
            self::$instances["grid"] = new self("grid");
        }

        return self::$instances["grid"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "layered") {
            return self::layered();
        }

        if ($value === "force") {
            return self::force();
        }

        if ($value === "grid") {
            return self::grid();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LayoutAlgorithm");
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
