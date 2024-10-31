---
title: <span class="badge object-type-enum"></span> ScaleDirection
---
# <span class="badge object-type-enum"></span> ScaleDirection

TODO docs

## Definition

```php
final class ScaleDirection implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, ScaleDirection>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function up(): self
    {
        if (!isset(self::$instances["Up"])) {
            self::$instances["Up"] = new self(1);
        }

        return self::$instances["Up"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["Right"])) {
            self::$instances["Right"] = new self(1);
        }

        return self::$instances["Right"];
    }

    public static function down(): self
    {
        if (!isset(self::$instances["Down"])) {
            self::$instances["Down"] = new self(-1);
        }

        return self::$instances["Down"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["Left"])) {
            self::$instances["Left"] = new self(-1);
        }

        return self::$instances["Left"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 1) {
            return self::up();
        }

        if ($value === 1) {
            return self::right();
        }

        if ($value === -1) {
            return self::down();
        }

        if ($value === -1) {
            return self::left();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScaleDirection");
    }

    public function jsonSerialize(): int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

```
