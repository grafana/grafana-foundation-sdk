---
title: <span class="badge object-type-enum"></span> VisibilityMode
---
# <span class="badge object-type-enum"></span> VisibilityMode

TODO docs

## Definition

```php
final class VisibilityMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VisibilityMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["auto"])) {
            self::$instances["auto"] = new self("auto");
        }

        return self::$instances["auto"];
    }

    public static function never(): self
    {
        if (!isset(self::$instances["never"])) {
            self::$instances["never"] = new self("never");
        }

        return self::$instances["never"];
    }

    public static function always(): self
    {
        if (!isset(self::$instances["always"])) {
            self::$instances["always"] = new self("always");
        }

        return self::$instances["always"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "never") {
            return self::never();
        }

        if ($value === "always") {
            return self::always();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VisibilityMode");
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
