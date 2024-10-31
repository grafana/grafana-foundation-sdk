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
        if (!isset(self::$instances["zero"])) {
            self::$instances["zero"] = new self("zero");
        }

        return self::$instances["zero"];
    }

    public static function coords(): self
    {
        if (!isset(self::$instances["coords"])) {
            self::$instances["coords"] = new self("coords");
        }

        return self::$instances["coords"];
    }

    public static function fit(): self
    {
        if (!isset(self::$instances["fit"])) {
            self::$instances["fit"] = new self("fit");
        }

        return self::$instances["fit"];
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
