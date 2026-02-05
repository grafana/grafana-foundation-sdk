---
title: <span class="badge object-type-enum"></span> BarAlignment
---
# <span class="badge object-type-enum"></span> BarAlignment

TODO docs

## Definition

```php
final class BarAlignment implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, BarAlignment>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function before(): self
    {
        if (!isset(self::$instances["Before"])) {
            self::$instances["Before"] = new self(-1);
        }

        return self::$instances["Before"];
    }

    public static function center(): self
    {
        if (!isset(self::$instances["Center"])) {
            self::$instances["Center"] = new self(0);
        }

        return self::$instances["Center"];
    }

    public static function after(): self
    {
        if (!isset(self::$instances["After"])) {
            self::$instances["After"] = new self(1);
        }

        return self::$instances["After"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === -1) {
            return self::before();
        }

        if ($value === 0) {
            return self::center();
        }

        if ($value === 1) {
            return self::after();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BarAlignment");
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
