---
title: <span class="badge object-type-enum"></span> MapCenterID
---
# <span class="badge object-type-enum"></span> MapCenterID

## Definition

```php
final class MapCenterID implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MapCenterID>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function zero(): self
    {
        if (!isset(self::$instances["Zero"])) {
            self::$instances["Zero"] = new self("zero");
        }

        return self::$instances["Zero"];
    }

    public static function coords(): self
    {
        if (!isset(self::$instances["Coords"])) {
            self::$instances["Coords"] = new self("coords");
        }

        return self::$instances["Coords"];
    }

    public static function fit(): self
    {
        if (!isset(self::$instances["Fit"])) {
            self::$instances["Fit"] = new self("fit");
        }

        return self::$instances["Fit"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "zero") {
            return self::zero();
        }

        if ($value === "coords") {
            return self::coords();
        }

        if ($value === "fit") {
            return self::fit();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MapCenterID");
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
