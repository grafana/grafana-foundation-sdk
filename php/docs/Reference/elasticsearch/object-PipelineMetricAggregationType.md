---
title: <span class="badge object-type-enum"></span> PipelineMetricAggregationType
---
# <span class="badge object-type-enum"></span> PipelineMetricAggregationType

## Definition

```php
final class PipelineMetricAggregationType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PipelineMetricAggregationType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function movingAvg(): self
    {
        if (!isset(self::$instances["moving_avg"])) {
            self::$instances["moving_avg"] = new self("moving_avg");
        }

        return self::$instances["moving_avg"];
    }

    public static function movingFn(): self
    {
        if (!isset(self::$instances["moving_fn"])) {
            self::$instances["moving_fn"] = new self("moving_fn");
        }

        return self::$instances["moving_fn"];
    }

    public static function derivative(): self
    {
        if (!isset(self::$instances["derivative"])) {
            self::$instances["derivative"] = new self("derivative");
        }

        return self::$instances["derivative"];
    }

    public static function serialDiff(): self
    {
        if (!isset(self::$instances["serial_diff"])) {
            self::$instances["serial_diff"] = new self("serial_diff");
        }

        return self::$instances["serial_diff"];
    }

    public static function cumulativeSum(): self
    {
        if (!isset(self::$instances["cumulative_sum"])) {
            self::$instances["cumulative_sum"] = new self("cumulative_sum");
        }

        return self::$instances["cumulative_sum"];
    }

    public static function bucketScript(): self
    {
        if (!isset(self::$instances["bucket_script"])) {
            self::$instances["bucket_script"] = new self("bucket_script");
        }

        return self::$instances["bucket_script"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "moving_avg") {
            return self::movingAvg();
        }

        if ($value === "moving_fn") {
            return self::movingFn();
        }

        if ($value === "derivative") {
            return self::derivative();
        }

        if ($value === "serial_diff") {
            return self::serialDiff();
        }

        if ($value === "cumulative_sum") {
            return self::cumulativeSum();
        }

        if ($value === "bucket_script") {
            return self::bucketScript();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PipelineMetricAggregationType");
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
