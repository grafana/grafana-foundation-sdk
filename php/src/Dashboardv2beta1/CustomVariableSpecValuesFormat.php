<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CustomVariableSpecValuesFormat implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, CustomVariableSpecValuesFormat>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function csv(): self
    {
        if (!isset(self::$instances["Csv"])) {
            self::$instances["Csv"] = new self("csv");
        }

        return self::$instances["Csv"];
    }

    public static function json(): self
    {
        if (!isset(self::$instances["Json"])) {
            self::$instances["Json"] = new self("json");
        }

        return self::$instances["Json"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "csv") {
            return self::csv();
        }

        if ($value === "json") {
            return self::json();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum CustomVariableSpecValuesFormat");
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

