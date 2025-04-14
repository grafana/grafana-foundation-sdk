<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorPropertyType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BuilderQueryEditorPropertyType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function number(): self
    {
        if (!isset(self::$instances["number"])) {
            self::$instances["number"] = new self("number");
        }

        return self::$instances["number"];
    }

    public static function string(): self
    {
        if (!isset(self::$instances["string"])) {
            self::$instances["string"] = new self("string");
        }

        return self::$instances["string"];
    }

    public static function boolean(): self
    {
        if (!isset(self::$instances["boolean"])) {
            self::$instances["boolean"] = new self("boolean");
        }

        return self::$instances["boolean"];
    }

    public static function datetime(): self
    {
        if (!isset(self::$instances["datetime"])) {
            self::$instances["datetime"] = new self("datetime");
        }

        return self::$instances["datetime"];
    }

    public static function timeSpan(): self
    {
        if (!isset(self::$instances["time_span"])) {
            self::$instances["time_span"] = new self("time_span");
        }

        return self::$instances["time_span"];
    }

    public static function function(): self
    {
        if (!isset(self::$instances["function"])) {
            self::$instances["function"] = new self("function");
        }

        return self::$instances["function"];
    }

    public static function interval(): self
    {
        if (!isset(self::$instances["interval"])) {
            self::$instances["interval"] = new self("interval");
        }

        return self::$instances["interval"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "number") {
            return self::number();
        }

        if ($value === "string") {
            return self::string();
        }

        if ($value === "boolean") {
            return self::boolean();
        }

        if ($value === "datetime") {
            return self::datetime();
        }

        if ($value === "time_span") {
            return self::timeSpan();
        }

        if ($value === "function") {
            return self::function();
        }

        if ($value === "interval") {
            return self::interval();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BuilderQueryEditorPropertyType");
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

