<?php

namespace Grafana\Foundation\Azuremonitor;

final class ResultFormat implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ResultFormat>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self("table");
        }

        return self::$instances["Table"];
    }

    public static function timeSeries(): self
    {
        if (!isset(self::$instances["TimeSeries"])) {
            self::$instances["TimeSeries"] = new self("time_series");
        }

        return self::$instances["TimeSeries"];
    }

    public static function trace(): self
    {
        if (!isset(self::$instances["Trace"])) {
            self::$instances["Trace"] = new self("trace");
        }

        return self::$instances["Trace"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "table") {
            return self::table();
        }

        if ($value === "time_series") {
            return self::timeSeries();
        }

        if ($value === "trace") {
            return self::trace();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ResultFormat");
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

