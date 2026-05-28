---
title: <span class="badge object-type-enum"></span> RepeatMode
---
# <span class="badge object-type-enum"></span> RepeatMode

other repeat modes will be added in the future: label, frame

## Definition

```php
final class RepeatMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, RepeatMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function variable(): self
    {
        if (!isset(self::$instances["variable"])) {
            self::$instances["variable"] = new self("variable");
        }

        return self::$instances["variable"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "variable") {
            return self::variable();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum RepeatMode");
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
