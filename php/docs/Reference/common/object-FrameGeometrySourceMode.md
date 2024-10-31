---
title: <span class="badge object-type-enum"></span> FrameGeometrySourceMode
---
# <span class="badge object-type-enum"></span> FrameGeometrySourceMode

## Definition

```php
final class FrameGeometrySourceMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, FrameGeometrySourceMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function geohash(): self
    {
        if (!isset(self::$instances["Geohash"])) {
            self::$instances["Geohash"] = new self("geohash");
        }

        return self::$instances["Geohash"];
    }

    public static function coords(): self
    {
        if (!isset(self::$instances["Coords"])) {
            self::$instances["Coords"] = new self("coords");
        }

        return self::$instances["Coords"];
    }

    public static function lookup(): self
    {
        if (!isset(self::$instances["Lookup"])) {
            self::$instances["Lookup"] = new self("lookup");
        }

        return self::$instances["Lookup"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "geohash") {
            return self::geohash();
        }

        if ($value === "coords") {
            return self::coords();
        }

        if ($value === "lookup") {
            return self::lookup();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FrameGeometrySourceMode");
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
