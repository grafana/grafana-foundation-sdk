---
title: <span class="badge object-type-enum"></span> ScaleOrientation
---
# <span class="badge object-type-enum"></span> ScaleOrientation

TODO docs

## Definition

```php
final class ScaleOrientation implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, ScaleOrientation>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function horizontal(): self
    {
        if (!isset(self::$instances["Horizontal"])) {
            self::$instances["Horizontal"] = new self(0);
        }

        return self::$instances["Horizontal"];
    }

    public static function vertical(): self
    {
        if (!isset(self::$instances["Vertical"])) {
            self::$instances["Vertical"] = new self(1);
        }

        return self::$instances["Vertical"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::horizontal();
        }

        if ($value === 1) {
            return self::vertical();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ScaleOrientation");
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
