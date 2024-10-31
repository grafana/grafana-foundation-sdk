---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).

## Definition

```php
final class VariableHide implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, VariableHide>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function dontHide(): self
    {
        if (!isset(self::$instances["dontHide"])) {
            self::$instances["dontHide"] = new self(0);
        }

        return self::$instances["dontHide"];
    }

    public static function hideLabel(): self
    {
        if (!isset(self::$instances["hideLabel"])) {
            self::$instances["hideLabel"] = new self(1);
        }

        return self::$instances["hideLabel"];
    }

    public static function hideVariable(): self
    {
        if (!isset(self::$instances["hideVariable"])) {
            self::$instances["hideVariable"] = new self(2);
        }

        return self::$instances["hideVariable"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::dontHide();
        }

        if ($value === 1) {
            return self::hideLabel();
        }

        if ($value === 2) {
            return self::hideVariable();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableHide");
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
