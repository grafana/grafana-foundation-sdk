<?php

namespace Grafana\Foundation\Bigquery;

final class QueryPriority implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryPriority>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function interactive(): self
    {
        if (!isset(self::$instances["Interactive"])) {
            self::$instances["Interactive"] = new self("INTERACTIVE");
        }

        return self::$instances["Interactive"];
    }

    public static function batch(): self
    {
        if (!isset(self::$instances["Batch"])) {
            self::$instances["Batch"] = new self("BATCH");
        }

        return self::$instances["Batch"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "INTERACTIVE") {
            return self::interactive();
        }

        if ($value === "BATCH") {
            return self::batch();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryPriority");
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

