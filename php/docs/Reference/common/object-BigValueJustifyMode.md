---
title: <span class="badge object-type-enum"></span> BigValueJustifyMode
---
# <span class="badge object-type-enum"></span> BigValueJustifyMode

TODO docs

## Definition

```php
final class BigValueJustifyMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BigValueJustifyMode>
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

    public static function center(): self
    {
        if (!isset(self::$instances["center"])) {
            self::$instances["center"] = new self("center");
        }

        return self::$instances["center"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "center") {
            return self::center();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BigValueJustifyMode");
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
