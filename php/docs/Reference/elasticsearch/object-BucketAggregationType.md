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
        if (!isset(self::$instances["Terms"])) {
            self::$instances["Terms"] = new self("terms");
        }

        return self::$instances["Terms"];
    }

    public static function filters(): self
    {
        if (!isset(self::$instances["Filters"])) {
            self::$instances["Filters"] = new self("filters");
        }

        return self::$instances["Filters"];
    }

    public static function geohashGrid(): self
    {
        if (!isset(self::$instances["GeohashGrid"])) {
            self::$instances["GeohashGrid"] = new self("geohash_grid");
        }

        return self::$instances["GeohashGrid"];
    }

    public static function dateHistogram(): self
    {
        if (!isset(self::$instances["DateHistogram"])) {
            self::$instances["DateHistogram"] = new self("date_histogram");
        }

        return self::$instances["DateHistogram"];
    }

    public static function histogram(): self
    {
        if (!isset(self::$instances["Histogram"])) {
            self::$instances["Histogram"] = new self("histogram");
        }

        return self::$instances["Histogram"];
    }

    public static function nested(): self
    {
        if (!isset(self::$instances["Nested"])) {
            self::$instances["Nested"] = new self("nested");
        }

        return self::$instances["Nested"];
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
