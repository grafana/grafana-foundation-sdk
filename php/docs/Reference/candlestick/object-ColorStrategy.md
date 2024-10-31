---
title: <span class="badge object-type-enum"></span> ColorStrategy
---
# <span class="badge object-type-enum"></span> ColorStrategy

## Definition

```php
final class ColorStrategy implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ColorStrategy>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function openClose(): self
    {
        if (!isset(self::$instances["OpenClose"])) {
            self::$instances["OpenClose"] = new self("open-close");
        }

        return self::$instances["OpenClose"];
    }

    public static function closeClose(): self
    {
        if (!isset(self::$instances["CloseClose"])) {
            self::$instances["CloseClose"] = new self("close-close");
        }

        return self::$instances["CloseClose"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "open-close") {
            return self::openClose();
        }

        if ($value === "close-close") {
            return self::closeClose();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ColorStrategy");
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
