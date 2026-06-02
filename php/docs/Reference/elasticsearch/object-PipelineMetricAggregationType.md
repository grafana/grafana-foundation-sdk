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
        if (!isset(self::$instances["MovingAvg"])) {
            self::$instances["MovingAvg"] = new self("moving_avg");
        }

        return self::$instances["MovingAvg"];
    }

    public static function movingFn(): self
    {
        if (!isset(self::$instances["MovingFn"])) {
            self::$instances["MovingFn"] = new self("moving_fn");
        }

        return self::$instances["MovingFn"];
    }

    public static function derivative(): self
    {
        if (!isset(self::$instances["Derivative"])) {
            self::$instances["Derivative"] = new self("derivative");
        }

        return self::$instances["Derivative"];
    }

    public static function serialDiff(): self
    {
        if (!isset(self::$instances["SerialDiff"])) {
            self::$instances["SerialDiff"] = new self("serial_diff");
        }

        return self::$instances["SerialDiff"];
    }

    public static function cumulativeSum(): self
    {
        if (!isset(self::$instances["CumulativeSum"])) {
            self::$instances["CumulativeSum"] = new self("cumulative_sum");
        }

        return self::$instances["CumulativeSum"];
    }

    public static function bucketScript(): self
    {
        if (!isset(self::$instances["BucketScript"])) {
            self::$instances["BucketScript"] = new self("bucket_script");
        }

        return self::$instances["BucketScript"];
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
