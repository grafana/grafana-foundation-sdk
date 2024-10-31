---
title: <span class="badge object-type-enum"></span> ParcaQueryType
---
# <span class="badge object-type-enum"></span> ParcaQueryType

## Definition

```php
final class ParcaQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ParcaQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function metrics(): self
    {
        if (!isset(self::$instances["metrics"])) {
            self::$instances["metrics"] = new self("metrics");
        }

        return self::$instances["metrics"];
    }

    public static function profile(): self
    {
        if (!isset(self::$instances["profile"])) {
            self::$instances["profile"] = new self("profile");
        }

        return self::$instances["profile"];
    }

    public static function both(): self
    {
        if (!isset(self::$instances["both"])) {
            self::$instances["both"] = new self("both");
        }

        return self::$instances["both"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "metrics") {
            return self::metrics();
        }

        if ($value === "profile") {
            return self::profile();
        }

        if ($value === "both") {
            return self::both();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ParcaQueryType");
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
