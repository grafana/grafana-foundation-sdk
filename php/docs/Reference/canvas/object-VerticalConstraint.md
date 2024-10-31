---
title: <span class="badge object-type-enum"></span> VerticalConstraint
---
# <span class="badge object-type-enum"></span> VerticalConstraint

## Definition

```php
final class VerticalConstraint implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VerticalConstraint>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function top(): self
    {
        if (!isset(self::$instances["Top"])) {
            self::$instances["Top"] = new self("top");
        }

        return self::$instances["Top"];
    }

    public static function bottom(): self
    {
        if (!isset(self::$instances["Bottom"])) {
            self::$instances["Bottom"] = new self("bottom");
        }

        return self::$instances["Bottom"];
    }

    public static function topBottom(): self
    {
        if (!isset(self::$instances["TopBottom"])) {
            self::$instances["TopBottom"] = new self("topbottom");
        }

        return self::$instances["TopBottom"];
    }

    public static function center(): self
    {
        if (!isset(self::$instances["Center"])) {
            self::$instances["Center"] = new self("center");
        }

        return self::$instances["Center"];
    }

    public static function scale(): self
    {
        if (!isset(self::$instances["Scale"])) {
            self::$instances["Scale"] = new self("scale");
        }

        return self::$instances["Scale"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "top") {
            return self::top();
        }

        if ($value === "bottom") {
            return self::bottom();
        }

        if ($value === "topbottom") {
            return self::topBottom();
        }

        if ($value === "center") {
            return self::center();
        }

        if ($value === "scale") {
            return self::scale();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VerticalConstraint");
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
