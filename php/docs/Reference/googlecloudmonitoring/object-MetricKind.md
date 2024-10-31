---
title: <span class="badge object-type-enum"></span> MetricKind
---
# <span class="badge object-type-enum"></span> MetricKind

## Definition

```php
final class MetricKind implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MetricKind>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function mETRICKINDUNSPECIFIED(): self
    {
        if (!isset(self::$instances["METRIC_KIND_UNSPECIFIED"])) {
            self::$instances["METRIC_KIND_UNSPECIFIED"] = new self("METRIC_KIND_UNSPECIFIED");
        }

        return self::$instances["METRIC_KIND_UNSPECIFIED"];
    }

    public static function gAUGE(): self
    {
        if (!isset(self::$instances["GAUGE"])) {
            self::$instances["GAUGE"] = new self("GAUGE");
        }

        return self::$instances["GAUGE"];
    }

    public static function dELTA(): self
    {
        if (!isset(self::$instances["DELTA"])) {
            self::$instances["DELTA"] = new self("DELTA");
        }

        return self::$instances["DELTA"];
    }

    public static function cUMULATIVE(): self
    {
        if (!isset(self::$instances["CUMULATIVE"])) {
            self::$instances["CUMULATIVE"] = new self("CUMULATIVE");
        }

        return self::$instances["CUMULATIVE"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "METRIC_KIND_UNSPECIFIED") {
            return self::mETRICKINDUNSPECIFIED();
        }

        if ($value === "GAUGE") {
            return self::gAUGE();
        }

        if ($value === "DELTA") {
            return self::dELTA();
        }

        if ($value === "CUMULATIVE") {
            return self::cUMULATIVE();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricKind");
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
