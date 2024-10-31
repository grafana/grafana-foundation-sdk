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
        if (!isset(self::$instances["none"])) {
            self::$instances["none"] = new self("none");
        }

        return self::$instances["none"];
    }

    public static function normal(): self
    {
        if (!isset(self::$instances["normal"])) {
            self::$instances["normal"] = new self("normal");
        }

        return self::$instances["normal"];
    }

    public static function percent(): self
    {
        if (!isset(self::$instances["percent"])) {
            self::$instances["percent"] = new self("percent");
        }

        return self::$instances["percent"];
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
