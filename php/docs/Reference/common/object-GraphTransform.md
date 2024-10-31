---
title: <span class="badge object-type-enum"></span> GraphTransform
---
# <span class="badge object-type-enum"></span> GraphTransform

TODO docs

## Definition

```php
final class GraphTransform implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, GraphTransform>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function constant(): self
    {
        if (!isset(self::$instances["Constant"])) {
            self::$instances["Constant"] = new self("constant");
        }

        return self::$instances["Constant"];
    }

    public static function negativeY(): self
    {
        if (!isset(self::$instances["NegativeY"])) {
            self::$instances["NegativeY"] = new self("negative-Y");
        }

        return self::$instances["NegativeY"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "constant") {
            return self::constant();
        }

        if ($value === "negative-Y") {
            return self::negativeY();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum GraphTransform");
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
