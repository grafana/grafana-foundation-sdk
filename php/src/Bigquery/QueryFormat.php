<?php

namespace Grafana\Foundation\Bigquery;

final class QueryFormat implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, QueryFormat>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function timeseries(): self
    {
        if (!isset(self::$instances["Timeseries"])) {
            self::$instances["Timeseries"] = new self(0);
        }

        return self::$instances["Timeseries"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self(1);
        }

        return self::$instances["Table"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::timeseries();
        }

        if ($value === 1) {
            return self::table();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryFormat");
    }

    public function jsonSerialize(): int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

