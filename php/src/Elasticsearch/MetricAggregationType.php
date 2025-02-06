<?php

namespace Grafana\Foundation\Elasticsearch;

final class MetricAggregationType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, MetricAggregationType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function count(): self
    {
        if (!isset(self::$instances["count"])) {
            self::$instances["count"] = new self("count");
        }

        return self::$instances["count"];
    }

    public static function avg(): self
    {
        if (!isset(self::$instances["avg"])) {
            self::$instances["avg"] = new self("avg");
        }

        return self::$instances["avg"];
    }

    public static function sum(): self
    {
        if (!isset(self::$instances["sum"])) {
            self::$instances["sum"] = new self("sum");
        }

        return self::$instances["sum"];
    }

    public static function min(): self
    {
        if (!isset(self::$instances["min"])) {
            self::$instances["min"] = new self("min");
        }

        return self::$instances["min"];
    }

    public static function max(): self
    {
        if (!isset(self::$instances["max"])) {
            self::$instances["max"] = new self("max");
        }

        return self::$instances["max"];
    }

    public static function extendedStats(): self
    {
        if (!isset(self::$instances["extended_stats"])) {
            self::$instances["extended_stats"] = new self("extended_stats");
        }

        return self::$instances["extended_stats"];
    }

    public static function percentiles(): self
    {
        if (!isset(self::$instances["percentiles"])) {
            self::$instances["percentiles"] = new self("percentiles");
        }

        return self::$instances["percentiles"];
    }

    public static function cardinality(): self
    {
        if (!isset(self::$instances["cardinality"])) {
            self::$instances["cardinality"] = new self("cardinality");
        }

        return self::$instances["cardinality"];
    }

    public static function rawDocument(): self
    {
        if (!isset(self::$instances["raw_document"])) {
            self::$instances["raw_document"] = new self("raw_document");
        }

        return self::$instances["raw_document"];
    }

    public static function rawData(): self
    {
        if (!isset(self::$instances["raw_data"])) {
            self::$instances["raw_data"] = new self("raw_data");
        }

        return self::$instances["raw_data"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["logs"])) {
            self::$instances["logs"] = new self("logs");
        }

        return self::$instances["logs"];
    }

    public static function rate(): self
    {
        if (!isset(self::$instances["rate"])) {
            self::$instances["rate"] = new self("rate");
        }

        return self::$instances["rate"];
    }

    public static function topMetrics(): self
    {
        if (!isset(self::$instances["top_metrics"])) {
            self::$instances["top_metrics"] = new self("top_metrics");
        }

        return self::$instances["top_metrics"];
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
        if ($value === "count") {
            return self::count();
        }

        if ($value === "avg") {
            return self::avg();
        }

        if ($value === "sum") {
            return self::sum();
        }

        if ($value === "min") {
            return self::min();
        }

        if ($value === "max") {
            return self::max();
        }

        if ($value === "extended_stats") {
            return self::extendedStats();
        }

        if ($value === "percentiles") {
            return self::percentiles();
        }

        if ($value === "cardinality") {
            return self::cardinality();
        }

        if ($value === "raw_document") {
            return self::rawDocument();
        }

        if ($value === "raw_data") {
            return self::rawData();
        }

        if ($value === "logs") {
            return self::logs();
        }

        if ($value === "rate") {
            return self::rate();
        }

        if ($value === "top_metrics") {
            return self::topMetrics();
        }

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

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricAggregationType");
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

