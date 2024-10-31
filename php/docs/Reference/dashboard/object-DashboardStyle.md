---
title: <span class="badge object-type-enum"></span> DashboardStyle
---
# <span class="badge object-type-enum"></span> DashboardStyle

## Definition

```php
final class DashboardStyle implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DashboardStyle>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function light(): self
    {
        if (!isset(self::$instances["Light"])) {
            self::$instances["Light"] = new self("light");
        }

        return self::$instances["Light"];
    }

    public static function dark(): self
    {
        if (!isset(self::$instances["Dark"])) {
            self::$instances["Dark"] = new self("dark");
        }

        return self::$instances["Dark"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "light") {
            return self::light();
        }

        if ($value === "dark") {
            return self::dark();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DashboardStyle");
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
