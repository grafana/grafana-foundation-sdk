---
title: <span class="badge object-type-enum"></span> VizOrientation
---
# <span class="badge object-type-enum"></span> VizOrientation

TODO docs

## Definition

```php
final class VizOrientation implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VizOrientation>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function vertical(): self
    {
        if (!isset(self::$instances["Vertical"])) {
            self::$instances["Vertical"] = new self("vertical");
        }

        return self::$instances["Vertical"];
    }

    public static function horizontal(): self
    {
        if (!isset(self::$instances["Horizontal"])) {
            self::$instances["Horizontal"] = new self("horizontal");
        }

        return self::$instances["Horizontal"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "vertical") {
            return self::vertical();
        }

        if ($value === "horizontal") {
            return self::horizontal();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VizOrientation");
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
