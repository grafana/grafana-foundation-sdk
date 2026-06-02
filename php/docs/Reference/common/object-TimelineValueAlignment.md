---
title: <span class="badge object-type-enum"></span> TimelineValueAlignment
---
# <span class="badge object-type-enum"></span> TimelineValueAlignment

Controls the value alignment in the TimelineChart component

## Definition

```php
final class TimelineValueAlignment implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TimelineValueAlignment>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function center(): self
    {
        if (!isset(self::$instances["Center"])) {
            self::$instances["Center"] = new self("center");
        }

        return self::$instances["Center"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["Left"])) {
            self::$instances["Left"] = new self("left");
        }

        return self::$instances["Left"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["Right"])) {
            self::$instances["Right"] = new self("right");
        }

        return self::$instances["Right"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "center") {
            return self::center();
        }

        if ($value === "left") {
            return self::left();
        }

        if ($value === "right") {
            return self::right();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TimelineValueAlignment");
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
