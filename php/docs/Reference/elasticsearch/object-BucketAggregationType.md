---
title: <span class="badge object-type-enum"></span> BucketAggregationType
---
# <span class="badge object-type-enum"></span> BucketAggregationType

## Definition

```php
final class BucketAggregationType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BucketAggregationType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function terms(): self
    {
        if (!isset(self::$instances["terms"])) {
            self::$instances["terms"] = new self("terms");
        }

        return self::$instances["terms"];
    }

    public static function filters(): self
    {
        if (!isset(self::$instances["filters"])) {
            self::$instances["filters"] = new self("filters");
        }

        return self::$instances["filters"];
    }

    public static function geohashGrid(): self
    {
        if (!isset(self::$instances["geohash_grid"])) {
            self::$instances["geohash_grid"] = new self("geohash_grid");
        }

        return self::$instances["geohash_grid"];
    }

    public static function dateHistogram(): self
    {
        if (!isset(self::$instances["date_histogram"])) {
            self::$instances["date_histogram"] = new self("date_histogram");
        }

        return self::$instances["date_histogram"];
    }

    public static function histogram(): self
    {
        if (!isset(self::$instances["histogram"])) {
            self::$instances["histogram"] = new self("histogram");
        }

        return self::$instances["histogram"];
    }

    public static function nested(): self
    {
        if (!isset(self::$instances["nested"])) {
            self::$instances["nested"] = new self("nested");
        }

        return self::$instances["nested"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "terms") {
            return self::terms();
        }

        if ($value === "filters") {
            return self::filters();
        }

        if ($value === "geohash_grid") {
            return self::geohashGrid();
        }

        if ($value === "date_histogram") {
            return self::dateHistogram();
        }

        if ($value === "histogram") {
            return self::histogram();
        }

        if ($value === "nested") {
            return self::nested();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BucketAggregationType");
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
