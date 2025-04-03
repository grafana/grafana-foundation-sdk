---
title: <span class="badge object-type-enum"></span> TooltipMode
---
# <span class="badge object-type-enum"></span> TooltipMode

## Definition

```php
final class TooltipMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TooltipMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function details(): self
    {
        if (!isset(self::$instances["Details"])) {
            self::$instances["Details"] = new self("details");
        }

        return self::$instances["Details"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "details") {
            return self::details();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TooltipMode");
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
