---
title: <span class="badge object-type-enum"></span> StackingMode
---
# <span class="badge object-type-enum"></span> StackingMode

TODO docs

## Definition

```php
final class StackingMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, StackingMode>
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

    public static function normal(): self
    {
        if (!isset(self::$instances["Normal"])) {
            self::$instances["Normal"] = new self("normal");
        }

        return self::$instances["Normal"];
    }

    public static function percent(): self
    {
        if (!isset(self::$instances["Percent"])) {
            self::$instances["Percent"] = new self("percent");
        }

        return self::$instances["Percent"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "normal") {
            return self::normal();
        }

        if ($value === "percent") {
            return self::percent();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum StackingMode");
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
