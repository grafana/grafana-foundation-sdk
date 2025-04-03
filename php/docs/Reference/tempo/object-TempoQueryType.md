---
title: <span class="badge object-type-enum"></span> TempoQueryType
---
# <span class="badge object-type-enum"></span> TempoQueryType

## Definition

```php
final class TempoQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TempoQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function traceql(): self
    {
        if (!isset(self::$instances["traceql"])) {
            self::$instances["traceql"] = new self("traceql");
        }

        return self::$instances["traceql"];
    }

    public static function traceqlSearch(): self
    {
        if (!isset(self::$instances["traceqlSearch"])) {
            self::$instances["traceqlSearch"] = new self("traceqlSearch");
        }

        return self::$instances["traceqlSearch"];
    }

    public static function serviceMap(): self
    {
        if (!isset(self::$instances["serviceMap"])) {
            self::$instances["serviceMap"] = new self("serviceMap");
        }

        return self::$instances["serviceMap"];
    }

    public static function upload(): self
    {
        if (!isset(self::$instances["upload"])) {
            self::$instances["upload"] = new self("upload");
        }

        return self::$instances["upload"];
    }

    public static function nativeSearch(): self
    {
        if (!isset(self::$instances["nativeSearch"])) {
            self::$instances["nativeSearch"] = new self("nativeSearch");
        }

        return self::$instances["nativeSearch"];
    }

    public static function traceId(): self
    {
        if (!isset(self::$instances["traceId"])) {
            self::$instances["traceId"] = new self("traceId");
        }

        return self::$instances["traceId"];
    }

    public static function clear(): self
    {
        if (!isset(self::$instances["clear"])) {
            self::$instances["clear"] = new self("clear");
        }

        return self::$instances["clear"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "traceql") {
            return self::traceql();
        }

        if ($value === "traceqlSearch") {
            return self::traceqlSearch();
        }

        if ($value === "serviceMap") {
            return self::serviceMap();
        }

        if ($value === "upload") {
            return self::upload();
        }

        if ($value === "nativeSearch") {
            return self::nativeSearch();
        }

        if ($value === "traceId") {
            return self::traceId();
        }

        if ($value === "clear") {
            return self::clear();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TempoQueryType");
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
